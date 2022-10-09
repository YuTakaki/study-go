package src

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
			// ... handle error
			panic(err)
	}

	for _, v := range TodoLists {
		if i == v.ID {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, map[string]string{
		"error": "id not found",
	})
}
func GetAllTodos(c *gin.Context) {
	isComplete := strings.ToLower(c.Query("isComplete"))

	fmt.Println(isComplete)

	if isComplete == "" {
		c.IndentedJSON(http.StatusOK, TodoLists)
	} else {
		todos := []Todo{}

		validValue := []string{"true", "false"}

		if !slices.Contains(validValue, isComplete) {
			c.IndentedJSON(http.StatusBadRequest, map[string]string{
				"error" : "isComplete is not a valid type bool",
			})
			return
		}

		for _, v := range TodoLists {
			if strconv.FormatBool(v.IsComplete) == isComplete {
				todos = append(todos, v)
			}
		}

		c.IndentedJSON(http.StatusOK, todos)
	}
}

func AddTodo(c *gin.Context) {
	var todo Todo

	if err :=c.BindJSON(&todo); err != nil {
		fmt.Println(err.Error())
		return
	}
	if todo.Todo == "" {
		c.IndentedJSON(http.StatusBadRequest, map[string]string{
			"error" : "todo has no value",
		})
		return
	}

	todo.ID = TodoLists[len(TodoLists) - 1].ID + 1


	TodoLists = append(TodoLists, todo)

	c.IndentedJSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
			// ... handle error
			panic(err)
	}
	var todo Todo
	
	if err :=c.BindJSON(&todo); err != nil {
		fmt.Println(err.Error())
		return
	}
	
	for i, v := range TodoLists {
		if v.ID == id_int {
			todos := &TodoLists[i]
			todos.Todo = todo.Todo
			todos.IsComplete = todo.IsComplete
			c.IndentedJSON(http.StatusAccepted, todos)
			return
		}
	}
}

func DeleteTodo(c *gin.Context) {
	todos := []Todo{}
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
			// ... handle error
			panic(err)
	}
	var todo Todo

	for _, v := range TodoLists {
		if v.ID != i {
			todos = append(todos, v)
		} else {
			todo = v
		}
	}

	TodoLists = todos
	c.IndentedJSON(http.StatusAccepted, todo)
}