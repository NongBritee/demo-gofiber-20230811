package query

import "github.com/gofiber/fiber/v2"

type repo interface {
	QueryAll() ([]QueryItem, error)
}

func NewHandler(r repo) *handler {
	return &handler{repo: r}
}

type handler struct {
	repo repo
}

type QueryItem struct {
	ActivePower int `json:"active_power" `
	PowerInput  int `json:"power_input" `
}

type SumResponse struct {
	TotalActivePower int `json:"total_active_power" `
	TotalPowerInput  int `json:"total_power_input" `
}

func (h handler) SumHandler(c *fiber.Ctx) error {

	resp, err := h.repo.QueryAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// sum all
	var totalActivePower int
	var totalPowerInput int
	for _, item := range resp {
		totalActivePower += item.ActivePower
		totalPowerInput += item.PowerInput
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_active_power": totalActivePower,
		"total_power_input":  totalPowerInput,
	})
}

func (h handler) GetAllHandler(c *fiber.Ctx) error {
	resp, err := h.repo.QueryAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": resp,
	})
}
