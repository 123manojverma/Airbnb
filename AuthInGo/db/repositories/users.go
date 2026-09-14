package db

import (
	"AuthInGo/dto"
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create(user *dto.CreateUserRequestDTO) (*models.User,error)
	GetById(id int64) (*models.User,error)
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
	GetByEmail(*dto.LoginUserRequestDTO) (*models.User,error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) Create(payload *dto.CreateUserRequestDTO) (*models.User,error) {

	query:="INSERT INTO users (username, email,password) VALUES (?, ?,?)"

	result, err := u.db.Exec(
		query,
		payload.Username,
		payload.Email,
		payload.Password,
	)

	if err != nil {
		fmt.Println("Error creating user:",err)
		return nil,err
	}

	lastInsertID, rowErr := result.LastInsertId()
	if rowErr != nil {
		fmt.Println("Error getting last insert ID:", rowErr)
		return nil, rowErr
	}

	user := &models.User{
		Id:       lastInsertID,
		Username: payload.Username,
		Email:    payload.Email,
	}

	fmt.Println("User created successfully:",user)
	
	return user,nil
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
	query := "SELECT id, username, email, created_at, updated_at FROM users"
	rows, err := u.db.Query(query)
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed after processing

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Created_at, &user.Updated_at); err != nil {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) DeleteById(id int64) error{
	query := "DELETE FROM users WHERE id = ?"
	result, err := u.db.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting user:", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()
	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}
	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not deleted")
		return nil
	}
	fmt.Println("User deleted successfully, rows affected:", rowsAffected)
	return nil
}

func (u *UserRepositoryImpl) GetByEmail(payload *dto.LoginUserRequestDTO) (*models.User,error) {
	query:="SELECT id,email,password FROM users WHERE email = ?"

	row := u.db.QueryRow(query, payload.Email)

	user:=&models.User{}

	err:=row.Scan(&user.Id,&user.Email,&user.Password)

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