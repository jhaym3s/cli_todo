package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct{
	Add string
	Edit string 
	Delete int
	Toggle int 
	List bool
}


func NewCmdFlags() *CmdFlags{
	cf := CmdFlags{}
	
	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit existing todo with ID ")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a todo using ID")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle completed and uncompleteed todo")
	flag.BoolVar(&cf.List, "list", false, "List all todo")

	flag.Parse()

	return &cf

}  

func (cf *CmdFlags) ExecuteFlags(todo *TodoList){
	switch  {
	case cf.List:
		todo.PrintAll()
	case cf.Add != "":
		todo.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, to edit use ':' to seperate id from new description")
			os.Exit(1)
			
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error, invalid id for edit:", parts[0])
			os.Exit(1)
		}
		if valErr := todo.validateID(id); valErr != nil{
			fmt.Println("this id does not exist") 
			os.Exit(1)

		}
		todo.MakeEdit(id, parts[1])

	case cf.Delete != -1:
		todo.Delete(cf.Delete)
	
	case cf.Toggle != -1:
		todo.MarkCompleted(cf.Toggle)

	default :
		fmt.Println("invalid command")
	}
}


