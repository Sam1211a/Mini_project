package handlers

import (
	"fmt"
	"mini_project/models"
	"net/http"
)

func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "post" {
		Name := r.FormValue("name")
		Email := r.FormValue("email")
		Phone := r.FormValue("phone")
		Country := r.FormValue("country")
		sqlst := `insert into information(name,email,phone,country) values($1,$2,$3,$4)`
		_, models.Err = models.Db.Exec(sqlst, Name, Email, Phone, Country)
		if models.Err != nil {
			panic(models.Err)
		}
		fmt.Println(w, "User Register succesfully")
		return
	}
	http.ServeFile(w, r, "templates/register.html")
}
