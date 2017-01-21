package main

import (
	"github.com/ridwanmsharif/mqueue/client"
	"log"
)

func main() {
	ch, err := client.Subscribe("topic_of_your_choice")
	if err != nil {
		log.Println(err)
		log.Println("wtf is happening")
		return
	}

	for e := range ch {
		log.Println(string(e))
	}

	log.Println("Channel closed")
}
