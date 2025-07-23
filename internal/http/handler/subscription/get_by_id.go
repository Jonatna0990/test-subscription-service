package handler

import (
	"errors"
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
// @Failure 404 {object} dto.ErrorResponse "Не найдено"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [get]
func (h *Handler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid UUID format")
	}

	result, err := h.usecase.GetByID(c.RequestCtx(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "subscription not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get subscription")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
