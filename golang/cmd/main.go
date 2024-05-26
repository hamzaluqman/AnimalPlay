package main

import (
	"animal-play/internal/adapters/primary/web"
	"animal-play/internal/adapters/secondary/mysql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Application is running!")

	err := godotenv.Load("../../env/mysql.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// load secondary db adapter
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	db := mysql.Connect(dbUser, dbPassword, database)
	defer db.Close()

	// load primary web adapter
	server := &http.Server{
		Handler:      web.NewWebHandler(db).MountRoutes(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         ":8080",
	}

	fmt.Println("Starting server on port 8080")

	err = server.ListenAndServe()
	if err != nil {
		// add logging
		fmt.Println("Problem Running Server.... ", err)
	}

}
