package handlers

import (
	"html/template"
	"mini_project/models"
	"net/http"
)

var Tmp = template.Must(template.ParseFiles("templates/user_post.html"))

func UserPost(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_email")
	if err != nil {
		http.Error(w, "Cookie not found", 400)
		return
	}
	Email := cookie.Value
	row, Err := models.Db.Query("select id , email, contant,create_time from user_post where email=$1 order by id desc ", Email)
	if Err != nil {
		http.Error(w, "Db Error", 400)
		return
	}
	defer row.Close()
	var post []Post
	var p Post
	for row.Next() {
		var t string
		row.Scan(&p.ID, &p.UserEmail, &p.Content, &t)
		p.CreatedAt = t
		post = append(post, p)
		// if Email == p.UserEmail {

		// }
	}
	if r.Method == "GET" {
		Tmp.Execute(w, map[string]interface{}{
			"Name":    User1.Name,
			"Email":   Email,
			"Country": User1.Country,
			"Phn":     User1.Phone,
			"Posts":   post,
		})
	}
}
