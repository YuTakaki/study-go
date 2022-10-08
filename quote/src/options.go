package src

import "fmt"
func Menu() {
	for {
		fmt.Println("Hello welcome to quote")
		var choice string
	
		fmt.Println("[1] quote")
		fmt.Println("[2] random quote")
		fmt.Println("[3] search quote")
		fmt.Println("[4] authors popular quote")
		fmt.Println("[5] close")
	
		fmt.Print("what is your choice: ")
		fmt.Scan(&choice)
		MenuLists(choice)
	}
}

func MenuLists(choice string) {
	switch choice {
	case "1":
		fmt.Println("hello")
	case "2":
		GetRandomQuote()
	case "3":
		fmt.Println("hello")
	case "4":
		fmt.Println("hello")
	case "5":
		fmt.Println("hello")
		break
	default:
		fmt.Printf("wring choices. try again \n\n\n")
		Menu()
	}
}