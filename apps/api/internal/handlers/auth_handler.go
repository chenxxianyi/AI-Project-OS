package handlers

import (
	"time"

	"github.com/ai-project-os/api/internal/services"
	"github.com/ai-project-os/api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	svc *services.AuthService
}

func NewAuthHandler(svc *services.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req services.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	resp, err := h.svc.Register(req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}
	h.setAuthCookie(c, resp.Token)
	return utils.Created(c, resp)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req services.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	resp, err := h.svc.Login(req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}
	h.setAuthCookie(c, resp.Token)
	return utils.Success(c, resp)
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteLaxMode,
		Path:     "/",
	})
	return utils.Success(c, nil, "logged out")
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	user, err := h.svc.GetCurrentUser(userID)
	if err != nil {
		return utils.NotFound(c, "user")
	}
	return utils.Success(c, user)
}

func (h *AuthHandler) setAuthCookie(c *fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(h.svc.TokenDuration()),
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: fiber.CookieSameSiteLaxMode,
		Path:     "/",
	})
}
