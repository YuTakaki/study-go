package src

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func GetRandomQuote() {
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

  var quote Quote

	if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
		panic(err.Error())
	}

	quote.TextOutput()
}


func SearchQuote() {
	var query string
	fmt.Printf("enter query: ")
	fmt.Scan(&query)
	
	resp, err := http.Get("https://api.quotable.io/search/quotes?query=" + query)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

  var quotes QuoteList

	if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		panic(err.Error())
	}

	quotes.ListOutput()

}