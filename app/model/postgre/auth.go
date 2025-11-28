package model

// login
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// user login response
type LoginUserResponse struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	FullName    string   `json:"fullName"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

// response login
type LoginResponse struct {
	Status string `json:"status"`
	Data   struct {
		Token        string            `json:"token"`
		RefreshToken string            `json:"refreshToken"`
		User         LoginUserResponse `json:"user"`
	} `json:"data"`
}

// response profile
type GetProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		FullName string `json:"full_name"`
		RoleID   string `json:"role_id"`
	} `json:"data"`
}
