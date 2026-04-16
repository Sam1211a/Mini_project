package handlers

import (
	"html/template"
	"mini_project/models"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/dashboard.html"))
var User1 models.Users

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cookie, err := r.Cookie("user_email")
		if err != nil {
			http.Error(w, "Cookies not Found", 400)
			return
		}
		Email := cookie.Value
		err = models.Db.QueryRow("select name,phone,country from information where email=$1", Email).Scan(&User1.Name, &User1.Phone, &User1.Country)

		tmpl.Execute(w, map[string]string{
			"Name":    User1.Name,
			"Email":   Email,
			"Country": User1.Country,
			"Phn":     User1.Phone,
		})
		// http.ServeFile(w, r, "templates/profile.html")
	}

}
