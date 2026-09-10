package db

import (
	"AuthInGo/models"
	"database/sql"
)

type UserRoleRepository interface {
	GetUserRoles(userId int64) ([]*models.Role, error)
	AssignRoleToUser(userId int64,roleId int64) error
	RemoveRoleFromUser(userId int64,roleId int64) error
	GetUserPermissions(userId int64) ([]*models.Permission,error)
	HasPermission(userId int64,permissionName string) (bool,error)
	HasRole(userId int64,roleName string) (bool,error)
}

type UserRoleRepositoryImpl struct {
	db *sql.DB
}

func NewUserRoleRepository(db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: db,
	}
}

func (r *UserRoleRepositoryImpl) GetUserRoles(userId int64) ([]*models.Role,error){
	query:=`SELECT r.id,r.name,r.description,r.created_at,r.updated_at FROM roles r JOIN user_roles ur ON r.id = ur.role_id WHERE ur.user_id = ?`
	roles := []*models.Role{}
	rows, err := r.db.Query(query, userId)
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

func (r *UserRoleRepositoryImpl) AssignRoleToUser(userId int64,roleId int64) error {
	query:=`INSERT INTO user_roles(user_id,role_id) VALUES(?,?)`
	_, err := r.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRoleRepositoryImpl) RemoveRoleFromUser(userId int64,roleId int64) error {
	query:=`DELETE FROM user_roles WHERE user_id = ? AND role_id = ?`
	_, err := r.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRoleRepositoryImpl) GetUserPermissions(userId int64) ([]*models.Permission,error){
	query:=`SELECT p.id,p.name,p.description,p.resource,p.action,p.created_at,p.updated_at FROM permissions p JOIN role_permissions rp ON p.id = rp.permission_id JOIN user_roles ur ON rp.role_id = ur.role_id WHERE ur.user_id = ?`
	permissions := []*models.Permission{}
	rows, err := r.db.Query(query, userId)
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

func (r *UserRoleRepositoryImpl) HasPermission(userId int64,permissionName string) (bool,error){
	query:=`SELECT EXISTS(SELECT 1 FROM permissions p JOIN role_permissions rp ON p.id = rp.permission_id JOIN user_roles ur ON rp.role_id = ur.role_id WHERE ur.user_id = ? AND p.name = ?)`
	var exists bool
	err := r.db.QueryRow(query, userId, permissionName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *UserRoleRepositoryImpl) HasRole(userId int64,roleName string) (bool,error){
	query:=`SELECT EXISTS(SELECT 1 FROM roles r JOIN user_roles ur ON r.id = ur.role_id WHERE ur.user_id = ? AND r.name = ?)`
	var exists bool
	err := r.db.QueryRow(query, userId, roleName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}