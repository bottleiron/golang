package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Admin_id string
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")
	r.Static("assets/", "./assets")

	r.GET("/", func(c *gin.Context) {

		db, err := sql.Open("mysql", "nvp_planty:planty@tcp(1.255.153.228:3306)/nvp_svc")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		rows, err := db.Query("select admin_id from tbl_admin")
		if err != nil {
			fmt.Println("Error running query")
			fmt.Println(err)
			return
		}
		defer rows.Close()

		var users []User
		var userOnce User

		for rows.Next() {
			err := rows.Scan(&userOnce.Admin_id)
			if err != nil {
				fmt.Println(err)
			}
			users = append(users, userOnce)
		}

		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "여기는 index 입니다.",
			"users": users,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}
