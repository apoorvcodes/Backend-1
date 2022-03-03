package routes

import "github.com/gominima/minima"

func Router() *minima.Router {
	rt := minima.NewRouter()
	//assign methods to routes
	rt.Get("/test/:id", TestRoute)
	return rt
}
