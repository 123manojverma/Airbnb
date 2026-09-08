package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/services"
	"AuthInGo/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserById called in UserController")
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	user, err := uc.UserService.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(
			w, "Invalid request body", http.StatusBadRequest,
		)
		return
	}

	err = uc.UserService.Create(&user)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created Successfully"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginUser called in UserController")

	var payload dto.LoginUserRequestDTO

	if jsonErr:= utils.ReadJsonBody(r,&payload); jsonErr != nil {
		http.Error(
			w, "Invalid request body", http.StatusBadRequest,
		)
		return
	}

	if validationErr:=utils.Validator.Struct(payload); validationErr!=nil{
		w.Write([]byte("Invalid input data"))
		fmt.Println("Validation Err",validationErr)
		return
	}

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Failed to login user",err)
		return
	}
	
	w.Header().Set("Authorization", "Bearer "+jwtToken)

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"User logged in successfully",jwtToken)
}
