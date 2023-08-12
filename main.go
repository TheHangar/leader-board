package main

import (
	"flag"

	handler "github.com/TheHangar/leader-board/handler/pages"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
    listenAddr := flag.String("p", ":3000", "Specify server's listening port (default :3000)")
    flag.Parse()

    engine := html.New("./www", ".html")
    page := handler.NewRenderHandler()
    
    app := fiber.New(fiber.Config{
        Views: engine,
    })
    apiv1 := app.Group("/api/v1")

    app.Static("/static", "./www/public")

    app.Get("/", page.GetHome)
    app.Get("/login", page.GetLogin)

    apiv1.Get("/test", func(c *fiber.Ctx) error {
        return c.JSON(map[string]string{"message": "Hello friend."})
    })

    if err := app.Listen(*listenAddr); err != nil {
        panic(err)
    }
}
