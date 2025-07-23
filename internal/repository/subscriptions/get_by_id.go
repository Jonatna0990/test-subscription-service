package subscriptions

import (
	"context"
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/pkg/utils"
	"time"
)

func (r *repo) GetByID(ctx context.Context, id string) (*dto.SubscriptionResponse, error) {
	query := `
		SELECT service_name, price, user_id, start_date, end_date
		FROM subscriptions
		WHERE user_id = $1
	`

	var sub dto.SubscriptionResponse
	var startDate *time.Time
	var endDate *time.Time

	row := r.postgres.Pool.QueryRow(ctx, query, id)
	err := row.Scan(
		&sub.ServiceName,
		&sub.Price,
		&sub.UserID,
		&startDate,
		&endDate,
	)
	if err != nil {
		r.logger.Error("get subscription by id error: " + err.Error())
		return nil, fmt.Errorf("get subscription by id: %w", err)
	}

	if startDate != nil {
		sub.StartDate = utils.FormatMonthYear(*startDate)
	}
	if endDate != nil {
		sub.EndDate = utils.FormatMonthYear(*endDate)
	}

	return &sub, nil
}
