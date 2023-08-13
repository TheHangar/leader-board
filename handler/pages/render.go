package pages

import "github.com/gofiber/fiber/v2"

type pageData struct {
    PageTitle string
}
type leaderboardData struct {
    PageTitle string
    GameUUID string
}

type RenderHandler struct {}

func NewRenderHandler() *RenderHandler {
    return &RenderHandler{}
}

func (h *RenderHandler) GetHome(c *fiber.Ctx) error {
    data := &pageData{ PageTitle: "Dashboard" }
    return c.Render("index", data, "layouts/main")
}

func (h *RenderHandler) GetLogin(c *fiber.Ctx) error {
    data := &pageData{ PageTitle: "Login" }
    return c.Render("login", data, "layouts/login")
}

func (h *RenderHandler) GetGame(c *fiber.Ctx) error {
    gameUUID := c.Params("id")
    data := &leaderboardData{ PageTitle: "My game", GameUUID: gameUUID }
    return c.Render("leaderboard", data, "layouts/main")
}
