package subscription

import (
	"context"
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/pkg/utils"
	"time"
)

func (r *repo) GetByID(ctx context.Context, id string) (*dto.SubscriptionResponse, error) {
	query := `
		SELECT id, service_name, price, user_id, start_date, end_date
		FROM subscriptions
		WHERE id = $1
	`

	var sub dto.SubscriptionResponse
	var startDate *time.Time
	var endDate *time.Time

	row := r.postgres.Pool.QueryRow(ctx, query, id)
	err := row.Scan(
		&sub.ID,
		&sub.ServiceName,
		&sub.Price,
		&sub.UserID,
		&startDate,
		&endDate,
	)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, fmt.Errorf("get subscription by id: %w", err)
	}

	sub.StartDate = utils.FormatMonthYear(*startDate)
	sub.EndDate = utils.FormatMonthYear(*endDate)

	return &sub, nil
}
