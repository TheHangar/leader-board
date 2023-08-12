package middleware

import "github.com/gofiber/fiber/v2"

type AuthUserHandler struct {}

func NewAuthUserHandler() *AuthUserHandler {
    return &AuthUserHandler{}
}

func (a *AuthUserHandler) CreateToken(c *fiber.Ctx) error {
    // create token and pass in http only cookie
    return c.Next()
}

func (a *AuthUserHandler) VerifyToken(c *fiber.Ctx) error {
    // verify cookie
    // token validation
    return c.Next()
}
