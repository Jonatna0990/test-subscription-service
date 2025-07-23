package handler

import (
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Create(c fiber.Ctx) error {

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

	h.usecase.Create(c.RequestCtx(), req)

	return c.SendStatus(fiber.StatusCreated)
}
