package main

import (
	"example/restapi/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", src.GetAllTodos)
	router.GET("/:id", src.GetTodoById)
	router.POST("/:id", src.UpdateTodo)
	router.POST("/", src.AddTodo)

	router.Run("localhost:4000")

}