package handlers

import (
	"html/template"
	"net/http"
)

var tmp = template.Must(template.ParseFiles("templates/edit_profile.html"))

func EditProfileHandle(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_email")
	if err != nil {
		panic(err)
	}
	Email := cookie.Value
	if r.Method == "GET" {
		tmp.Execute(w, map[string]string{
			"Name":    "",
			"Email":   Email,
			"Country": "",
			"Phn":     "",
		})

	}
}

// http.ServeFile(w, r, "templates/edit_profile.html")
