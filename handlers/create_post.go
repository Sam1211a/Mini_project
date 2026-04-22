package handlers

import (
	"mini_project/models"
	"net/http"
)

// var tmp= template.Must(template.ParseFiles("templates/dashboard.html"))
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// tmp.Execute(w,map[string]string{
	// })
	cookie, err := r.Cookie("user_email")
	if err != nil {
		http.Error(w, "Cookie not found", 400)
		return
	}
	models.User_Email = cookie.Value
	if r.Method == "POST" {
		user_post := r.FormValue("content")
		var user_image string
		models.Db.QueryRow(`select image from information where email=$1`, models.User_Email).Scan(&user_image)
		sqlstatement = "insert into user_post (email,contant,image) values($1,$2,$3)"
		_, err := models.Db.Exec(sqlstatement, models.User_Email, user_post, user_image)
		if err != nil {
			http.Error(w, "Post Not Executed", 400)
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	// http.ServeFile(w, r, "templates/dashboard.html")
}
