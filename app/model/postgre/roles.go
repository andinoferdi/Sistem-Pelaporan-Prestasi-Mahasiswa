package model

import "time"

// role model
type Role struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// create role
type CreateRoleRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

// update role
type UpdateRoleRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

// response semua role
type GetAllRolesResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []Role `json:"data"`
}

// response role by id
type GetRoleByIDResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Role   `json:"data"`
}

// response create role
type CreateRoleResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Role   `json:"data"`
}

// response update role
type UpdateRoleResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Role   `json:"data"`
}

// response delete role
type DeleteRoleResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
