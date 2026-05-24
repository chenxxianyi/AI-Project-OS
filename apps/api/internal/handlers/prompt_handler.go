package handlers

import (
	"github.com/ai-project-os/api/internal/services"
	"github.com/ai-project-os/api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type PromptHandler struct {
	svc *services.PromptService
}

func NewPromptHandler(svc *services.PromptService) *PromptHandler {
	return &PromptHandler{svc: svc}
}

func (h *PromptHandler) Create(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	var req services.CreatePromptRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	userID := c.Locals("user_id").(string)
	p, err := h.svc.Create(projectID, userID, req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}
	return utils.Created(c, p)
}

func (h *PromptHandler) List(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	promptType := c.Query("type")
	prompts, err := h.svc.ListByProject(projectID, promptType)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, prompts)
}

func (h *PromptHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	p, err := h.svc.GetByID(id)
	if err != nil {
		return utils.NotFound(c, "prompt")
	}
	return utils.Success(c, p)
}

func (h *PromptHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req services.UpdatePromptRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	p, err := h.svc.Update(id, req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}
	return utils.Success(c, p)
}

func (h *PromptHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.Delete(id); err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, nil, "deleted")
}

func (h *PromptHandler) ListVersions(c *fiber.Ctx) error {
	id := c.Params("id")
	versions, err := h.svc.ListVersions(id)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, versions)
}
