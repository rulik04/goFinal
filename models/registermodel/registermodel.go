package registermodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func Create(user entities.User) bool {
	result, err := config.DB.Exec(`
		INSERT INTO users (name, email, password) 
		VALUE (?, ?, ?)`,
		user.Name,
		user.Email,
		user.Password,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}