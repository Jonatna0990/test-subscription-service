package handler

import "github.com/gofiber/fiber/v3"

// @Summary Подсчитать сумму подписок
// @Description Возвращает сумму всех подписок, соответствующих фильтру
// @Tags Подписки
// @Accept json
// @Produce json
// @Param data body dto.GetSubscriptionFilterListRequest true "Фильтр для подсчёта"
// @Success 200 {object} dto.GetSubscriptionFilterListResponse "Результат подсчёта"
// @Failure 400 {object} dto.ErrorResponse "Неверный запрос"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions/calculate-total [post]
func (h *Handler) CalculateTotal(ctx fiber.Ctx) error {
	return nil

}
