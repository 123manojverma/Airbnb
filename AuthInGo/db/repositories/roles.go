package db

import (
	"AuthInGo/models"
	"database/sql"
)

type RoleRepository interface {
	GetRoleById(int64) (*models.Role, error)
	GetRoleByName(string) (*models.Role,error)
	GetAllRoles() ([]*models.Role,error)
	CreateRole(string,string) (*models.Role,error)
	UpdateRole(int64,string,string) (*models.Role,error)
	DeleteRole(int64) (error)	
}

type RoleRepositoryImpl struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: db,
	}
}

func (r *RoleRepositoryImpl) GetRoleById(id int64) (*models.Role, error) {
	query:="SELECT id,name,description,created_at,updated_at FROM roles WHERE id = ?"
	role := &models.Role{}
	row := r.db.QueryRow(query, id)
	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepositoryImpl) GetRoleByName(name string) (*models.Role, error) {
	query:="SELECT id,name,description,created_at,updated_at FROM roles WHERE name = ?"
	role := &models.Role{}
	row := r.db.QueryRow(query, name)
	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepositoryImpl) GetAllRoles() ([]*models.Role, error) {
	query:="SELECT id,name,description,created_at,updated_at FROM roles"
	roles := []*models.Role{}
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		role := &models.Role{}
		err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepositoryImpl) CreateRole(name,description string) (*models.Role, error) {
	query:="INSERT INTO roles(name,description) VALUES(?,?)"
	result, err := r.db.Exec(query, name, description)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &models.Role{
		Id:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (r *RoleRepositoryImpl) UpdateRole(id int64, name string, description string) (*models.Role, error) {
	query:="UPDATE roles SET name=?,description=?,updated_at=CURRENT_TIMESTAMP WHERE id=?"
	_, err := r.db.Exec(query, name, description, id)
	if err != nil {
		return nil, err
	}
	return &models.Role{
		Id:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (r *RoleRepositoryImpl) DeleteRole(id int64) (error) {
	query:="DELETE FROM roles WHERE id=?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}