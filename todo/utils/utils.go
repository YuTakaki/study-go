package utils

import (
	"todo/types"
)

func Filter(data []types.Todo, callback types.SliceFilterType) []types.Todo {
	newTodoList := []types.Todo{}
	for index, val := range data {
		if !callback(index) {
			newTodoList = append(newTodoList, val)
		}
	}

	return newTodoList
}

func Map(data []types.Todo, callback types.SliceMapType) []types.Todo {
	newTodoList := []types.Todo{}
	for index, val := range data {
		newTodoList = append(newTodoList, callback(val, index))
	}
	return newTodoList
}