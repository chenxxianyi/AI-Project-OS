package handlers

import (
	"github.com/ai-project-os/api/internal/services"
	"github.com/ai-project-os/api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type AIGenerationHandler struct {
	svc *services.AIGenerationService
}

func NewAIGenerationHandler(svc *services.AIGenerationService) *AIGenerationHandler {
	return &AIGenerationHandler{svc: svc}
}

func (h *AIGenerationHandler) Generate(c *fiber.Ctx) error {
	var req services.GeneratePromptRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	userID := c.Locals("user_id").(string)
	gen, err := h.svc.Generate(userID, req)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Created(c, gen)
}

func (h *AIGenerationHandler) List(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	gens, err := h.svc.ListByUser(userID, 1, 20)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, gens)
}
