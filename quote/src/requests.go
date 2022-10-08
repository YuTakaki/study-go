package src

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetRandomQuote() {
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
      log.Fatal(err)
  }

  fmt.Println(string(body))

}