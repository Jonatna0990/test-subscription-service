package dto

import "time"

// TODO
// подходит для create/update - потом можно сделать для каждого request'a свою структуру
type SubscriptionRequest struct {
	ServiceName string `json:"service_name" validate:"required"`
	Price       int    `json:"price" validate:"required,gt=0"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date,omitempty"`
}

type GetSubscriptionFilterListRequest struct {
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
	ID          string `json:"id"`
	ServiceName string `json:"service_name"`
}

// TODO
// подходит для create/read/update/list - потом можно сделать для каждого request'a свою структуру
type SubscriptionResponse struct {
	ID          string `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type GetSubscriptionFilterListResponse struct {
	TotalCost string     `json:"id"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}
