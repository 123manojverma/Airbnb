package services

import (
	config "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById(id int64) (*models.User,error)
	Create(user *models.User) error
	LoginUser(user *models.User) (string, error)
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

func(u *UserServiceImpl) Create(user *models.User) error{
	hashPassword,err:=utils.HashPassword(user.Password)
	if err!=nil{
		return err
	}
	user.Password=hashPassword
	err=u.userRepository.Create(user)
	if err!=nil{
		return err
	}

	return nil
}

func (u *UserServiceImpl) LoginUser(user *models.User) (string,error) {
	user1,err:=u.userRepository.GetByEmail(user)
	if err!=nil {
		return "",err
	}
	response:=utils.CheckPasswordHash(user.Password,user1.Password)
	fmt.Println("Login response:",response)

	payload:=jwt.MapClaims{
		"email":user1.Email,
		"id":user1.Id,
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,payload)

	tokenString,err:=token.SignedString([]byte(config.GetString("JWT_SECRET","TOKEN")))

	if err!=nil{
		fmt.Println("Error signing tokens:",err)
		return "",err
	}

	fmt.Println("JWT Tokens:",tokenString)

	return tokenString,nil
}