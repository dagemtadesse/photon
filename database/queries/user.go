package queries

import (
	"photon/database"
	"photon/model"
)

func CreateUserCreds(user *model.Credential) error {

	conn := database.GetDBInstance()
	query := `
		INSERT INTO user_credentials(email, password) 
		VALUES ($1, $2)
	`
	_, err := conn.Exec(query, user.Email, user.Password)

	return err
}

func GetUserCreds(email string) (model.Credential, error) {
	var user model.Credential

	conn := database.GetDBInstance()
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
