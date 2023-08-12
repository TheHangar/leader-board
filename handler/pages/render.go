package handler

import "github.com/gofiber/fiber/v2"

type homeData struct {
    PageTitle string
}

type loginData struct {
    PageTitle string
}

type RenderHandler struct {}

func NewRenderHandler() *RenderHandler {
    return &RenderHandler{}
}

func (h *RenderHandler) GetHome(c *fiber.Ctx) error {
    data := &homeData{ PageTitle: "Leaderboard" }
    return c.Render("index", data, "layouts/main")
}

func (h *RenderHandler) GetLogin(c *fiber.Ctx) error {
    data := &loginData{ PageTitle: "Leaderboard" }
    return c.Render("login", data, "layouts/login")
}
