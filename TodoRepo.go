package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type TodoRepo struct {
}

var todos = []Todo{
	{
		Id:        "1",
		Completed: false,
		Name:      "Take out trash",
		Due:       time.Now()},
	{
		Id:        "2",
		Completed: false,
		Name:      "Clean windows",
		Due:       time.Now().Add(3 * 24 * 8 * 60 * 60 * 1000)}}

func Write(todo Todo) {
	f, err := os.Create("todo-db.csv")
	if err != nil {
		panic(err)
	}

	str := fmt.Sprint(todo.Id, ";", todo.Name, ";", todo.Due, ";", todo.Completed, "\n")
	f.WriteString(str)
}

func (r TodoRepo) add(t Todo) Todo {
	t.Id = strconv.Itoa(len(todos))

	todos = append(todos, t)

	return t
}

func (r TodoRepo) delete(id string) error {
	idx := indexOf(id)
	if idx < 0 {
		return errors.New("not found")
	}
	todos = append(todos[:idx], todos[idx+1:]...)

	return nil
}

func (r TodoRepo) getAll() []Todo {
	return todos
}

func indexOf(id string) int {
	for idx, t := range todos {
		if t.Id == id {
			return idx
		}
	}

	return -1
}

func (r TodoRepo) get(id string) (Todo, error) {
	idx := indexOf(id)

	return todos[idx], nil
}
