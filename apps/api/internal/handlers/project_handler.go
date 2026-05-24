package handlers

import (
	"strconv"

	"github.com/ai-project-os/api/internal/services"
	"github.com/ai-project-os/api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type ProjectHandler struct {
	svc *services.ProjectService
}

func NewProjectHandler(svc *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{svc: svc}
}

func (h *ProjectHandler) Create(c *fiber.Ctx) error {
	var req services.CreateProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	userID := c.Locals("user_id").(string)
	p, err := h.svc.Create(userID, req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}
	return utils.Created(c, p)
}

func (h *ProjectHandler) List(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	projectType := c.Query("project_type")
	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "20"))

	result, err := h.svc.List(userID, projectType, search, page, pageSize)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, result)
}

func (h *ProjectHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	p, err := h.svc.GetByID(id)
	if err != nil {
		return utils.NotFound(c, "project")
	}
	return utils.Success(c, p)
}

func (h *ProjectHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req services.UpdateProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	p, err := h.svc.Update(id, req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}
	return utils.Success(c, p)
}

func (h *ProjectHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.Delete(id); err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, nil, "deleted")
}
