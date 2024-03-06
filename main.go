package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/logincontroller"
	"go-web-native/controllers/productcontroller"
	"go-web-native/controllers/registercontroller"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("rulan-aisultan-golang-final-project"))

func main() {
	// Database connection
	config.ConnectDB()
	// Routes
	// 1. Homepage
	http.HandleFunc("/", CheckSession(homecontroller.Welcome))

	// 2. Category
	http.HandleFunc("/categories", CheckSession(categorycontroller.Index))
	http.HandleFunc("/categories/add", CheckSession(categorycontroller.Add))
	http.HandleFunc("/categories/edit", CheckSession(categorycontroller.Edit))
	http.HandleFunc("/categories/delete", CheckSession(categorycontroller.Delete))

	// 3. Products
	http.HandleFunc("/products", CheckSession(productcontroller.Index))
	http.HandleFunc("/products/add", CheckSession(productcontroller.Add))
	http.HandleFunc("/products/detail", CheckSession(productcontroller.Detail))
	http.HandleFunc("/products/edit", CheckSession(productcontroller.Edit))
	http.HandleFunc("/products/delete", CheckSession(productcontroller.Delete))

	// 4. Login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        logincontroller.Verify(w, r, store)
    })
	http.HandleFunc("/logout", logoutHandler)

	// 5. Register
	http.HandleFunc("/register", registercontroller.Create)

	// Run server
	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CheckSession(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "secret")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		handler(w, r)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "secret")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = false
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
}
