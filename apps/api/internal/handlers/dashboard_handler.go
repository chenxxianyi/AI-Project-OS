package handlers

import (
	"github.com/ai-project-os/api/internal/services"
	"github.com/ai-project-os/api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	svc *services.DashboardService
}

func NewDashboardHandler(svc *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{svc: svc}
}

func (h *DashboardHandler) Stats(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	stats, err := h.svc.GetStats(userID)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, stats)
}
