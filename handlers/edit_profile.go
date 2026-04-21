package handlers

import (
	"html/template"
	"mini_project/models"
	"net/http"
	"strings"
)

var tmp = template.Must(template.ParseFiles("templates/edit_profile.html"))

func EditProfileHandle(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_email")
	if err != nil {
		panic(err)
	}
	models.User_Email = cookie.Value
	if r.Method == "GET" {
		tmp.Execute(w, map[string]string{
			"Name":    User1.Name,
			"Email":   models.User_Email,
			"Country": User1.Country,
			"Phn":     User1.Phone[3:],
		})
	}
	if r.Method == "POST" {
		user := models.Users{}
		user.Name = strings.TrimSpace(r.FormValue("name"))
		user.Country = r.FormValue("country")
		user.Phone = "+88" + r.FormValue("phn")
		if user.Name == "" || len(user.Name) < 4 {
			tmp.Execute(w, map[string]string{
				"ErrName": "Invalid Name",
				"Name":    User1.Name,
				"Email":   models.User_Email,
				"Country": User1.Country,
				"Phn":     User1.Phone[3:],
			})
			return
		}
		if strings.HasPrefix(user.Phone, "+8801") == false || len(user.Phone) != 14 {
			tmp.Execute(w, map[string]string{
				"Err":     "Invalid Phone",
				"Name":    User1.Name,
				"Email":   models.User_Email,
				"Country": User1.Country,
				"Phn":     User1.Phone[3:],
			})
			return
		}
		_, err = models.Db.Exec(`update information set name=$1, country=$2, phone=$3 where email=$4`, user.Name, user.Country, user.Phone, models.User_Email)
		if err != nil {
			http.Error(w, "Update fail Bad Req", 400)
			return
		}
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
}

// http.ServeFile(w, r, "templates/edit_profile.html")
