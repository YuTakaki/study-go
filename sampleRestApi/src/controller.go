package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAllTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, TodoLists)
}

