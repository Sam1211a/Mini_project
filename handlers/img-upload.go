package handlers

import (
	"io"
	"mini_project/models"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadImg(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_email")
	if err != nil {
		http.Error(w, "Cookie not exits", 400)
		return
	}
	Email := cookie.Value
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "File Error", 400)
		return
	}
	file.Close()

	filename := time.Now().Format("20060102150405") + filepath.Ext(handler.Filename)
	filepath := "static/upload/" + filename
	dst, err := os.Create(filepath)
	if err != nil {
		http.Error(w, "create path error", 400)
		return
	}
	defer dst.Close()
	io.Copy(dst, file)
	_, err = models.Db.Exec("update information set image=$1 where email =$2 ", filepath, Email)
	if err != nil {
		http.Error(w, "DB Error img", 400)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
