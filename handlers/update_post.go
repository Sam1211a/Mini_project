package handlers

import (
	"mini_project/models"
	"net/http"
)

// var Tmpl = template.Must(template.ParseFiles("templates/user_post.html"))

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		id := r.FormValue("id")
		content := r.FormValue("content")
		_, err := models.Db.Exec("update user_post set contant=$1 where id=$2", content, id)
		if err != nil {
			http.Error(w, "Update Error", 400)
			return
		}
		w.Write([]byte("ok"))

	}
}
