package models

import "database/sql"

var Err error
var Db *sql.DB

type Users struct {
	Name           string
	Phone string
	Country string
}
