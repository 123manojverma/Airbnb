package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	GetUserById(id int64) (*models.User,error)
	Create(user *models.User) error
	LoginUser(user *models.User) (token string, err error)
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

func (u *UserServiceImpl) LoginUser(user *models.User) (token string,err error) {
	user1,err:=u.userRepository.GetByEmail(user)
	if err!=nil {
		return "",err
	}
	response:=utils.CheckPasswordHash(user.Password,user1.Password)
	fmt.Println("Login response:",response)
	return user1.Password,nil
}