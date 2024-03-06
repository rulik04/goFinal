package logincontroller

import (
	"go-web-native/entities"
	"go-web-native/models/loginmodel"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

func Verify(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/login/index.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var user entities.User
		session, err := store.Get(r, "secret")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user.Name = r.FormValue("name")
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		ok := loginmodel.Verify(user)
		if !ok {
			temp, _ := template.ParseFiles("views/login/index.html")
			temp.Execute(w, nil)
		}
		session.Values["authenticated"] = true
		session.Save(r, w)
	
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}