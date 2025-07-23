package handler

import (
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/gofiber/fiber/v3"
)

// @Summary Получить все подписки
// @Description Возвращает список всех подписок
// @Tags Подписки
// @Produce json
// @Success 200 {array} dto.SubscriptionResponse "Список подписок"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions [get]
func (h *Handler) GetAll(c fiber.Ctx) error {
	result, err := h.usecase.GetAll(c.RequestCtx())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
