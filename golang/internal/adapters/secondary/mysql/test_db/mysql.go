package test_db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func Connect() *sqlx.DB {

	// err := godotenv.Load("../../../../../../env/mysql.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	database := os.Getenv("MYSQL_TEST_DATABASE")

	connectionUrl := fmt.Sprintf("%v:%v@(localhost:3306)/%v", dbUser, dbPassword, database)

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

func TruncateTable(db *sqlx.DB, tableName string) error {
	query := fmt.Sprintf("TRUNCATE TABLE %v", tableName)

	_, err := db.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func Insert(db *sqlx.DB, tableName string) {
	//TODO: Add generic insert
}
