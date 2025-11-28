package model

import "time"

//student model
type Student struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	StudentID    string    `json:"student_id"`
	ProgramStudy string    `json:"program_study"`
	AcademicYear string    `json:"academic_year"`
	AdvisorID    string    `json:"advisor_id"`
	CreatedAt    time.Time `json:"created_at"`
}

//create student
type CreateStudentRequest struct {
	UserID       string `json:"user_id" validate:"required"`
	StudentID    string `json:"student_id" validate:"required"`
	ProgramStudy string `json:"program_study"`
	AcademicYear string `json:"academic_year"`
	AdvisorID    string `json:"advisor_id"`
}

//update student
type UpdateStudentRequest struct {
	StudentID    string `json:"student_id" validate:"required"`
	ProgramStudy string `json:"program_study"`
	AcademicYear string `json:"academic_year"`
	AdvisorID    string `json:"advisor_id"`
}

//response semua student
type GetAllStudentsResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Student `json:"data"`
}

//response student by id
type GetStudentByIDResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    Student `json:"data"`
}

//response create student
type CreateStudentResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    Student `json:"data"`
}

//response update student
type UpdateStudentResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    Student `json:"data"`
}

//response delete student
type DeleteStudentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
