package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func InitToDB() *sql.DB {
	var connectionString = os.Getenv("DATABASE")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error making connection", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connecting to database", errPing.Error())
	} else {
		fmt.Println("koneksi berhasil")
	}
	return db
}
