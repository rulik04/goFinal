package loginmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func Verify(user entities.User) bool {
	row := config.DB.QueryRow(`SELECT name, email, password FROM users WHERE email = ? AND password = ?`, user.Email, user.Password)

	var u entities.User

	if err := row.Scan(&u.Name, &u.Email, &u.Password); err != nil {
		return false
	}

	return true
}