package middleware

import "github.com/gofiber/fiber/v2"

type AuthAPIHandler struct {}

func NewAuthAPIHandler() *AuthAPIHandler {
    return &AuthAPIHandler{}
}

func (a *AuthAPIHandler) VerifyCredentials(c *fiber.Ctx) error {
    // fetch credentials from db
    // check credentials
    return c.Next()
}
