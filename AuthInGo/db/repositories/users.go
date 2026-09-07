package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create(user *models.User) error
	GetById(id int64) (*models.User,error)
	GetAll() ([]*models.User,error)
	DeleteById() error
	GetByEmail(user *models.User) (*models.User,error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) Create(user *models.User) error {

	query:="INSERT INTO users (username, email,password) VALUES (?, ?,?)"

	result, err := u.db.Exec(
		query,
		user.Username,
		user.Email,
		user.Password,
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

func (u *UserRepositoryImpl) GetById(id int64) (*models.User,error) {
	fmt.Println("Fetching user in UserRepository")

	// Step 1: Prepare the query
	query:="SELECT id,username,email,password,created_at,updated_at FROM users WHERE id = ?"

	// Step 2: Execute the query
	row := u.db.QueryRow(query, id)

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

func (u *UserRepositoryImpl) GetByEmail(data *models.User) (*models.User,error) {
	query:="SELECT id,username,email,password,created_at,updated_at FROM users WHERE email = ?"

	row := u.db.QueryRow(query, data.Email)

	user:=&models.User{}

	err:=row.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.Created_at,&user.Updated_at)

	if err!=nil{
		if err==sql.ErrNoRows{
			fmt.Println("No user found with the given Email")
			return nil,err
		}else{
			fmt.Println("Error scanning user:",err)
			return nil,err
		}
	}

	fmt.Println("User fetched successfully:",user)

	return user,nil
}