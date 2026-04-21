package models

import "database/sql"

var Err error
var Db *sql.DB
var User_Email string

type Users struct {
	Name      string
	Phone     string
	Country   string
	Content   string
	CreatedAt string
}
