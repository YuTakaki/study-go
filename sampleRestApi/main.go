package main

import (
	"example/restapi/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", src.GetAllTodos)
	router.GET("/:id", src.GetTodoById)
	router.PUT("/:id", src.UpdateTodo)
	router.POST("/", src.AddTodo)
	router.DELETE("/:id", src.DeleteTodo)

	router.Run("localhost:4000")

}