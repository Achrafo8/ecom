package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Achrafo8/ecom/cmd/api"
	"github.com/Achrafo8/ecom/config"
	"github.com/Achrafo8/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{

		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	iniStorage(db)

	server := api.NEWAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func iniStorage(db *sql.DB) {

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: successfully connected!")
}

func hello() {

	fmt.Println("Hello there")
}
