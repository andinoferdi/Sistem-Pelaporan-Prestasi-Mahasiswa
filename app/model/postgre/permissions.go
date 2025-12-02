package model

type Permission struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Description string `json:"description"`
}

type CreatePermissionRequest struct {
	Name        string `json:"name" validate:"required"`
	Resource    string `json:"resource" validate:"required"`
	Action      string `json:"action" validate:"required"`
	Description string `json:"description"`
}

type UpdatePermissionRequest struct {
	Name        string `json:"name" validate:"required"`
	Resource    string `json:"resource" validate:"required"`
	Action      string `json:"action" validate:"required"`
	Description string `json:"description"`
}

type GetAllPermissionsResponse struct {
	Status string       `json:"status"`
	Data   []Permission `json:"data"`
}

type GetPermissionByIDResponse struct {
	Status string     `json:"status"`
	Data   Permission `json:"data"`
}

type CreatePermissionResponse struct {
	Status string     `json:"status"`
	Data   Permission `json:"data"`
}

type UpdatePermissionResponse struct {
	Status string     `json:"status"`
	Data   Permission `json:"data"`
}

type DeletePermissionResponse struct {
	Status string `json:"status"`
}
