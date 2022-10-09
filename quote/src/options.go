package src

import "fmt"
func Menu() {
	loop : for {
		fmt.Println("Hello welcome to quote")
		var choice string
	
		fmt.Println("[1] random quote")
		fmt.Println("[2] search quote")
		fmt.Println("[3] close")
	
		fmt.Print("what is your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			GetRandomQuote()
		case "2":
			SearchQuote()
		case "3":
			fmt.Println("hello")
			break loop
		default:
			fmt.Printf("wring choices. try again \n\n\n")
			Menu()
		}
	}
}
