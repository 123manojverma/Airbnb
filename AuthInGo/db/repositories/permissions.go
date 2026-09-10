package db

import (
	"AuthInGo/models"
	"database/sql"
)

type PermissionRepository interface {
	GetPermissionById(id int64) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission,error)
	GetAllPermissions() ([]*models.Permission,error)
	CreatePermission(name string,description string,resource string,action string) (*models.Permission,error)
	UpdatePermission(id int64,name string,description string,resource string,action string) (*models.Permission,error)
	DeletePermission(id int64) (error)	
}

type PermissionRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: db,
	}
}

func (r *PermissionRepositoryImpl) GetPermissionById(id int64) (*models.Permission, error) {
	query:="SELECT id,name,description,resource,action,created_at,updated_at FROM permissions WHERE id = ?"
	permission := &models.Permission{}
	row := r.db.QueryRow(query, id)
	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (r *PermissionRepositoryImpl) GetPermissionByName(name string) (*models.Permission, error) {
	query:="SELECT id,name,description,resource,action,created_at,updated_at FROM permissions WHERE name = ?"
	permission := &models.Permission{}
	row := r.db.QueryRow(query, name)
	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (r *PermissionRepositoryImpl) GetAllPermissions() ([]*models.Permission, error) {
	query:="SELECT id,name,description,resource,action,created_at,updated_at FROM permissions"
	permissions := []*models.Permission{}
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		permission := &models.Permission{}
		err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *PermissionRepositoryImpl) CreatePermission(name,description,resource,action string) (*models.Permission,error) {
	query:="INSERT INTO permissions(name,description,resource,action) VALUES(?,?,?,?)"
	result, err := r.db.Exec(query, name, description, resource, action)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &models.Permission{
		Id:          id,
		Name:        name,
		Description: description,
		Resource:    resource,
		Action:      action,
	}, nil
}

func (r *PermissionRepositoryImpl) UpdatePermission(id int64,name,description,resource,action string) (*models.Permission,error) {
	query:="UPDATE permissions SET name=?,description=?,resource=?,action=?,updated_at=CURRENT_TIMESTAMP WHERE id=?"
	_, err := r.db.Exec(query, name, description, resource, action, id)
	if err != nil {
		return nil, err
	}
	return &models.Permission{
		Id:          id,
		Name:        name,
		Description: description,
		Resource:    resource,
		Action:      action,
	}, nil
}

func (r *PermissionRepositoryImpl) DeletePermission(id int64) (error) {
	query:="DELETE FROM permissions WHERE id=?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}