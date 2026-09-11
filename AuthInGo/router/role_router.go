package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	RoleController *controllers.RoleController
}

func NewRoleRouter(_roleController *controllers.RoleController) Router{
	return &RoleRouter{
		RoleController: _roleController,
	}
}

func (rr *RoleRouter) Register(r chi.Router){
	// Role CRUD operations
	r.Get("/roles/{id}", rr.RoleController.GetRoleById)
	r.Get("/roles", rr.RoleController.GetAllRoles)
}