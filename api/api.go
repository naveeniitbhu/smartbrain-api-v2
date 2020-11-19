package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Users struct {
	ID      int64        `json:"id,omitempty" db:"id"`
	Name    string       `json:"name,omitempty" db:"name"`
	Email   string       `json:"email,omitempty" db:"email"`
	Entries int64        `json:"entries,omitempty" db:"entries"`
	Joined  sql.NullTime `json:"joined,omitempty" db:"joined"`
}

type Login struct {
	ID       int64  `json:"id,omitempty" db:"id"`
	Name     string `json:"name,omitempty" db:"name"`
	Email    string `json:"email,omitempty" db:"email"`
	Password string `json:"password,omitempty" db:"hash"`
}

type SignIn struct {
	Email    string `json:"email,omitempty" db:"email"`
	Password string `json:"password,omitempty" db:"hash"`
}

func PostRegister(c *gin.Context, db *sqlx.DB) {

	var (
		login Login
		id    int64
	)

	if err := c.ShouldBindJSON(&login); err == nil {
		fmt.Println("INFO: json binding is done")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"json binding error": err.Error()})
	}

	err := db.QueryRow(`INSERT INTO login(hash, email) VALUES($1, $2) RETURNING id`, login.Password, login.Email).Scan(&id)

	if err == nil {
		fmt.Printf("INFO: Login Details Inserted with id:%d & email:%s & name:%s\n", id, login.Email, login.Name)
	} else {
		log.Fatalln(err)
	}

	joined := time.Now()
	err = db.QueryRow(`INSERT INTO users(email, name, joined) VALUES($1, $2, $3) RETURNING id`, login.Email, login.Name, joined).Scan(&id)

	if err == nil {
		fmt.Printf("INFO: Users Details Inserted with id:%d\n", id)
	} else {
		log.Fatalln(err)
	}

	users := Users{}
	rows, err := db.Queryx("SELECT * FROM users WHERE id=$1;", id)

	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&users)
		if err != nil {
			log.Fatalln(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      users.ID,
		"name":    users.Name,
		"email":   users.Email,
		"entries": users.Entries,
		"joined":  users.Joined,
	})
}

func PostSignin(c *gin.Context, db *sqlx.DB) {

	var (
		signin   SignIn
		email    string
		password string
	)

	if err := c.ShouldBindJSON(&signin); err == nil {
		fmt.Println("INFO: json binding is done")
		email = signin.Email
		password = signin.Password
		if len(email) < 0 || len(password) < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"incorrect form submission": "email or password is missing"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"json binding error": err.Error()})
	}

	users := Users{}
	_, err := db.Queryx("SELECT * FROM login WHERE email=$1 AND hash=$2;", email, password)

	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Queryx("SELECT * FROM users WHERE email=$1;", email)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&users)
		if err != nil {
			log.Fatalln(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      users.ID,
		"name":    users.Name,
		"email":   users.Email,
		"entries": users.Entries,
		"joined":  users.Joined,
	})
}
