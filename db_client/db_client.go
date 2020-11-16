package db_client

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitialiseDBConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "user=postgres password=pspasswd dbname=smart-brain sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to database")

	return db, err
}

// const (
// 	port   = 5432
// 	DbName = "smart-brain-go"
// 	pgUser = "postgres"
// 	pgPass = "pspasswd"
// 	DbType = "postgres"
// )

// var DBClient *sqlx.DB
