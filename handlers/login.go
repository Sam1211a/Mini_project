package handlers

import (
	"html/template"
	"mini_project/models"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var temp = template.Must(template.ParseFiles("templates/login.html"))
	if r.Method == "GET" {
		temp.Execute(w, map[string]string{
			"EmailError": "",
		})
		return
	}
	if r.Method == "POST" {
		Emil := r.FormValue("emil")
		Pass := r.FormValue("pass")
		var Uname, checkPass string
		err := models.Db.QueryRow("select name,password from information where email = $1", Emil).Scan(&Uname, &checkPass)
		// fmt.Fprintln(w, Uname, checkPass, Emil)
		if err != nil {
			temp.Execute(w, map[string]string{
				"EmailError": "User Email is not valid !",
				"Email":      Emil,
			})
			return
		}

		if CheckPass(checkPass, Pass) == false {
			temp.Execute(w, map[string]string{
				"PassEr": "Incorrect Password !",
				"Email":  Emil,
			})
			return
		}
		cookie := &http.Cookie{
			Name:  "user_email",
			Value: Emil,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		// http.ServeFile(w, r, "templates/user_post.html")
	}
	// http.ServeFile(w, r, "templates/login.html")
}
