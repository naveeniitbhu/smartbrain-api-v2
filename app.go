package main

import (
	"smartgo/api"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	R  *gin.Engine
	Db *sqlx.DB
}

func (app *App) initialize() {
	R := app.R
	R.POST("/register", app.PostRegister)
}

func (app *App) PostRegister(c *gin.Context) {
	api.PostRegister(c, app.Db)
}
