package models
import "database/sql"
var Err error
var Db *sql.DB
type Student struct{
	Name string
	
	Class int
}