package main

import (
	"github.com/TeamEvie/API/middlewares"
	"github.com/TeamEvie/API/routes"
	"github.com/gominima/cors"
	"github.com/gominima/minima"
)

func main() {
	app := minima.New()
	crs := cors.New()
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello World")
	})
	app.UseRouter(routes.Router())
	app.Use(middlewares.Logger)
	app.Use(crs.NewCors(cors.Options{
		AllowedOrigins:   []string{"https://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}))
	app.Listen(":3000")
}
