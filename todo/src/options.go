package src

import (
	"fmt"
	"strconv"
	types "todo/types"
	"todo/utils"
)

func AddTodo() {
	var todo string
	fmt.Print("add todo: ")
	fmt.Scan(&todo)

	newTodo := types.Todo{
		Todo:       todo,
		IsComplete: false,
	}

	newTodoList := append(Todos, newTodo)
	Todos = newTodoList
	fmt.Printf("Added todo\n\n\n")
}

func RemoveTodo() {
	var indexChoice string
	ListTodo()
	fmt.Print("what to remove: ")
	fmt.Scan(&indexChoice)

	intVar, err := strconv.Atoi(indexChoice)
	if err != nil || intVar >= len(Todos) {
		fmt.Printf("incorrect index. try again\n\n\n")
		return 
	}

	updatedTodoList := utils.Filter(Todos, func(index int) bool {
		fmt.Println(indexChoice == strconv.Itoa(index))
		return indexChoice == strconv.Itoa(index)
	})
	Todos = updatedTodoList
	fmt.Printf("remove todo\n\n\n")
}


func ToggleTodo(){
	var indexChoice string
	ListTodo()
	fmt.Print("what to finish: ")
	fmt.Scan(&indexChoice)

	intVar, err := strconv.Atoi(indexChoice)
	if err != nil || intVar >= len(Todos) {
		fmt.Printf("incorrect index. try again\n\n\n")
		return 
	}

	updatedTodoList := utils.Map(Todos, func(x types.Todo, index int) types.Todo {
		if indexChoice == strconv.Itoa(index) {
			x.IsComplete = !x.IsComplete
		}
		return x
	})
	Todos = updatedTodoList
	fmt.Printf("toggle complete todo\n\n\n")
}

func ListTodo() {
	fmt.Println("list of todo: ")
	for i, val := range Todos {
		var v string
		if val.IsComplete {
			v = "complete"
		} else {
			v = "incomplete"
		}
		fmt.Printf("[%v] %v : %v\n", i, val.Todo, v)
	}
	fmt.Printf("\n\n\n\n")
}