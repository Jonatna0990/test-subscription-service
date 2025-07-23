package subscriptions

import (
	"context"
	"errors"
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/pkg/utils"
	"time"
)

var (
	ErrDateFormat = errors.New("wrong date format")
)

func (r *repo) Create(ctx context.Context, s *dto.SubscriptionRequest) (string, error) {
	startDate, err := utils.ParseMonthYear(s.StartDate)
	if err != nil {
		return "", ErrDateFormat
	}

	var endDate *time.Time
	if s.EndDate != "" {
		parsedEndDate, err := utils.ParseMonthYear(s.EndDate)
		if err != nil {
			return "", ErrDateFormat
		}
		endDate = &parsedEndDate
	}

	id := utils.GenerateUUID()

	query := `INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
              VALUES ($1, $2, $3, $4, $5)`

	_, err = r.db.Exec(
		ctx,
		query,
		s.ServiceName,
		s.Price,
		id,
		startDate,
		endDate, // может быть NULL
	)

	if err != nil {
		r.logger.Error("failed to create subscription: " + err.Error())
		return "", err
	}

	r.logger.Info(fmt.Sprintf("Created subscription with ID: %s", id))
	return id, nil
}
