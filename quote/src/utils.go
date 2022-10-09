package src

import (
	"fmt"
)

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

type QuoteList struct {
	Results []Quote `json:"results"`
}

func (quote Quote) TextOutput(){
	fmt.Printf("Content: %v\n", quote.Content)
	fmt.Printf("Author: %v\n", quote.Author)
}

func (quotes QuoteList) ListOutput(){
	for _, v:= range quotes.Results {
		fmt.Printf("%v\n", v.Content)
	}
}