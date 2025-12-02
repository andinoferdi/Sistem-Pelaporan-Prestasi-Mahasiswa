package model

import "time"

type Role struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateRoleRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type UpdateRoleRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type GetAllRolesResponse struct {
	Status string `json:"status"`
	Data   []Role `json:"data"`
}

type GetRoleByIDResponse struct {
	Status string `json:"status"`
	Data   Role   `json:"data"`
}

type CreateRoleResponse struct {
	Status string `json:"status"`
	Data   Role   `json:"data"`
}

type UpdateRoleResponse struct {
	Status string `json:"status"`
	Data   Role   `json:"data"`
}

type DeleteRoleResponse struct {
	Status string `json:"status"`
}

