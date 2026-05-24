package handlers

import (
	"github.com/ai-project-os/api/internal/services"
	"github.com/ai-project-os/api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type ProjectDocHandler struct {
	svc *services.ProjectDocService
}

func NewProjectDocHandler(svc *services.ProjectDocService) *ProjectDocHandler {
	return &ProjectDocHandler{svc: svc}
}

func (h *ProjectDocHandler) Create(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	var req services.CreateDocRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	userID := c.Locals("user_id").(string)
	doc, err := h.svc.Create(projectID, userID, req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}
	return utils.Created(c, doc)
}

func (h *ProjectDocHandler) List(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	docType := c.Query("doc_type")
	docs, err := h.svc.ListByProject(projectID, docType)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, docs)
}

func (h *ProjectDocHandler) Get(c *fiber.Ctx) error {
	id := c.Params("docId")
	doc, err := h.svc.GetByID(id)
	if err != nil {
		return utils.NotFound(c, "document")
	}
	return utils.Success(c, doc)
}

func (h *ProjectDocHandler) Update(c *fiber.Ctx) error {
	id := c.Params("docId")
	var req services.UpdateDocRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	doc, err := h.svc.Update(id, req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}
	return utils.Success(c, doc)
}

func (h *ProjectDocHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("docId")
	if err := h.svc.Delete(id); err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, nil, "deleted")
}
