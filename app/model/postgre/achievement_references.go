package model

import "time"

//status prestasi
const (
	AchievementStatusDraft     = "draft"
	AchievementStatusSubmitted = "submitted"
	AchievementStatusVerified  = "verified"
	AchievementStatusRejected  = "rejected"
)

//achievement reference model
type AchievementReference struct {
	ID                 string     `json:"id"`
	StudentID          string     `json:"student_id"`
	MongoAchievementID string     `json:"mongo_achievement_id"`
	Status             string     `json:"status"`
	SubmittedAt        *time.Time `json:"submitted_at"`
	VerifiedAt         *time.Time `json:"verified_at"`
	VerifiedBy         *string    `json:"verified_by"`
	RejectionNote      string     `json:"rejection_note"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

//create achievement reference
type CreateAchievementReferenceRequest struct {
	StudentID          string `json:"student_id" validate:"required"`
	MongoAchievementID string `json:"mongo_achievement_id" validate:"required"`
	Status             string `json:"status" validate:"required"`
}

//update achievement reference
type UpdateAchievementReferenceRequest struct {
	Status        string `json:"status" validate:"required"`
	RejectionNote string `json:"rejection_note"`
}

//verify achievement
type VerifyAchievementRequest struct {
	Status        string `json:"status" validate:"required"`
	RejectionNote string `json:"rejection_note"`
}

//response semua achievement reference
type GetAllAchievementReferencesResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    []AchievementReference `json:"data"`
}

//response achievement reference by id
type GetAchievementReferenceByIDResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    AchievementReference `json:"data"`
}

//response create achievement reference
type CreateAchievementReferenceResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    AchievementReference `json:"data"`
}

//response update achievement reference
type UpdateAchievementReferenceResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    AchievementReference `json:"data"`
}

//response delete achievement reference
type DeleteAchievementReferenceResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
