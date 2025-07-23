package subscriptions

import (
	"context"
	"errors"
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/pkg/utils"
)

var ErrDateRange = errors.New("endDate cannot be before startDate")

func (r *repo) CalculateTotal(ctx context.Context, filter *dto.GetSubscriptionFilterListRequest) (dto.GetSubscriptionFilterListResponse, error) {
	startDate, err := utils.ParseMonthYear(filter.StartDate)
	if err != nil {
		return dto.GetSubscriptionFilterListResponse{}, ErrDateFormat
	}

	endDate, err := utils.ParseMonthYear(filter.EndDate)
	if err != nil {
		return dto.GetSubscriptionFilterListResponse{}, ErrDateFormat
	}

	// Проверяем, что endDate не меньше startDate
	if endDate.Before(startDate) {
		return dto.GetSubscriptionFilterListResponse{}, ErrDateRange
	}

	// Базовый запрос с условиями
	query := `
    SELECT COALESCE(SUM(price), 0)
    FROM subscriptions
    WHERE start_date <= $1 AND (end_date IS NULL OR end_date >= $2)
`

	// Аргументы для запроса
	args := []interface{}{endDate, startDate}
	argPos := 3 // для нумерации $3, $4 и т.д.

	if filter.UserID != "" {
		query += fmt.Sprintf(" AND user_id = $%d", argPos)
		args = append(args, filter.UserID)
		argPos++
	}

	if filter.ServiceName != "" {
		query += fmt.Sprintf(" AND service_name = $%d", argPos)
		args = append(args, filter.ServiceName)
		argPos++
	}

	var totalCost int
	err = r.db.QueryRow(ctx, query, args...).Scan(&totalCost)
	if err != nil {
		r.logger.Error("CalculateTotal query error: " + err.Error())
		return dto.GetSubscriptionFilterListResponse{}, err
	}

	return dto.GetSubscriptionFilterListResponse{
		TotalCost: totalCost,
		StartDate: &startDate,
		EndDate:   &endDate,
	}, nil
}
