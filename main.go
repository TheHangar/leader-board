package main

import (
	"flag"

	"github.com/TheHangar/leader-board/handler"
	"github.com/TheHangar/leader-board/handler/api"
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

    if seedErr := store.PostgresSeed(); seedErr != nil {
        panic(seedErr)
    }

    page := pages.NewRenderHandler()
    apiAuth := middleware.NewAuthAPIHandler(postgres.Game)
    userAuth := middleware.NewAuthUserHandler()
    adminHandler := handler.NewAdminHandler(postgres)
    apiHandler := api.NewApiHandler(postgres.Leaderboard)
    
    app := fiber.New(fiber.Config{
        Views: engine,
    })
    apiv1 := app.Group("/api/v1/games/:uuid", apiAuth.VerifyCredentials)
    hm := app.Group("/hm", userAuth.VerifyToken)

    app.Static("/static", "./www/public")

    app.Get("/", userAuth.VerifyToken, page.GetHome)
    app.Get("/login", page.GetLogin)
    app.Get("/games/:id", userAuth.VerifyToken, page.GetGame)
    app.Post("/login", adminHandler.HandlePostLogin, userAuth.CreateToken)
    app.Post("/games", userAuth.VerifyToken, adminHandler.HandlePostGame)

    hm.Get("/games", userAuth.VerifyToken, adminHandler.HandleGetGames)
    hm.Get("/games/:uuid/usage", userAuth.VerifyToken, adminHandler.HandleGetGameUsage)
    hm.Get("/games/:uuid/leaderboard", userAuth.VerifyToken, adminHandler.HandleGetGameLeaderboard)
    hm.Delete("/games/:uuid", userAuth.VerifyToken, adminHandler.HandleDeleteGame)

    apiv1.Get("/leaderboard/:top?", apiHandler.HandleGetGameLeaderboard)
    apiv1.Post("/leaderboard", apiHandler.HandlePostLeaderboard)

    if err := app.Listen(*listenAddr); err != nil {
        panic(err)
    }
}
