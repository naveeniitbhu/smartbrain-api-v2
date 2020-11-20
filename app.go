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
	R.POST("/signin", app.PostSignin)
	R.GET("/profile/:id", app.GetProfile)
	R.PUT("/image", app.ImageCount)
}

func (app *App) PostRegister(c *gin.Context) {
	api.PostRegister(c, app.Db)
}

// func (app *App) GetUsers(c *gin.Context) {
// 	api.GetUsers(c, app.Db)
// }

func (app *App) PostSignin(c *gin.Context) {
	api.PostSignin(c, app.Db)
}

func (app *App) GetProfile(c *gin.Context) {
	api.GetProfile(c, app.Db)
}

func (app *App) ImageCount(c *gin.Context) {
	api.ImageCount(c, app.Db)
}
