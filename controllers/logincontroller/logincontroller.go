package logincontroller

import (
	"go-web-native/entities"
	"go-web-native/models/loginmodel"
	"net/http"
	"text/template"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/login/index.html")
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

		ok := loginmodel.Verify(user)
		if !ok {
			temp, _ := template.ParseFiles("views/login/index.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}