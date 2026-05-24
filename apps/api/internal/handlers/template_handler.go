package handlers

import (
	"strconv"

	"github.com/ai-project-os/api/internal/services"
	"github.com/ai-project-os/api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type TemplateHandler struct {
	svc *services.TemplateService
}

func NewTemplateHandler(svc *services.TemplateService) *TemplateHandler {
	return &TemplateHandler{svc: svc}
}

func (h *TemplateHandler) List(c *fiber.Ctx) error {
	projectType := c.Query("project_type")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "20"))
	result, err := h.svc.List(projectType, page, pageSize)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, result)
}

func (h *TemplateHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	t, err := h.svc.GetByID(id)
	if err != nil {
		return utils.NotFound(c, "template")
	}
	return utils.Success(c, t)
}
