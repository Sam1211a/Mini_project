package handlers

import (
	"html/template"
	"mini_project/models"
	"net/http"
)

var Tmpl = template.Must(template.ParseFiles("templates/dashboard.html"))
var User1 models.Users

type Post struct {
	ID        int
	UserEmail string
	Content   string
	CreatedAt string
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cookie, err := r.Cookie("user_email")
		if err != nil {
			http.Error(w, "Cookies not Found", 400)
			return
		}
		Email := cookie.Value
		err = models.Db.QueryRow("select name,phone,country from information where email=$1", Email).Scan(&User1.Name, &User1.Phone, &User1.Country)

		var P Post
		row, Err := models.Db.Query("select id, email, contant, create_time from user_post order by create_time desc")
		if err != nil {
			http.Error(w, "Bad Reqest row query", 400)
			return
		}

		if Err != nil {
			http.Error(w, "Post not exist", 400)
			return
		}
		defer row.Close()

		var posts []Post
		for row.Next() {
			var t string
			row.Scan(&P.ID, &P.UserEmail, &P.Content, &t)
			P.CreatedAt = t

			posts = append(posts, P)
		}
		Tmpl.Execute(w, map[string]interface{}{
			"Name":    User1.Name,
			"Email":   Email,
			"Country": User1.Country,
			"Phn":     User1.Phone,
			"Posts":   posts,
		})
		// http.ServeFile(w, r, "templates/profile.html")
	}

}
