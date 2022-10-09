package main

import (
	"example/restapi/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", src.GetAllTodos)

	router.Run("localhost:4000")

}