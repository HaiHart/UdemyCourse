package database

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var connectionString=("root:donQuiote2@tcp(127.0.0.1:3306)/test_api")

var db, err = sql.Open("mysql", connectionString)

