package model

import "time"

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	FullName     string    `json:"full_name"`
	RoleID       string    `json:"role_id"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	RoleID   string `json:"role_id" validate:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	RoleID   string `json:"role_id" validate:"required"`
	IsActive *bool  `json:"is_active"`
}

type GetAllUsersResponse struct {
	Status string `json:"status"`
	Data   []User `json:"data"`
}

type GetUserByIDResponse struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
}

type CreateUserResponse struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
}

type UpdateUserResponse struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
}

type DeleteUserResponse struct {
	Status string `json:"status"`
}

