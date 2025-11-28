package postgre

import (
	"database/sql"
	"os"
	model "sistem-pelaporan-prestasi-mahasiswa/app/model/postgre"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//jwt claims
type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	RoleID string `json:"role_id"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(getJWTSecret())

//ambil jwt secret
func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "sistem-pelaporan-prestasi-mahasiswa-jwt-secret-key-minimum-32-characters-long-for-production-security"
	}
	return secret
}

//generate token
func GenerateToken(user model.User) (string, error) {
	claims := JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "sistem-pelaporan-prestasi-mahasiswa-api",
			Subject:   "user-authentication",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

//generate refresh token
func GenerateRefreshToken(user model.User) (string, error) {
	claims := JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "sistem-pelaporan-prestasi-mahasiswa-api",
			Subject:   "user-refresh-token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

//validate token
func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}

//extract token dari header
func ExtractTokenFromHeader(authHeader string) string {
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		return authHeader[7:]
	}
	return ""
}

//cek permission user
func CheckUserPermission(db *sql.DB, userID string, permission string) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM role_permissions rp
		INNER JOIN permissions p ON rp.permission_id = p.id
		INNER JOIN users u ON u.role_id = rp.role_id
		WHERE u.id = $1 AND p.name = $2
	`

	var hasPermission bool
	err := db.QueryRow(query, userID, permission).Scan(&hasPermission)
	if err != nil {
		return false, err
	}

	return hasPermission, nil
}

