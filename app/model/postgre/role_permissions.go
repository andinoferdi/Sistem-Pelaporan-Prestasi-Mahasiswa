package model

//role permission model
type RolePermission struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

//create role permission
type CreateRolePermissionRequest struct {
	RoleID       string `json:"role_id" validate:"required"`
	PermissionID string `json:"permission_id" validate:"required"`
}

//response semua role permission
type GetAllRolePermissionsResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    []RolePermission `json:"data"`
}

//response create role permission
type CreateRolePermissionResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    RolePermission `json:"data"`
}

//response delete role permission
type DeleteRolePermissionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
