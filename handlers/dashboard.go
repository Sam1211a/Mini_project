package handlers

import (
	"html/template"
	"mini_project/models"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/dashboard.html"))

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cookie, err := r.Cookie("user_email")
		if err != nil {
			http.Error(w, "Cookies not Found", 400)
			return
		}
		Email := cookie.Value
		var user1 models.Users
		err = models.Db.QueryRow("select name,phone,country from information where email=$1", Email).Scan(&user1.Name, &user1.Phone, &user1.Country)

		tmpl.Execute(w, map[string]string{
			"Name":    user1.Name,
			"Email":   Email,
			"Country": user1.Country,
			"Phn":     user1.Phone,
		})
		// http.ServeFile(w, r, "templates/profile.html")
	}

}
