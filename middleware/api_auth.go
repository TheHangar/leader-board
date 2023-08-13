package middleware

import (

	"github.com/TheHangar/leader-board/store"
	"github.com/gofiber/fiber/v2"
)

type AuthAPIHandler struct {
    gameStore store.GameStore
}

func NewAuthAPIHandler(gs store.GameStore) *AuthAPIHandler {
    return &AuthAPIHandler{ gameStore: gs }
}

func (a *AuthAPIHandler) VerifyCredentials(c *fiber.Ctx) error {
    api_key := c.Get("X-Api-Key")
    if api_key == "" {
        return c.Status(401).JSON(map[string]string{"message": "no api key provided"})
    }
    gameUUID := c.Params("uuid")

    game, err := a.gameStore.GetGameByUUID(gameUUID)

    if err != nil {
        return c.Status(500).JSON(map[string]string{"message": "database error"})
    }

    if game.ApiKey != api_key {
        return c.Status(401).JSON(map[string]string{"message": "unauthorized"})
    }

    return c.Next()
}
