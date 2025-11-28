package service

import (
	"database/sql"
	model "sistem-pelaporan-prestasi-mahasiswa/app/model/postgre"
	repository "sistem-pelaporan-prestasi-mahasiswa/app/repository/postgre"
	utilspostgre "sistem-pelaporan-prestasi-mahasiswa/utils/postgre"

	"github.com/gofiber/fiber/v2"
)

// login service
func LoginService(c *fiber.Ctx, db *sql.DB) error {
	var req model.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Format request body tidak valid. Pastikan JSON format benar. Detail: " + err.Error(),
		})
	}

	if req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Username/email dan password wajib diisi.",
		})
	}

	user, err := repository.GetUserByUsername(db, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			user, err = repository.GetUserByEmail(db, req.Username)
			if err != nil {
				if err == sql.ErrNoRows {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"success": false,
						"message": "Username/email atau password tidak valid.",
					})
				}
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"success": false,
					"message": "Error mengambil data user dari database. Detail: " + err.Error(),
				})
			}
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Error mengambil data user dari database. Detail: " + err.Error(),
			})
		}
	}

	if !user.IsActive {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Akun Anda tidak aktif. Silakan hubungi administrator.",
		})
	}

	if !utilspostgre.CheckPassword(req.Password, user.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Username/email atau password tidak valid.",
		})
	}

	roleName, err := repository.GetRoleNameByID(db, user.RoleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error mengambil data role dari database. Detail: " + err.Error(),
		})
	}

	permissions, err := repository.GetUserPermissions(db, user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error mengambil data permissions dari database. Detail: " + err.Error(),
		})
	}

	token, err := utilspostgre.GenerateToken(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error generating token. Detail: " + err.Error(),
		})
	}

	refreshToken, err := utilspostgre.GenerateRefreshToken(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error generating refresh token. Detail: " + err.Error(),
		})
	}

	response := model.LoginResponse{
		Status: "success",
		Data: struct {
			Token        string                  `json:"token"`
			RefreshToken string                  `json:"refreshToken"`
			User         model.LoginUserResponse `json:"user"`
		}{
			Token:        token,
			RefreshToken: refreshToken,
			User: model.LoginUserResponse{
				ID:          user.ID,
				Username:    user.Username,
				Email:       user.Email,
				FullName:    user.FullName,
				Role:        roleName,
				Permissions: permissions,
			},
		},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// ambil profile user
func GetProfileService(c *fiber.Ctx, db *sql.DB) error {
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "User ID tidak ditemukan. Silakan login ulang.",
		})
	}

	user, err := repository.GetUserByID(db, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Data user tidak ditemukan di database.",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error mengambil data user dari database. Detail: " + err.Error(),
		})
	}

	response := model.GetProfileResponse{
		Success: true,
		Message: "Data profile berhasil diambil.",
		Data: struct {
			UserID   string `json:"user_id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			FullName string `json:"full_name"`
			RoleID   string `json:"role_id"`
		}{
			UserID:   user.ID,
			Username: user.Username,
			Email:    user.Email,
			FullName: user.FullName,
			RoleID:   user.RoleID,
		},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// refresh token service
func RefreshTokenService(c *fiber.Ctx, db *sql.DB) error {
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "User ID tidak ditemukan. Silakan login ulang.",
		})
	}

	user, err := repository.GetUserByID(db, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Data user tidak ditemukan di database.",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error mengambil data user dari database. Detail: " + err.Error(),
		})
	}

	if !user.IsActive {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Akun Anda tidak aktif. Silakan hubungi administrator.",
		})
	}

	token, err := utilspostgre.GenerateToken(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error generating token. Detail: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Token berhasil di-refresh.",
		"data": fiber.Map{
			"token": token,
		},
	})
}

// logout service
func LogoutService(c *fiber.Ctx, db *sql.DB) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Logout berhasil.",
	})
}
