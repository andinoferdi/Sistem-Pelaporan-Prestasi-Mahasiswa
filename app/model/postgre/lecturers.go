package model

import "time"

//lecturer model
type Lecturer struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	LecturerID string    `json:"lecturer_id"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"created_at"`
}

//create lecturer
type CreateLecturerRequest struct {
	UserID     string `json:"user_id" validate:"required"`
	LecturerID string `json:"lecturer_id" validate:"required"`
	Department string `json:"department"`
}

//update lecturer
type UpdateLecturerRequest struct {
	LecturerID string `json:"lecturer_id" validate:"required"`
	Department string `json:"department"`
}

//response semua lecturer
type GetAllLecturersResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    []Lecturer `json:"data"`
}

//response lecturer by id
type GetLecturerByIDResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    Lecturer `json:"data"`
}

//response create lecturer
type CreateLecturerResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    Lecturer `json:"data"`
}

//response update lecturer
type UpdateLecturerResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    Lecturer `json:"data"`
}

//response delete lecturer
type DeleteLecturerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
