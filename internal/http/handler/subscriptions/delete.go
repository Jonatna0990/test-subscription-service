package handler

import (
	"database/sql"
	"errors"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/gofiber/fiber/v3"
)

// @Summary Удалить подписку
// @Description Удаляет подписку по её ID
// @Tags Подписки
// @Param id path string true "ID подписки"
// @Success 204 {object} dto.SuccessResponse "Удалено"
// @Failure 404 {object} dto.ErrorResponse "Не найдено"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [delete]
func (h *Handler) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	err := h.usecase.Delete(c.RequestCtx(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Message: "subscription not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "failed to delete subscription",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse{
		Message: "subscription deleted successfully",
	})
}
