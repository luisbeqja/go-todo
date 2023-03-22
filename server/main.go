package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisbeqja/go-todo/config"
	"github.com/luisbeqja/go-todo/controllers"
)

func init() {
	config.ConncectToDB()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../client/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/todos", func(c *gin.Context) {
		controllers.TodoGet(c)
	})
	router.POST("/create", func(c *gin.Context) {
		controllers.TodoCreate(c)
	})
	router.DELETE("/delete", func(c *gin.Context) {
		controllers.TodoDelete(c)
	})
	router.POST("/update", func(c *gin.Context) {
		controllers.TodoUpdate(c)
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
