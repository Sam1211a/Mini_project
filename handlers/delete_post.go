package handlers

import (
	"mini_project/models"
	"net/http"
)

// var tmpl= template.Must(template.ParseFiles(""))
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "POST" {
	id := r.URL.Query().Get("id")
	_, err := models.Db.Exec("delete from user_post where id=$1", id)
	if err != nil {
		http.Error(w, "Delete error", 400)
		return
	}
	http.Redirect(w, r, "/user_post", http.StatusSeeOther)
}
