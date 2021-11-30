package queries

import (
	"photon/database"
	"photon/model"

	"github.com/google/uuid"
)

func CreateUserCreds(user model.Credential) (uuid.UUID, error) {
	var userId uuid.UUID

	conn := database.GetInstance()
	query := `
		INSERT INTO user_credentials(email, password) 
		VALUES ($1, $2) 
		RETURNING id
	`
	err := conn.
		QueryRow(query, user.Email, user.Password).
		Scan(&userId)

	return userId, err
}

func GetUserCreds(email string) (model.Credential, error) {
	var user model.Credential

	conn := database.GetInstance()
	query := `
		SELECT * FROM user_credentials
		wHERE email = $1
		LIMIT 1
	`

	err := conn.
		QueryRow(query, email).
		Scan(&user.Id, &user.Email, &user.Password)

	return user, err
}
