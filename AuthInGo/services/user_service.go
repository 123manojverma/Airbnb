package services

import (
	config "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById(int64) (*models.User,error)
	Create(*dto.CreateUserRequestDTO) (*models.User,error)
	LoginUser(*dto.LoginUserRequestDTO) (string, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById(id int64) (*models.User,error) {
	fmt.Println("Creating user in UserService")
	user,err:=u.userRepository.GetById(id)
	if err!=nil{
		return nil,err
	}
	return user,nil
}

func(u *UserServiceImpl) Create(payload *dto.CreateUserRequestDTO) (*models.User,error) {
	hashPassword,err:=utils.HashPassword(payload.Password)
	if err!=nil{
		return nil,err
	}
	payload.Password=hashPassword
	user,err:=u.userRepository.Create(payload)
	if err!=nil{
		fmt.Println("Error creating user:",err)
		return nil,err
	}

	return user,nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string,error) {
	user1,err:=u.userRepository.GetByEmail(payload)
	if err!=nil {
		return "",err
	}
	response:=utils.CheckPasswordHash(payload.Password,user1.Password)
	fmt.Println("Login response:",response)

	key:=jwt.MapClaims{
		"email":user1.Email,
		"id":user1.Id,
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,key)

	tokenString,err:=token.SignedString([]byte(config.GetString("JWT_SECRET","TOKEN")))

	if err!=nil{
		fmt.Println("Error signing tokens:",err)
		return "",err
	}

	fmt.Println("JWT Tokens:",tokenString)

	return tokenString,nil
}