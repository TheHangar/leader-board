package main

import (
	"flag"

	"github.com/TheHangar/leader-board/handler"
	"github.com/TheHangar/leader-board/handler/pages"
	"github.com/TheHangar/leader-board/middleware"
	"github.com/TheHangar/leader-board/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
    listenAddr := flag.String("p", ":3000", "Specify server's listening port (default :3000)")
    flag.Parse()

    engine := html.New("./www", ".html")
    postgres := store.NewPostgresStore()

    if err := store.PostgresTest(); err != nil {
        panic(err)
    }

    /*
    if seedErr := store.PostgresSeed(); seedErr != nil {
        panic(seedErr)
    }
    */

    page := pages.NewRenderHandler()
    apiAuth := middleware.NewAuthAPIHandler()
    userAuth := middleware.NewAuthUserHandler()
    adminHandler := handler.NewAdminHandler(postgres)
    
    app := fiber.New(fiber.Config{
        Views: engine,
    })
    apiv1 := app.Group("/api/v1", apiAuth.VerifyCredentials)

    app.Static("/static", "./www/public")

    app.Get("/", userAuth.VerifyToken, page.GetHome)
    app.Get("/login", page.GetLogin)
    app.Get("/games", userAuth.VerifyToken, adminHandler.HandleGetGames)
    app.Post("/login", adminHandler.HandlePostLogin, userAuth.CreateToken)
    app.Post("/games", userAuth.VerifyToken, adminHandler.HandlePostGame)

    apiv1.Get("/test", func(c *fiber.Ctx) error {
        return c.JSON(map[string]string{"message": "Hello friend."})
    })

    if err := app.Listen(*listenAddr); err != nil {
        panic(err)
    }
}
