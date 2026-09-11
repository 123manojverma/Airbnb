package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/middlewares"
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
	// userId := r.URL.Query().Get("id")
	userId,ok:=r.Context().Value(middlewares.ContextKeyUserID).(string)
	fmt.Println("User ID:",userId,ok)
	if !ok{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("missing user ID"))
		return
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Invalid user ID format", err)
		return
	}
	user, err := uc.UserService.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "User not found", fmt.Errorf("user with ID %d not found", id))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	payload, ok := r.Context().Value("payload").(dto.CreateUserRequestDTO)
	if !ok {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", fmt.Errorf("missing payload"))
		return
	}

	fmt.Println("Payload received:", payload)

	user, err := uc.UserService.Create(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(
			w,
			http.StatusInternalServerError,
			"Failed to create user",
			err)

		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", user)
	fmt.Println("User created successfully:", user)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginUser called in UserController")

	payload, ok := r.Context().Value("payload").(dto.LoginUserRequestDTO)
	if !ok{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", fmt.Errorf("missing payload"))
		return
	}

	fmt.Println("Payload received:", payload)

	// if validationErr:=utils.Validator.Struct(payload); validationErr!=nil{
	// 	w.Write([]byte("Invalid input data"))
	// 	fmt.Println("Validation Err",validationErr)
	// 	return
	// }

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}

	w.Header().Set("Authorization", "Bearer "+jwtToken)

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}
