package main

import (
	"smartgo/db_client"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db, _ := db_client.InitialiseDBConnection()

	app := App{
		R:  gin.Default(),
		Db: db,
	}

	app.R.Use(cors.Default())

	app.initialize()

	app.R.Run(":3000")
}
