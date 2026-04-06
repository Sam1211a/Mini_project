package handlers

import (
	"database/sql"
	"fmt"
	"mini_project/models"

	_ "github.com/lib/pq"
)

var sqlstatement string
var name, email, mob, country string

func ConnectDB() {
	connecStr := "user=postgres password=260897 dbname=postgres sslmode=disable"
	models.Db, models.Err = sql.Open("postgres", connecStr)
	if models.Err != nil {
		panic(models.Err)
	}
	// defer db.Close()
	fmt.Println("DB Connected")
}

func AddUsers() {
	fmt.Println("Enter Name :")
	fmt.Scanln(&name)
	fmt.Println("Enter Email :")
	fmt.Scanln(&email)
	fmt.Println("Enter Country :")
	fmt.Scanln(&country)
	fmt.Print("Enter Phone No :\n+88")
	fmt.Scanln(&mob)

	sqlstatement = `insert into information(name,email,phone,country) values($1,$2,$3,$4)`
	_, models.Err = models.Db.Exec(sqlstatement, name, email, "+88"+mob, country)
	if models.Err != nil {
		panic(models.Err)
	}
	fmt.Println("User Added")

}
func UserDetails() {
	sqlstatement = `select *from information where phone like $1`
	fmt.Print("Enter Phone Number to Show User Details\n+88")
	fmt.Scanln(&mob)
	models.Err = models.Db.QueryRow(sqlstatement, "+88"+mob).Scan(&name, &email, &mob, &country)
	if models.Err != nil {
		if models.Err == sql.ErrNoRows {
			fmt.Println("No User found !")
			return
		}
		panic(models.Err)
	}
	fmt.Println("Name: ", name, "Email: ", email, "Country: ", country, "Phone No- ", mob)
}
