package handler

import (
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetAll(c fiber.Ctx) error {
	result, err := h.usecase.GetAll(c.RequestCtx())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
