package model

//permission model
type Permission struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Description string `json:"description"`
}

//create permission
type CreatePermissionRequest struct {
	Name        string `json:"name" validate:"required"`
	Resource    string `json:"resource" validate:"required"`
	Action      string `json:"action" validate:"required"`
	Description string `json:"description"`
}

//update permission
type UpdatePermissionRequest struct {
	Name        string `json:"name" validate:"required"`
	Resource    string `json:"resource" validate:"required"`
	Action      string `json:"action" validate:"required"`
	Description string `json:"description"`
}

//response semua permission
type GetAllPermissionsResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    []Permission `json:"data"`
}

//response permission by id
type GetPermissionByIDResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    Permission `json:"data"`
}

//response create permission
type CreatePermissionResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    Permission `json:"data"`
}

//response update permission
type UpdatePermissionResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    Permission `json:"data"`
}

//response delete permission
type DeletePermissionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
