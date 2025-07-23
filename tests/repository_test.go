package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/internal/repository/subscriptions"
	"github.com/Jonatna0990/test-subscription-service/pkg/utils"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
	"io"
	"log/slog"
	"testing"
	"time"
)

func newNoopLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func TestCalculateTotal_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()
	logger := newNoopLogger()

	repo := subscriptions.NewRepository(mock, logger)

	filter := &dto.GetSubscriptionFilterListRequest{
		StartDate:   "06-2025",
		EndDate:     "07-2025",
		UserID:      "user-uuid",
		ServiceName: "Netflix",
	}

	start, _ := utils.ParseMonthYear(filter.StartDate)
	end, _ := utils.ParseMonthYear(filter.EndDate)

	mock.ExpectQuery(`SELECT COALESCE\(SUM\(price\), 0\) FROM subscriptions`).
		WithArgs(end, start, filter.UserID, filter.ServiceName).
		WillReturnRows(pgxmock.NewRows([]string{"coalesce"}).AddRow(900))

	res, err := repo.CalculateTotal(context.Background(), filter)
	require.NoError(t, err)
	require.Equal(t, 900, res.TotalCost)
	require.Equal(t, &start, res.StartDate)
	require.Equal(t, &end, res.EndDate)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()

	repo := subscriptions.NewRepository(mock, logger)

	subReq := &dto.SubscriptionRequest{
		ServiceName: "Yandex Plus",
		Price:       400,
		StartDate:   "07-2025",
		EndDate:     "",
	}

	mock.ExpectExec(`INSERT INTO subscriptions \(service_name, price, user_id, start_date, end_date\)`).
		WithArgs(
			subReq.ServiceName,
			subReq.Price,
			pgxmock.AnyArg(),
			mustParseMonthYear(subReq.StartDate),
			nil,
		).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	ctx := context.Background()
	id, err := repo.Create(ctx, subReq)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDelete_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()
	id := "60601fee-2bf1-4721-ae6f-7636e79a0cba"

	mock.ExpectExec(`DELETE FROM subscriptions WHERE user_id = \$1`).
		WithArgs(id).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = repo.Delete(ctx, id)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDelete_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()
	id := "nonexistent-id"

	mock.ExpectExec(`DELETE FROM subscriptions WHERE user_id = \$1`).
		WithArgs(id).
		WillReturnResult(pgxmock.NewResult("DELETE", 0))

	err = repo.Delete(ctx, id)
	require.ErrorIs(t, err, sql.ErrNoRows)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDelete_DBError(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()
	id := "some-id"

	mock.ExpectExec(`DELETE FROM subscriptions WHERE user_id = \$1`).
		WithArgs(id).
		WillReturnError(sql.ErrConnDone)
	err = repo.Delete(ctx, id)
	require.ErrorIs(t, err, sql.ErrConnDone)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAll_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()

	startDate1 := time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)
	endDate1 := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)

	rows := pgxmock.NewRows([]string{"service_name", "price", "user_id", "start_date", "end_date"}).
		AddRow("Netflix", 500, "user-1-uuid", startDate1, endDate1).
		AddRow("Spotify", 300, "user-2-uuid", startDate1, nil)

	mock.ExpectQuery(`SELECT service_name, price, user_id, start_date, end_date FROM subscriptions`).
		WillReturnRows(rows)

	result, err := repo.GetAll(ctx)
	require.NoError(t, err)
	require.Len(t, result, 2)

	require.Equal(t, "Netflix", result[0].ServiceName)
	require.Equal(t, 500, result[0].Price)
	require.Equal(t, "user-1-uuid", result[0].UserID)
	require.Equal(t, utils.FormatMonthYear(startDate1), result[0].StartDate)
	require.Equal(t, utils.FormatMonthYear(endDate1), result[0].EndDate)

	require.Equal(t, "Spotify", result[1].ServiceName)
	require.Equal(t, 300, result[1].Price)
	require.Equal(t, "user-2-uuid", result[1].UserID)
	require.Equal(t, utils.FormatMonthYear(startDate1), result[1].StartDate)
	require.Equal(t, "", result[1].EndDate)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAll_QueryError(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()

	mock.ExpectQuery(`SELECT service_name, price, user_id, start_date, end_date FROM subscriptions`).
		WillReturnError(io.ErrUnexpectedEOF)

	result, err := repo.GetAll(ctx)
	require.Error(t, err)
	require.Nil(t, result)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAll_ScanError(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()

	rows := pgxmock.NewRows([]string{"service_name", "price", "user_id", "start_date", "end_date"}).
		AddRow("Netflix", "not-an-int", "user-1-uuid", time.Now(), nil)

	mock.ExpectQuery(`SELECT service_name, price, user_id, start_date, end_date FROM subscriptions`).
		WillReturnRows(rows)

	result, err := repo.GetAll(ctx)
	require.Error(t, err)
	require.Nil(t, result)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetByID_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()
	userID := "user-1-uuid"

	startDate := time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)

	rows := pgxmock.NewRows([]string{"service_name", "price", "user_id", "start_date", "end_date"}).
		AddRow("Netflix", 500, userID, startDate, endDate)

	mock.ExpectQuery(`SELECT service_name, price, user_id, start_date, end_date FROM subscriptions WHERE user_id = \$1`).
		WithArgs(userID).
		WillReturnRows(rows)

	sub, err := repo.GetByID(ctx, userID)
	require.NoError(t, err)
	require.NotNil(t, sub)

	require.Equal(t, "Netflix", sub.ServiceName)
	require.Equal(t, 500, sub.Price)
	require.Equal(t, userID, sub.UserID)
	require.Equal(t, utils.FormatMonthYear(startDate), sub.StartDate)
	require.Equal(t, utils.FormatMonthYear(endDate), sub.EndDate)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetByID_NoRows(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()
	userID := "nonexistent-user-id"

	mock.ExpectQuery(`SELECT service_name, price, user_id, start_date, end_date FROM subscriptions WHERE user_id = \$1`).
		WithArgs(userID).
		WillReturnError(fmt.Errorf("no rows in result set"))

	sub, err := repo.GetByID(ctx, userID)
	require.Error(t, err)
	require.Nil(t, sub)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetByID_ScanError(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()
	userID := "user-1-uuid"

	rows := pgxmock.NewRows([]string{"service_name", "price", "user_id", "start_date", "end_date"}).
		AddRow("Netflix", "invalid-price", userID, time.Now(), nil)

	mock.ExpectQuery(`SELECT service_name, price, user_id, start_date, end_date FROM subscriptions WHERE user_id = \$1`).
		WithArgs(userID).
		WillReturnRows(rows)

	sub, err := repo.GetByID(ctx, userID)
	require.Error(t, err)
	require.Nil(t, sub)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdate_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()
	id := "user-uuid-1"

	req := &dto.SubscriptionRequest{
		ServiceName: "Spotify",
		Price:       999,
		StartDate:   "2023-07",
		EndDate:     "2024-07",
	}

	startDate, err := utils.ParseMonthYear(req.StartDate)
	require.NoError(t, err)
	endDate, err := utils.ParseMonthYear(req.EndDate)
	require.NoError(t, err)

	mock.ExpectExec(`UPDATE subscriptions
		SET service_name = \$1,
		    price = \$2,
		    start_date = \$3,
		    end_date = \$4
		WHERE user_id = \$5`).
		WithArgs(req.ServiceName, req.Price, startDate, &endDate, id).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = repo.Update(ctx, req, id)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdate_InvalidStartDate(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	req := &dto.SubscriptionRequest{
		ServiceName: "Spotify",
		Price:       999,
		StartDate:   "invalid-date",
	}

	err = repo.Update(context.Background(), req, "some-id")
	require.ErrorIs(t, err, subscriptions.ErrDateFormat)
}

func TestUpdate_InvalidEndDate(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	req := &dto.SubscriptionRequest{
		ServiceName: "Spotify",
		Price:       999,
		StartDate:   "2023-07",
		EndDate:     "invalid-end-date",
	}

	err = repo.Update(context.Background(), req, "some-id")
	require.ErrorIs(t, err, subscriptions.ErrDateFormat)
}

func TestUpdate_DBExecError(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	logger := newNoopLogger()
	repo := subscriptions.NewRepository(mock, logger)

	ctx := context.Background()
	id := "user-uuid-1"

	req := &dto.SubscriptionRequest{
		ServiceName: "Spotify",
		Price:       999,
		StartDate:   "2023-07",
		EndDate:     "2024-07",
	}

	startDate, err := utils.ParseMonthYear(req.StartDate)
	require.NoError(t, err)
	endDate, err := utils.ParseMonthYear(req.EndDate)
	require.NoError(t, err)

	mock.ExpectExec(`UPDATE subscriptions
		SET service_name = \$1,
		    price = \$2,
		    start_date = \$3,
		    end_date = \$4
		WHERE user_id = \$5`).
		WithArgs(req.ServiceName, req.Price, startDate, &endDate, id).
		WillReturnError(fmt.Errorf("db exec error"))

	err = repo.Update(ctx, req, id)
	require.Error(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func mustParseMonthYear(s string) time.Time {
	t, err := utils.ParseMonthYear(s)
	if err != nil {
		panic(err)
	}
	return t
}
