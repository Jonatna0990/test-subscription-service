package handler

import (
	"errors"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func (h *Handler) Update(c fiber.Ctx) error {

	id := c.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid UUID format")
	}

	_, err := h.usecase.GetByID(c.RequestCtx(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "subscription not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get subscription")
	}

	req := new(dto.SubscriptionRequest)

	if err := c.Bind().JSON(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorsMap := make(map[string]string)
			for _, e := range validationErrors {
				errorsMap[e.Field()] = e.Tag()
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "validation failed",
				"fields": errorsMap,
			})
		}
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
