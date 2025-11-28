package model

import "time"

// user model
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

// create user
type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	RoleID   string `json:"role_id" validate:"required"`
}

// update user
type UpdateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	RoleID   string `json:"role_id" validate:"required"`
	IsActive *bool  `json:"is_active"`
}

// response semua user
type GetAllUsersResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

// response user by id
type GetUserByIDResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

// response create user
type CreateUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

// response update user
type UpdateUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

// response delete user
type DeleteUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
