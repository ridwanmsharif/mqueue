package main

import (
	"github.com/ridwanmsharif/mqueue/client"
	"log"
)

func main() {
	err := client.Publish("topic_of_your_choice", []byte("ridwan"))
	if err != nil {
		log.Println(err)
	}
}
