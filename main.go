package main

import (
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type mqueue struct {
	topics map[string][]chan []byte
	mtx    sync.RWMutex
}

var (
	defaultMqueue = &mqueue{
		topics: make(map[string][]chan []byte),
	}
)

// Method to add subscriber to specefied topic in mqueue
func (m *mqueue) sub(topic string) (<-chan []byte, error) {
	channel := make(chan []byte, 100)
	m.mtx.Lock()
	m.topics[topic] = append(m.topics[topic], channel)
	m.mtx.Unlock()
	return channel, nil
}

// Method to remove subscriber from specefied topic in mqueue
func (m *mqueue) unsub(topic string, sub <-chan []byte) error {
	m.mtx.RLock()
	subscribers, ok := m.topics[topic]
	m.mtx.RUnlock()

	if !ok {
		return nil
	}
	var subs []chan []byte
	for _, subscriber := range subscribers {
		if subscriber != sub {
			subs = append(subs, subscriber)
		}
		continue
	}

	m.mtx.Lock()
	m.topics[topic] = subs
	m.mtx.Unlock()

	return nil
}

// Method to publish/push payload to every subscriber
func (m *mqueue) pub(topic string, payload []byte) error {
	m.mtx.RLock()
	subscribers, ok := m.topics[topic]
	m.mtx.RUnlock()

	if !ok {
		return nil
	}

	go func() {
		for _, subscriber := range subscribers {
			select {
			case subscriber <- payload:
			default:
			}
		}
	}()
	return nil
}

// Subscribe to a specefic topic in mqueue
func sub(w http.ResponseWriter, r *http.Request) {
	connection, err := websocket.Upgrade(w, r, w.Header(),
		1024, 1024)
	if err != nil {
		log.Println("Websocket connection failed:", err)
		http.Error(w, "Could not open websocket	connection",
			http.StatusBadRequest)
		return
	}

	topic := r.URL.Query().Get("topic")
	channel, err := defaultMqueue.sub(topic)
	if err != nil {
		log.Println("Could not retrieve %s.", topic)
		http.Error(w, "Could not retrieve events",
			http.StatusInternalServerError)
		return
	}
	defer defaultMqueue.unsub(topic, channel)

	for {
		select {
		case e := <-channel:
			err = connection.WriteMessage(websocket.BinaryMessage, e)
			if err != nil {
				log.Printf("Error sending event: %v",
					err.Error())
				return
			}
		}
	}
}

// Publishes and prints to console of every subscriber
func pub(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Pub Error", http.StatusInternalServerError)
		return
	}
	r.Body.Close()

	err = defaultMqueue.pub(topic, b)
	if err != nil {
		http.Error(w, "Pub Error", http.StatusInternalServerError)
		return
	}
}

// Entry Point
func main() {
	http.HandleFunc("/pub", pub)
	http.HandleFunc("/sub", sub)
	log.Println("Mqueue listening on :8081")
	http.ListenAndServe(":8081", nil)
}
