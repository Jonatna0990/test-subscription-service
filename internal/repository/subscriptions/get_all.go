package subscriptions

import (
	"context"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/pkg/utils"
	"time"
)

func (r *repo) GetAll(ctx context.Context) ([]dto.SubscriptionResponse, error) {
	// TODO: добавить лимит и смещение (пагинацию)
	query := `SELECT service_name, price, user_id, start_date, end_date FROM subscriptions`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var subs []dto.SubscriptionResponse
	for rows.Next() {
		var s dto.SubscriptionResponse
		var startDate *time.Time
		var endDate *time.Time

		err = rows.Scan(&s.ServiceName, &s.Price, &s.UserID, &startDate, &endDate)
		if err != nil {
			r.logger.Error("failed to scan subscription: " + err.Error())
			return nil, err
		}

		if startDate != nil {
			s.StartDate = utils.FormatMonthYear(*startDate)
		}

		if endDate != nil {
			s.EndDate = utils.FormatMonthYear(*endDate)
		}

		subs = append(subs, s)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("error after row iteration: " + err.Error())
		return nil, err
	}

	return subs, nil
}
