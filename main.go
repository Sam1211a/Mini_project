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
	fmt.Println(`Server running at http://localhost:8080/register`)
	http.ListenAndServe(":8080", nil)
	defer models.Db.Close()
}

// http.HandleFunc("/register",registerHandle)
// http.ListenAndServe(":8080",nil)
// fmt.Println("Press 1 to Add User")
// fmt.Println("Press 2 to Find User")
// var n int
// fmt.Scanln(&n)
// if n == int(1) {
// 	handlers.AddUsers()
// } else if n == 2 {
// 	handlers.UserDetails()
// }
// defer models.Db.Close()
// }

// connecstr:="user=postgres password=260897 dbname=postgres sslmode=disable"
// db,err:= sql.Open("postgres",connecstr)
// if err!=nil{
// 	panic(err)
// }
// defer db.Close()
// fmt.Println("Database connected Successfully")
// nam:="Efat"
// eml:="efathasan@gmail.com"
// phn:="+8801718355998"
// cntry:="Bangladesh"

// sqlStatement:= `insert into information(name, email,country,phone) values($1,$2,$3,$4)`
// _,err = db.Exec(sqlStatement,nam,eml,cntry,phn)
// if err!=nil{
// 	fmt.Println("Error:",err)
// 	return

// }
// db.Close()
// fmt.Println("Insert Data Succesfully")
// }

// type student struct {
// 	Name  string
// 	Class int
// }

// var st []student

//	func adduser(name string, class int) {
//		std := student{Name: name, Class: class}
//		st = append(st, std)
//		fmt.Println("Add users Succesfully")
//	}
//
//	func showuser(){
//		for i:=0;i<len(st);i++{
//			fmt.Println("Name - ",st[i].Name," Class - ",st[i].Class)
//		}
//	}
// func main() {
// 	fmt.Println("Welcome to my First project")
// 	handlers.Adduser("Sabbir", 10)
// 	handlers.Adduser("Efat", 9)
// 	handlers.Showuser()
// }
