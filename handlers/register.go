package handlers

import (
	"fmt"
	"mini_project/models"
	"net/http"
	"strings"
)

func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Name1 := r.FormValue("name")
		Email := r.FormValue("email")
		Phone := r.FormValue("mobile")
		Country := r.FormValue("country")
		Pass := r.FormValue("pass")
		ConfPass := r.FormValue("confpass")
		if Pass != ConfPass {
			http.Error(w, "Password doesn't match", 400)
			return
		}
		isValidemail := strings.Contains(Email, "@")
		isValidphn := strings.Contains(Phone, "01") && len(Phone) == 11
		if isValidemail == false {
			http.Error(w, "Invalid Email", 400)
			return
		}
		if isValidphn == false {
			http.Error(w, "Invalid Phone", 400)
			return
		}
		if Name1 == "" || Country == "" || Pass == "" {
			http.Error(w, "All Fields are required", 400)
			return
		}
		sqlst := `insert into information (name,email,phone,country,password) values($1,$2,$3,$4,$5)`
		_, err := models.Db.Exec(sqlst, Name1, Email, Phone, Country, Pass)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintln(w, "User Register succesfully")
		return
	}
	http.ServeFile(w, r, "templates/register.html")
}
