// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"smartgo/db_client"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jmoiron/sqlx"
// )

// type App struct {
// 	R  *gin.Engine
// 	Db *sqlx.DB
// }

// type Users struct {
// 	ID      int64        `json:"id,omitempty" db:"id"`
// 	Name    string       `json:"name,omitempty" db:"name"`
// 	Email   string       `json:"email,omitempty" db:"email"`
// 	Entries int64        `json:"entries,omitempty" db:"entries"`
// 	Joined  sql.NullTime `json:"joined,omitempty" db:"joined"`
// }

// type Login struct {
// 	ID    int64  `json:"id,omitempty" db:"id"`
// 	Hash  string `json:"hash,omitempty" db:"hash"`
// 	Email string `json:"email,omitempty" db:"email"`
// }

// func main() {
// 	db, _ := db_client.InitialiseDBConnection()

// 	app := App{
// 		R:  gin.Default(),
// 		Db: db,
// 	}

// 	app.R.POST("/register", app.PostRegister)

// 	app.R.Run(":3000")
// }

// func (app *App) PostRegister(c *gin.Context) {

// 	var (
// 		json Login
// 		id   int64
// 	)

// 	if err := c.ShouldBindJSON(&json); err == nil {
// 		fmt.Println("INFO: json binding is done")
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"json binding error": err.Error()})
// 	}

// 	db := app.Db
// 	err := db.QueryRow(`INSERT INTO login(hash, email) VALUES($1, $2) RETURNING id`, json.Hash, json.Email).Scan(&id)

// 	if err == nil {
// 		fmt.Printf("INFO: Login Details Inserted with id: %d\n", &id)
// 	} else {
// 		log.Fatalln(err)
// 	}
// }

// if err := c.ShouldBindJSON(&json); err == nil {
// 	if json.ID == 1 && json.Hash == "qwe" && json.Email == "new@gm.com" {
// 		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
// 	} else {
// 		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
// 	}
// } else {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// }

// errSlice := rows.StructScan(&usersSlice[i])
// i++
// if errSlice != nil {
// 	log.Fatalln(errSlice)
// }

// usersSlice := [4]Users{}
// i := 0

// users := Users{}
// rows, err := db.Queryx("SELECT * FROM users")

// if err != nil {
// 	log.Fatalln(err)
// }

// for rows.Next() {

// 	err := rows.StructScan(&users)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println("\nFifth", users)
// 	fmt.Println()
// }

// xy := []struct {
// 	id      int
// 	name    string
// 	email   string
// 	entries string
// 	joined  int
// }{
// 	{16, "jojo", "jo@gm.com", "0", 12},
// }

// fmt.Printf("%+v", xy)

// row := db.QueryRow(`SELECT name FROM users WHERE id=$1;`, 30)
// var name string
// err = row.Scan(&name)
// fmt.Println(err, name)

// if errr == nil {
// 	fmt.Printf("INFO: Users Details Inserted with id:%d\n", id)
// 	c.JSON(http.StatusOK, gin.H{
// 		"id":      24,
// 		"name":    "john",
// 		"email":   "john@gm.com",
// 		"entries": "0",
// 		"joined":  "2020-11-16T13:17:45.269Z",
// 	})
// } else {
// 	log.Fatalln(err)
// }
