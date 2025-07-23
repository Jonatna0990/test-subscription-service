package usecase

import (
	"context"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
)

type UseCase interface {
	Create(ctx context.Context, s *dto.SubscriptionRequest) (string, error)
	GetAll(ctx context.Context) ([]dto.SubscriptionResponse, error)
	GetByID(ctx context.Context, id string) (*dto.SubscriptionResponse, error)
	Update(ctx context.Context, s *dto.SubscriptionRequest, id string) error
	Delete(ctx context.Context, id string) error
	CalculateTotal(ctx context.Context, filter *dto.GetSubscriptionFilterListRequest) (dto.GetSubscriptionFilterListResponse, error)
}
