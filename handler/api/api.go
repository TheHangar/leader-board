package api

import (
	custometype "github.com/TheHangar/leader-board/custome_type"
	"github.com/TheHangar/leader-board/store"
	"github.com/gofiber/fiber/v2"
)

type ApiHandler struct {
    lbStore store.LeaderboardStore
}

type getLeaderboardResponse struct {
    Leaderboard []*custometype.Leaderboard
}

type postLeaderboardRequest struct {
    Pseudo string `json:"pseudo"`
    Score float64 `json:"score"`
}

func NewApiHandler(lbs store.LeaderboardStore) *ApiHandler {
    return &ApiHandler{ lbStore: lbs }
}

func (h *ApiHandler) HandleGetGameLeaderboard(c *fiber.Ctx) error {
    gameUUID := c.Params("uuid")
    top, _ := c.ParamsInt("top")

    if top == 0 {
        leaderboard, err := h.lbStore.GetLeaderboardByGameUUID(gameUUID)

        if err != nil {
            return c.Status(500).JSON(map[string]string{"message": "database error"})
        }

        return c.JSON(leaderboard)
    }

    leaderboard, err := h.lbStore.GetTopPlayerFromGameUUID(gameUUID, top)

    if err != nil {
        return c.Status(500).JSON(map[string]string{"message": "database error"})
    }

    response := &getLeaderboardResponse{Leaderboard: leaderboard}

    return c.JSON(response)
}

func (h *ApiHandler) HandlePostLeaderboard(c *fiber.Ctx) error {
    var leaderboard *custometype.Leaderboard
    gameUUID := c.Params("uuid")

    err := c.BodyParser(&leaderboard)

    if err != nil {
        return c.Status(400).JSON(map[string]string{"message": "bad request"})
    }
    leaderboard.Game_UUID = gameUUID

    if leaderboard.User_UUID == "" {
        return c.Status(400).JSON(map[string]string{"message": "pseudo is empty"})
    }

    err = h.lbStore.AddLeaderboard(leaderboard)

    if err != nil {
        return c.Status(500).JSON(map[string]string{"message": "database error"})
    }

    return c.Status(201).JSON(map[string]string{"message": "score added to database"})
}
