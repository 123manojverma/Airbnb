package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router{
	return &UserRouter{
		UserController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router){
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("user", "admin")).Get("/profile",ur.UserController.GetUserById)
	r.With(middlewares.UserCreateRequestValidator).Post("/signup",ur.UserController.Create)
	r.With(middlewares.UserLoginRequestValidator).Post("/login",ur.UserController.LoginUser)
}