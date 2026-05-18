package main

import (
	"errors"
	"time"
)

type Todo struct {
	ID          int
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type TodoList []Todo

func (t *TodoList) Add(description string) {
	newTodo := Todo{
		ID:          len(*t) + 1,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	*t = append(*t, newTodo)
}

func (t TodoList) validateID(id int) error {
	
	for _, todo := range t {
		if todo.ID == id {
			return nil
		}
	}
	return errors.New("ID DOES NOT EXIST")
}

func (t *TodoList) Delete(id int) error {
	todos := *t
	if err := t.validateID(id); err != nil {
		return err
	}
	 newList := append(todos[:id], todos[id+1:]...)
	 *t = newList
	 return nil
}

func (t *TodoList) MakeEdit(id int, description string) error {
	todos := *t
	if err := t.validateID(id); err != nil {
		return err
	}
	for _, todo := range todos {
		if todo.ID == id {
			todos[id].Description = description
			return nil
		}
	}
		return nil
	}

func (t *TodoList) MarkCompleted(id int) error {
	todos := *t
	if err := t.validateID(id); err != nil {
		return err
	}
	isCompleted := todos[id].Completed
	now := time.Now()
	if !isCompleted {
		todos[id].Completed = true
		todos[id].CompletedAt = &now
	}
	todos[id].Completed = !isCompleted
	return nil
}