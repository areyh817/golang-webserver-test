package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	QueryRegister = `
	insert into users(id, password)
	values($1, $2)`
)

var DB *sql.DB

func main() {

	r := gin.Default()
	r.GET("/ping", pong)
	r.POST("/orangina", orangina)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func orangina(ctx *gin.Context) {
	characterName := ctx.PostForm("character")
	age := ctx.PostForm("age")

	DB.ExecContext()
	rows, err := DB.QueryContext(ctx, QueryRegister, id, password)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error()
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
