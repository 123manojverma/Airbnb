package db

import (
	"AuthInGo/models"
	"database/sql"
)

type RolePermissionRepository interface {
	GetRolePermissionById(id int64) (*models.RolePermission, error)
	GetRolePermissionByRoleId(roleId int64) ([]*models.	RolePermission, error)
	AddPermissionToRole(roleId int64,permissionId int64) (*models.RolePermission,error)
	RemovePermissionFromRole(roleId int64,permissionId int64) error
	GetAllRolePermissions() ([]*models.RolePermission,error)
}

type RolePermissionRepositoryImpl struct {
	db *sql.DB
}

func NewRolePermissionRepository(db *sql.DB) RolePermissionRepository {
	return &RolePermissionRepositoryImpl{
		db: db,
	}
}

func (r *RolePermissionRepositoryImpl) GetRolePermissionById(id int64) (*models.RolePermission, error) {
	query := `SELECT id,role_id,permission_id,created_at,updated_at FROM role_permissions WHERE id = ?`
	rolePermission := &models.RolePermission{}
	row := r.db.QueryRow(query, id)
	err := row.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.CreatedAt, &rolePermission.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return rolePermission, nil
}

func (r *RolePermissionRepositoryImpl) GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error) {
	query := `SELECT id,role_id,permission_id,created_at,updated_at FROM role_permissions WHERE role_id = ?`
	rolesPermissions := []*models.RolePermission{}
	rows, err := r.db.Query(query, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rolePermission := &models.RolePermission{}
		err := rows.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.CreatedAt, &rolePermission.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rolesPermissions = append(rolesPermissions, rolePermission)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rolesPermissions, nil
}

func (r *RolePermissionRepositoryImpl) AddPermissionToRole(roleId int64,permissionId int64) (*models.RolePermission,error) {
	query := `INSERT INTO role_permissions(role_id,permission_id) VALUES(?,?)`
	result, err := r.db.Exec(query, roleId, permissionId)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &models.RolePermission{
		Id:           id,
		RoleId:       roleId,
		PermissionId: permissionId,
	}, nil
}

func (r *RolePermissionRepositoryImpl) RemovePermissionFromRole(roleId int64,permissionId int64) error {
	query := `DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?`
	_, err := r.db.Exec(query, roleId, permissionId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RolePermissionRepositoryImpl) GetAllRolePermissions() ([]*models.RolePermission,error) {
	query := `SELECT id,role_id,permission_id,created_at,updated_at FROM role_permissions`
	rolesPermissions := []*models.RolePermission{}
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rolePermission := &models.RolePermission{}
		err := rows.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.CreatedAt, &rolePermission.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rolesPermissions = append(rolesPermissions, rolePermission)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rolesPermissions, nil
}