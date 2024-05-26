package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect(user string, password string, database string) *sqlx.DB {

	connectionUrl := user + ":" + password + "@(localhost:3306)/" + database

	db, err := sqlx.Connect("mysql", connectionUrl)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		panic("Failed to pint db. error: " + err.Error())
	}

	return db

}
