package registercontroller

import (
	"go-web-native/entities"
	"go-web-native/models/registermodel"
	"net/http"
	"text/template"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/register/index.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var user entities.User

		user.Name = r.FormValue("name")
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")

		ok := registermodel.Create(user)
		if !ok {
			temp, _ := template.ParseFiles("views/register/index.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}