package handler

import (
	"go-fiber-api-template/internals"
	"go-fiber-api-template/internals/entity"
	"go-fiber-api-template/internals/usecase"

	"github.com/gofiber/fiber/v2"
)

type ForexFactoryHandler struct {
	usecase usecase.ForexfactoryUsecase
}

func NewForexFactoryHandler(usecase usecase.ForexfactoryUsecase) *ForexFactoryHandler {
	return &ForexFactoryHandler{usecase: usecase}
}

func (h *ForexFactoryHandler) GetForexFactory(c *fiber.Ctx) error {
	var body *entity.GetForexFactoryRequestBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	timeString := internals.PadZero(body.Day) + "/" + internals.PadZero(body.Month) + "/" + internals.PadZero(body.Year)

	data, err := h.usecase.GetForexFactory(timeString)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"data": data})
}

func (h *ForexFactoryHandler) UpsertForexFactory(c *fiber.Ctx) error {
	var body []*entity.CreateForexFactoryRequest
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	data, err := h.usecase.UpsertForexFactory(body)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"data": data})
}
