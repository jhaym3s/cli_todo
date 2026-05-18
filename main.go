package main


func main(){
	todoList := TodoList{}
	storage := NewStorage[TodoList]("todos.json")
	storage.Load(&todoList)
	// todoList.Add("Buy groceries")
	// todoList.Add("Walk the dog")
	// todoList.Add("Read a book")
	storage.Save(todoList)
	todoList.PrintAll()
	
}	