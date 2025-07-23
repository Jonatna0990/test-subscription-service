package handler

import "github.com/gofiber/fiber/v3"

// @Summary Удалить подписку
// @Description Удаляет подписку по её ID
// @Tags Подписки
// @Param id path string true "ID подписки"
// @Success 204 {string} string "Удалено"
// @Failure 404 {object} dto.ErrorResponse "Не найдено"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [delete]
func (h *Handler) Delete(ctx fiber.Ctx) error {
	return nil
}
