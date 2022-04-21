package perfhandler

import (
	"github.com/artback/hygh/pkg/queryperf"
	"github.com/gofiber/fiber/v2"
)

type PerfHandler struct {
	queryperf.Repository
}

func (p PerfHandler) GetQueryPerformance(c *fiber.Ctx) error {
	var options queryperf.Options
	if err := c.QueryParser(&options); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	time, err := p.QueriesByMeanTime(c.Context(), options)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// Vague description: do we want JSON response?
	return c.JSON(time)
}
