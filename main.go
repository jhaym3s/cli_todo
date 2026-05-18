package main

import "fmt"


func main(){
	todoList := TodoList{}
	todoList.Add("Buy groceries")
	todoList.Add("Walk the dog")
	todoList.Add("Read a book")
	fmt.Println("Todo List:")	
	editErr := todoList.MakeEdit(1, "Buy groceries and cook dinner")
	if editErr != nil {
		fmt.Println("Error editing todo:", editErr)
	}
	for _, todo := range todoList {
		fmt.Printf("ID: %d, Description: %s, Completed: %t, Created At: %v\n", todo.ID, todo.Description, todo.Completed, todo.CreatedAt)
	}
	todoList.PrintAll()
	err := todoList.Delete(9)
	if err != nil {
		fmt.Println("Error deleting todo:", err)
	} else {
		fmt.Println("Todo with ID 2 deleted successfully.")
	}
}	