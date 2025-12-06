package repository

import (
	"database/sql"
	model "sistem-pelaporan-prestasi-mahasiswa/app/model/postgre"
)

func GetUserByEmail(db *sql.DB, email string) (*model.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.password_hash, u.full_name, 
		       u.role_id, u.is_active, u.created_at, u.updated_at
		FROM users u
		WHERE u.email = $1
	`

	user := new(model.User)
	err := db.QueryRow(query, email).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash,
		&user.FullName, &user.RoleID, &user.IsActive,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByID(db *sql.DB, id string) (*model.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.password_hash, u.full_name, 
		       u.role_id, u.is_active, u.created_at, u.updated_at
		FROM users u
		WHERE u.id = $1
	`

	user := new(model.User)
	err := db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash,
		&user.FullName, &user.RoleID, &user.IsActive,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByUsernameOrEmail(db *sql.DB, usernameOrEmail string) (*model.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.password_hash, u.full_name, 
		       u.role_id, u.is_active, u.created_at, u.updated_at
		FROM users u
		WHERE u.username = $1 OR u.email = $1
	`

	user := new(model.User)
	err := db.QueryRow(query, usernameOrEmail).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash,
		&user.FullName, &user.RoleID, &user.IsActive,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

type RefreshToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt string
	CreatedAt string
}

func SaveRefreshToken(db *sql.DB, userID string, token string, expiresAt string) error {
	query := `
		INSERT INTO refresh_tokens (user_id, token, expires_at)
		VALUES ($1, $2, $3)
	`
	_, err := db.Exec(query, userID, token, expiresAt)
	return err
}

func GetRefreshToken(db *sql.DB, token string) (*RefreshToken, error) {
	query := `
		SELECT id, user_id, token, expires_at, created_at
		FROM refresh_tokens
		WHERE token = $1 AND expires_at > NOW()
	`
	rt := new(RefreshToken)
	err := db.QueryRow(query, token).Scan(
		&rt.ID, &rt.UserID, &rt.Token, &rt.ExpiresAt, &rt.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return rt, nil
}

func DeleteRefreshToken(db *sql.DB, token string) error {
	query := `DELETE FROM refresh_tokens WHERE token = $1`
	_, err := db.Exec(query, token)
	return err
}

func DeleteUserRefreshTokens(db *sql.DB, userID string) error {
	query := `DELETE FROM refresh_tokens WHERE user_id = $1`
	_, err := db.Exec(query, userID)
	return err
}

func GetUserPermissions(db *sql.DB, userID string) ([]string, error) {
	query := `
		SELECT p.name
		FROM role_permissions rp
		INNER JOIN permissions p ON rp.permission_id = p.id
		INNER JOIN users u ON u.role_id = rp.role_id
		WHERE u.id = $1
		ORDER BY p.name
	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []string
	for rows.Next() {
		var perm string
		if err := rows.Scan(&perm); err != nil {
			return nil, err
		}
		permissions = append(permissions, perm)
	}

	return permissions, nil
}

func GetRoleName(db *sql.DB, roleID string) (string, error) {
	query := `SELECT name FROM roles WHERE id = $1`
	var roleName string
	err := db.QueryRow(query, roleID).Scan(&roleName)
	if err != nil {
		return "", err
	}
	return roleName, nil
}

func GetLecturerByUserID(db *sql.DB, userID string) (*model.Lecturer, error) {
	query := `
		SELECT l.id, l.user_id, l.lecturer_id, l.department, l.created_at
		FROM lecturers l
		WHERE l.user_id = $1
	`

	lecturer := new(model.Lecturer)
	err := db.QueryRow(query, userID).Scan(
		&lecturer.ID, &lecturer.UserID, &lecturer.LecturerID,
		&lecturer.Department, &lecturer.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return lecturer, nil
}

