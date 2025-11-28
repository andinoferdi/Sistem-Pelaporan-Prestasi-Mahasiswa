package helper

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//response success
func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

//response error
func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}

//response validation error
func ValidationErrorResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusBadRequest, message)
}

//response unauthorized
func UnauthorizedResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusUnauthorized, message)
}

//response not found
func NotFoundResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusNotFound, message)
}

//response internal server error
func InternalServerErrorResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusInternalServerError, message)
}

//response forbidden
func ForbiddenResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusForbidden, message)
}

//handle database error
func HandleDatabaseError(c *fiber.Ctx, err error) error {
	if err == sql.ErrNoRows {
		return NotFoundResponse(c, "Data tidak ditemukan di database.")
	}
	return InternalServerErrorResponse(c, "Error mengakses database. Detail: "+err.Error())
}

//parse uuid
func ParseUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

//cek valid uuid
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

//ambil query int
func GetQueryInt(c *fiber.Ctx, key string, defaultValue int) int {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

//ambil query string
func GetQueryString(c *fiber.Ctx, key string, defaultValue string) string {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

//ambil user id dari context
func GetUserIDFromContext(c *fiber.Ctx) (string, bool) {
	userID, ok := c.Locals("user_id").(string)
	return userID, ok
}

//ambil email dari context
func GetEmailFromContext(c *fiber.Ctx) (string, bool) {
	email, ok := c.Locals("email").(string)
	return email, ok
}

//ambil role id dari context
func GetRoleIDFromContext(c *fiber.Ctx) (string, bool) {
	roleID, ok := c.Locals("role_id").(string)
	return roleID, ok
}

//validasi pagination
func ValidatePagination(page, limit int) (int, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	return page, limit
}

//hitung offset
func CalculateOffset(page, limit int) int {
	return (page - 1) * limit
}

//sanitize search
func SanitizeSearch(search string) string {
	search = strings.TrimSpace(search)
	search = strings.ReplaceAll(search, "%", "")
	search = strings.ReplaceAll(search, "_", "")
	return search
}

//cek string kosong
func IsEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}

