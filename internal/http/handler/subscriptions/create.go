package handler

import (
	"errors"
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/internal/repository/subscriptions"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// @Summary Создать подписку
// @Description Создаёт новую подписку
// @Tags Подписки
// @Accept json
// @Produce json
// @Param data body dto.SubscriptionRequest true "Данные подписки"
// @Success 201 {object} dto.SuccessResponse "Создано"
// @Failure 400 {object} dto.ErrorResponse "Неверный запрос"
// @Failure 502 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions [post]
func (h *Handler) Create(c fiber.Ctx) error {
	req := new(dto.SubscriptionRequest)
	// TODO можно улучшить валидацию и возвращаемые ошибки

	if len(c.Body()) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "request body cannot be empty",
		})
	}

	if err := c.Bind().JSON(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorsMap := make(map[string]string)
			for _, e := range validationErrors {
				errorsMap[e.Field()] = e.Tag()
			}
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Message: "wrong param",
				Fields:  errorsMap,
			})

		}
		return c.Status(fiber.StatusBadGateway).JSON(dto.ErrorResponse{
			Message: "unexpected error",
		})
	}

	id, err := h.usecase.Create(c.RequestCtx(), req)
	if err != nil {
		var status int
		var message string

		switch {
		case errors.Is(err, subscriptions.ErrDateFormat):
			status = fiber.StatusBadRequest
			message = "invalid date format (expected MM-YYYY)"
		default:
			status = fiber.StatusBadGateway
			message = "failed to create subscription"
		}

		return c.Status(status).JSON(dto.ErrorResponse{
			Message: message,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.SuccessResponse{
		Message: fmt.Sprintf("subscription created: %s", id),
	})
}
