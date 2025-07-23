package handler

import (
	"errors"
	"github.com/Jonatna0990/test-subscription-service/internal/dto"
	"github.com/Jonatna0990/test-subscription-service/internal/repository/subscriptions"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// @Summary Подсчитать сумму подписок
// @Description Возвращает сумму всех подписок, соответствующих фильтру
// @Tags Подписки
// @Accept json
// @Produce json
// @Param data body dto.GetSubscriptionFilterListRequest true "Фильтр для подсчёта"
// @Success 200 {object} dto.GetSubscriptionFilterListResponse "Результат подсчёта"
// @Failure 400 {object} dto.ErrorResponse "Неверный запрос"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions/summary [get]
func (h *Handler) CalculateTotal(c fiber.Ctx) error {
	req := new(dto.GetSubscriptionFilterListRequest)
	if err := c.Bind().Query(req); err != nil {
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

	result, err := h.usecase.CalculateTotal(c.RequestCtx(), req)
	if err != nil {
		if errors.Is(err, subscriptions.ErrDateFormat) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Message: "invalid date format (expected MM-YYYY)",
			})
		}

		if errors.Is(err, subscriptions.ErrDateRange) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Message: "start date older than end date)",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "failed to calculate total",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)

}
