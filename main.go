package main

import (
	// "database/sql"
	"fmt"
	"mini_project/handlers"
	"mini_project/models"
	"net/http"

	_ "github.com/lib/pq"
)

//	func registerHandle(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "Hello world")
//	}
func main() {
	handlers.ConnectDB()

	http.HandleFunc("/register", handlers.RegisterHandle)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/edit_profile", handlers.EditProfileHandle)
	http.HandleFunc("/create-post", handlers.CreatePost)
	http.HandleFunc("/delete-post", handlers.DeletePost)
	http.HandleFunc("/user_post", handlers.UserPost)
	http.HandleFunc("/update-post", handlers.UpdatePost)
	http.HandleFunc("/upload-image", handlers.UploadImg)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println(`Server running at http://localhost:8080/login`)
	http.ListenAndServe(":8080", nil)
	defer models.Db.Close()
}
