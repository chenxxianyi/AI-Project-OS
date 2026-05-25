package handlers

import (
	"github.com/ai-project-os/api/internal/services"
	"github.com/ai-project-os/api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	svc *services.AdminService
}

func NewAdminHandler(svc *services.AdminService) *AdminHandler {
	return &AdminHandler{svc: svc}
}

// GET /admin/stats
func (h *AdminHandler) Stats(c *fiber.Ctx) error {
	stats, err := h.svc.GetSystemStats()
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, stats)
}

// GET /admin/health
func (h *AdminHandler) Health(c *fiber.Ctx) error {
	return utils.Success(c, h.svc.GetSystemHealth())
}

// GET /admin/settings
func (h *AdminHandler) GetSettings(c *fiber.Ctx) error {
	return utils.Success(c, h.svc.GetSettings())
}

// PUT /admin/settings
func (h *AdminHandler) UpdateSettings(c *fiber.Ctx) error {
	var req services.AdminSystemSettings
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	return utils.Success(c, h.svc.UpdateSettings(req))
}

// GET /admin/users?search=&role=&page=&page_size=
func (h *AdminHandler) ListUsers(c *fiber.Ctx) error {
	search := c.Query("search")
	role := c.Query("role")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 20)

	resp, err := h.svc.ListUsers(search, role, page, pageSize)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, resp)
}

// PUT /admin/users/:id
func (h *AdminHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var req services.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	user, err := h.svc.UpdateUser(id, req)
	if err != nil {
		if err.Error() == "user not found" {
			return utils.NotFound(c, "user")
		}
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, user)
}

// DELETE /admin/users/:id
func (h *AdminHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.DeleteUser(id); err != nil {
		if err.Error() == "user not found" {
			return utils.NotFound(c, "user")
		}
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, nil)
}

// GET /admin/projects?search=&status=&page=&page_size=
func (h *AdminHandler) ListProjects(c *fiber.Ctx) error {
	search := c.Query("search")
	status := c.Query("status")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 20)

	resp, err := h.svc.ListProjects(search, status, page, pageSize)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, resp)
}

// DELETE /admin/projects/:id
func (h *AdminHandler) DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.DeleteProject(id); err != nil {
		if err.Error() == "project not found" {
			return utils.NotFound(c, "project")
		}
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, nil)
}

// GET /admin/generations?status=&generation_type=&page=&page_size=
func (h *AdminHandler) ListGenerations(c *fiber.Ctx) error {
	status := c.Query("status")
	genType := c.Query("generation_type")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 20)

	resp, err := h.svc.ListGenerations(status, genType, page, pageSize)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}
	return utils.Success(c, resp)
}
