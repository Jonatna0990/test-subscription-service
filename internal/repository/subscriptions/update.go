package subscriptions

import (
	"context"
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/pkg/utils"
	"time"
)

func (r *repo) Update(ctx context.Context, s *dto.SubscriptionRequest, id string) error {
	startDate, err := utils.ParseMonthYear(s.StartDate)
	if err != nil {
		return ErrDateFormat
	}

	var endDate *time.Time
	if s.EndDate != "" {
		parsedEnd, err := utils.ParseMonthYear(s.EndDate)
		if err != nil {
			return ErrDateFormat
		}
		endDate = &parsedEnd
	}

	query := `
		UPDATE subscriptions
		SET service_name = $1,
		    price = $2,
		    start_date = $3,
		    end_date = $4
		WHERE user_id = $5
	`

	_, err = r.postgres.Pool.Exec(ctx, query,
		s.ServiceName,
		s.Price,
		startDate,
		endDate,
		id,
	)

	if err != nil {
		r.logger.Error("failed to update subscription: " + err.Error())
		return err
	}

	r.logger.Info(fmt.Sprintf("Updated subscription with ID: %s", id))
	return nil
}
