package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/samridhi-sahu/pich-booking-system/WithFiber/database"
	"github.com/samridhi-sahu/pich-booking-system/WithFiber/routes"
	"go.uber.org/fx"
)

func start(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := routes.Routes(app)
			go app.Listen(":4000")
			return err
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}

func main() {
	database.AddData()
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	fx.New(
		fx.Provide(func() *fiber.App {
			return app
		}),
		fx.Invoke(
			start,
		),
	).Run()
}
