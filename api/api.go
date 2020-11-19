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

func PostRegister(c *gin.Context, db *sqlx.DB) {

	var (
		json Login
		id   int64
	)

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("INFO: json binding is done")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"json binding error": err.Error()})
	}

	err := db.QueryRow(`INSERT INTO login(hash, email) VALUES($1, $2) RETURNING id`, json.Password, json.Email).Scan(&id)

	if err == nil {
		fmt.Printf("INFO: Login Details Inserted with id:%d & email:%s & name:%s\n", id, json.Email, json.Name)
	} else {
		log.Fatalln(err)
	}

	joined := time.Now()
	err = db.QueryRow(`INSERT INTO users(email, name, joined) VALUES($1, $2, $3) RETURNING id`, json.Email, json.Name, joined).Scan(&id)

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
