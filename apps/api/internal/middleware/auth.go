package middleware

import (
	"strings"
	"time"

	"github.com/ai-project-os/api/internal/config"
	"github.com/ai-project-os/api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		tokenString := c.Cookies("access_token")
		if authHeader != "" {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return utils.Unauthorized(c)
			}
			tokenString = parts[1]
		}
		if tokenString == "" {
			return utils.Unauthorized(c)
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			return utils.Unauthorized(c)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utils.Unauthorized(c)
		}

		userID, ok := claims["sub"].(string)
		if !ok || userID == "" {
			return utils.Unauthorized(c)
		}
		role, _ := claims["role"].(string)

		c.Locals("user_id", userID)
		c.Locals("user_role", role)
		return c.Next()
	}
}

func GenerateToken(cfg *config.Config, userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(time.Duration(cfg.JWTExpirationHrs) * time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}
