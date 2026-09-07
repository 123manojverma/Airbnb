package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create() error
	GetById() (*models.User,error)
	GetAll() ([]*models.User,error)
	DeleteById() error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) Create() error {

	query:="INSERT INTO users (username, email,password) VALUES (?, ?,?)"

	username:="testuser"
	email:="test@gmail.com"
	password:="password123"

	result, err := u.db.Exec(
		query,
		username,
		email,
		password,
	)

	if err != nil {
		fmt.Println("Error inserting user:",err)
		return err
	}

	rowsAffected,rowErr:=result.RowsAffected()

	if rowErr!=nil{
		fmt.Println("Error getting rows affected:",rowErr)
		return rowErr
	}

	if rowsAffected==0{
		fmt.Println("No rows were affected, user not created")
		return  nil
	}

	fmt.Println("User created successfully, rows affected:",rowsAffected)
	
	return nil
}

func (u *UserRepositoryImpl) GetById() (*models.User,error) {
	fmt.Println("Fetching user in UserRepository")

	// Step 1: Prepare the query
	query:="SELECT id,username,email,password,created_at,updated_at FROM users WHERE id = ?"

	// Step 2: Execute the query
	row := u.db.QueryRow(query, 1)

	// Step 3: Process the result
	user:=&models.User{}

	err:=row.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.Created_at,&user.Updated_at)

	if err!=nil{
		if err==sql.ErrNoRows{
			fmt.Println("No user found with the given ID")
			return nil,err
		}else{
			fmt.Println("Error scanning user:",err)
			return nil,err
		}
	}

	// Step 4: Print the user details
	fmt.Println("User fetched successfully:",user)

	return user,nil
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User,error){
	return nil,nil
}

func (u *UserRepositoryImpl) DeleteById() error{
	return nil
}