package handlers

import (
	"html/template"
	"mini_project/models"
	"net/http"
)

var Tmpl = template.Must(template.ParseFiles("templates/home.html"))
var User1 models.Users

type Post struct {
	ID           int
	UserEmail    string
	Content      string
	CreatedAt    string
	ProfileImage string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		cookie, err := r.Cookie("user_email")
		if err != nil {
			http.Error(w, "Cookies not Found", 400)
			return
		}
		Email := cookie.Value
		err = models.Db.QueryRow("select name,phone,country,image from information where email=$1", Email).Scan(&User1.Name, &User1.Phone, &User1.Country, &User1.ProfileImage)
		if err != nil {
			http.Error(w, "Db img error", 400)
			return
		}
		var P Post
		row, Err := models.Db.Query("select id, email, contant, create_time, image from user_post order by create_time desc")

		if Err != nil {
			http.Error(w, "Post not exist", 400)
			return
		}
		defer row.Close()

		var posts []Post
		for row.Next() {
			var t string
			row.Scan(&P.ID, &P.UserEmail, &P.Content, &t, &P.ProfileImage)
			P.CreatedAt = t
			posts = append(posts, P)
		}
		Tmpl.Execute(w, map[string]interface{}{
			"ProfileImage": User1.ProfileImage,
			"Name":         User1.Name,
			"Email":        Email,
			"Country":      User1.Country,
			"Phn":          User1.Phone,
			"Posts":        posts,
		})
		// http.ServeFile(w, r, "templates/profile.html")
	}

}
