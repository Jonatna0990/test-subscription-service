package handler

import (
	"github.com/Jonatna0990/test-subscription-service/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	usecase usecase.UseCase
}

func New(u usecase.UseCase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	r := router.Group("/subscriptions")
	r.Post("/", h.Create)
	r.Get("/", h.GetAll)
	r.Get("/summary", h.CalculateTotal)
	r.Get("/:id", h.GetByID)
	r.Put("/:id", h.Update)
	r.Delete("/:id", h.Delete)
}
