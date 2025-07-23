package handler

import (
	"errors"
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/internal/repository/subscriptions"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

// @Summary Обновить подписку
// @Description Обновляет данные существующей подписки
// @Tags Подписки
// @Accept json
// @Produce json
// @Param data body dto.SubscriptionRequest true "Обновлённые данные подписки"
// @Success 200 {object} dto.SuccessResponse "Обновлено"
// @Failure 400 {object} dto.ErrorResponse "Неверный запрос"
// @Failure 404 {object} dto.ErrorResponse "Не найдена подписка"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions [put]
func (h *Handler) Update(c fiber.Ctx) error {
	id := c.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "invalid UUID format",
		})
	}

	if len(c.Body()) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "request body cannot be empty",
		})
	}

	_, err := h.usecase.GetByID(c.RequestCtx(), id)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Message: "subscription not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Message: "internal server error",
			})
		}
	}

	req := new(dto.SubscriptionRequest)
	if err := c.Bind().JSON(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorsMap := make(map[string]string)
			for _, e := range validationErrors {
				errorsMap[e.Field()] = e.Tag()
			}
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Message: "validation failed",
				Fields:  errorsMap,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "invalid request format",
		})
	}

	err = h.usecase.Update(c.RequestCtx(), req, id)
	if err != nil {
		var status int
		var message string

		switch {
		case errors.Is(err, subscriptions.ErrDateFormat):
			status = fiber.StatusBadRequest
			message = "invalid date format (expected MM-YYYY)"
		default:
			status = fiber.StatusBadGateway
			message = "failed to update subscription"
		}

		return c.Status(status).JSON(dto.ErrorResponse{
			Message: message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse{
		Message: fmt.Sprintf("subscription updated: %s", id),
	})
}
