package handler

import (
	"errors"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

// @Summary Получить подписку по ID
// @Description Возвращает подписку по её идентификатору
// @Tags Подписки
// @Produce json
// @Param id path string true "ID подписки"
// @Success 200 {object} dto.SubscriptionResponse "Найденная подписка"
// @Failure 400 {object} dto.ErrorResponse "Неправильный uuid"
// @Failure 404 {object} dto.ErrorResponse "Не найдено"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [get]
func (h *Handler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "invalid UUID format",
		})
	}

	result, err := h.usecase.GetByID(c.RequestCtx(), id)
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

	return c.Status(fiber.StatusOK).JSON(result)
}
