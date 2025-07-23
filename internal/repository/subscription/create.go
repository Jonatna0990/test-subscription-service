package subscription

import (
	"context"
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/pkg/utils"
)

func (r *repo) Create(ctx context.Context, s *dto.SubscriptionRequest) error {
	startDate, err := utils.ParseMonthYear(s.StartDate)
	if err != nil {
		return err
	}
	id := utils.GenerateUUID()
	query := `INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
	              VALUES ($1, $2, $3, $4, $5)`
	_, err = r.postgres.Pool.Exec(ctx, query, s.ServiceName, s.Price, id, startDate, startDate)
	if err != nil {
		r.logger.Error(err.Error())
	}
	r.logger.Info(fmt.Sprintf("Created subscription with ID: %s ", id))
	return err
}
