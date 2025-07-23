package subscription

import (
	"context"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/pkg/utils"
	"time"
)

func (r *repo) GetAll(ctx context.Context) ([]dto.SubscriptionResponse, error) {
	//TODO сделать лимит и смещение(пагинацию)
	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions`
	rows, err := r.postgres.Pool.Query(ctx, query)
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
		err = rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &startDate, &endDate)
		if err != nil {
			return nil, err
		}
		s.StartDate = utils.FormatMonthYear(*startDate)
		s.EndDate = utils.FormatMonthYear(*endDate)
		subs = append(subs, s)
	}
	return subs, nil
}
