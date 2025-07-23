package usecase

import (
	"context"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/internal/repository/subscriptions"
)

type usecase struct {
	repo subscriptions.Repository
}

func New(repo subscriptions.Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) Create(ctx context.Context, s *dto.SubscriptionRequest) (string, error) {
	return u.repo.Create(ctx, s)
}

func (u *usecase) GetAll(ctx context.Context) ([]dto.SubscriptionResponse, error) {
	return u.repo.GetAll(ctx)
}

func (u *usecase) GetByID(ctx context.Context, id string) (*dto.SubscriptionResponse, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *usecase) Update(ctx context.Context, s *dto.SubscriptionRequest, id string) error {
	return u.repo.Update(ctx, s, id)
}

func (u *usecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

func (u *usecase) CalculateTotal(ctx context.Context, filter *dto.GetSubscriptionFilterListRequest) (dto.GetSubscriptionFilterListResponse, error) {
	return u.repo.CalculateTotal(ctx, filter)
}
