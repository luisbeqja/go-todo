package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luisbeqja/go-todo/config"
	"github.com/luisbeqja/go-todo/models"
)

func TodoCreate(c *gin.Context) {
	/* 	var title string = c.PostForm("title")
	 */ /* 	var description string = c.PostForm("description")
	 */ /* 	var status string = c.PostForm("checked") */
	todo := models.TodoItem{Title: "title2", Description: "description", Status: false}
	config.DB.AutoMigrate(&models.TodoItem{})
	config.DB.Create(&todo)
	fmt.Println(config.DB)
	c.JSON(200, gin.H{
		"message": todo,
	})
}
func TodoDelete(c *gin.Context) {
	config.DB.Where("title = ?", "title").Delete(&models.TodoItem{})
	config.DB.AutoMigrate(&models.TodoItem{})
}
func TodoUpdate(c *gin.Context) {
	config.DB.Model(&models.TodoItem{}).Where("title = ?", "title4").Update("title", "title2")
	config.DB.AutoMigrate(&models.TodoItem{})
}
func TodoGet(c *gin.Context) {
	var todos []models.TodoItem
	config.DB.Find(&todos)
	c.JSON(200, gin.H{
		"todos": todos,
	})
}
