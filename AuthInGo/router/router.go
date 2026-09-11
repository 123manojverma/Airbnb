package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"
	"AuthInGo/utils"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Router interface{
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router,RoleRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	// chiRouter.Use(middlewares.RequestLogger) // Middleware for logging requests
	chiRouter.Use(middleware.Logger) // Built-in Chi middleware for logging requests
	
	// chiRouter.Use(middlewares.RateLimitMiddleware)
	chiRouter.Get("/ping",controllers.PingHandler)

	chiRouter.With(middlewares.JWTAuthMiddleware).HandleFunc("/fakestoreservice/*",utils.ProxyToService("https://fakestoreapi.com/","/fakestoreservice"))

	UserRouter.Register(chiRouter)

	RoleRouter.Register(chiRouter)

	return chiRouter
}