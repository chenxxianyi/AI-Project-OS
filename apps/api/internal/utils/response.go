package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Code    string      `json:"code"`
	Details interface{} `json:"details,omitempty"`
}

func Success(c *fiber.Ctx, data interface{}, message ...string) error {
	msg := "ok"
	if len(message) > 0 {
		msg = message[0]
	}
	return c.Status(http.StatusOK).JSON(APIResponse{
		Success: true,
		Data:    data,
		Message: msg,
	})
}

func Created(c *fiber.Ctx, data interface{}, message ...string) error {
	msg := "created"
	if len(message) > 0 {
		msg = message[0]
	}
	return c.Status(http.StatusCreated).JSON(APIResponse{
		Success: true,
		Data:    data,
		Message: msg,
	})
}

func Error(c *fiber.Ctx, status int, code string, details interface{}) error {
	return c.Status(status).JSON(APIResponse{
		Success: false,
		Data:    nil,
		Message: code,
		Error:   &APIError{Code: code, Details: details},
	})
}

func BadRequest(c *fiber.Ctx, details interface{}) error {
	return Error(c, http.StatusBadRequest, "INVALID_INPUT", details)
}

func Unauthorized(c *fiber.Ctx) error {
	return Error(c, http.StatusUnauthorized, "UNAUTHORIZED", nil)
}

func Forbidden(c *fiber.Ctx) error {
	return Error(c, http.StatusForbidden, "FORBIDDEN", nil)
}

func NotFound(c *fiber.Ctx, resource string) error {
	return Error(c, http.StatusNotFound, "NOT_FOUND", resource)
}

func InternalError(c *fiber.Ctx, details interface{}) error {
	return Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", details)
}
