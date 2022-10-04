package src

import "fmt"

func Menu() {
	var choice string
	menuLoop: for {
		fmt.Println("What do you want to do")
		fmt.Println("[1] add")
		fmt.Println("[2] remove")
		fmt.Println("[3] toggle complete")
		fmt.Println("[4] list")
		fmt.Println("[5] exist")
		fmt.Print("answer: ")
		fmt.Scan(&choice)
		fmt.Printf("your choice is %v\n", choice)
		switch choice {
		case "1":
			AddTodo()
		case "2":
			RemoveTodo()
		case "3":
			ToggleTodo()
		case "4":
			ListTodo()
		case "5":
			fmt.Println("thank you. see you again")
			break menuLoop
		default:
			fmt.Printf("wrong choice try again\n\n\n")
			Menu()	
		}
	}
}