package handlers

import (
	"html/template"
	"mini_project/models"
	"net/http"
	"strings"
)

// type PageData struct {
// 	Error    string

// 	PhnError string

// }

var temp = template.Must(template.ParseFiles("templates/register.html"))

func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp.Execute(w, map[string]string{
			"Error": "",
		})
		return
	}
	if r.Method == "POST" {
		Name1 := r.FormValue("name")
		Email := r.FormValue("email")
		Phone := "+88" + r.FormValue("mobile")
		Country := r.FormValue("country")
		Pass := r.FormValue("pass")
		ConfPass := r.FormValue("confpass")
		isValidemail := strings.Contains(Email, "@")
		isValidphn := strings.Contains(Phone, "01") && len(Phone) == 14
		if Name1 == "" {
			temp.Execute(w, map[string]string{
				"ErrName": "All Fields are required !",
				"Name":    Name1,
				"Email":   Email,
				"Country": Country,
				"Phone":   Phone,
			})
			return
		}
		if isValidemail == false {
			temp.Execute(w, map[string]string{
				"Error":   "invalid email",
				"Name":    Name1,
				"Email":   Email,
				"Country": Country,
				"Phone":   Phone,
			})
			return
		}
		if Country == "" {
			temp.Execute(w, map[string]string{
				"ErrCountry": "All Fields are required !",
				"Name":       Name1,
				"Email":      Email,
				"Country":    Country,
				"Phone":      Phone,
			})
			// temp.Execute(w, Error)
			return
		}

		// temp.Execute(w, map[string]string{})
		if isValidphn == false {
			temp.Execute(w, map[string]string{
				"PhnError": "Invalid phone",
				"Name":     Name1,
				"Email":    Email,
				"Country":  Country,
				"Phone":    Phone,
			})
			return
		}

		if Pass != ConfPass || Pass == "" {
			temp.Execute(w, map[string]string{
				"PassError": "Password doesn't match",
				"Name":      Name1,
				"Email":     Email,
				"Country":   Country,
				"Phone":     Phone,
			})
			return
		}
		sqlst := `insert into information (name,email,phone,country,password) values($1,$2,$3,$4,$5)`
		_, err := models.Db.Exec(sqlst, Name1, Email, Phone, Country, Pass)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.ServeFile(w, r, "templates/register.html")
}
