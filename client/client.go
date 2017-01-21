package client

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	mqueueServer = "127.0.0.1:8081"
)

func Publish(topic string, payload []byte) error {
	rsp, err := http.Post(fmt.Sprintf("http://%s/pub?topic=%s",
		mqueueServer, topic), "application/json",
		bytes.NewBuffer(payload))

	if err != nil {
		return err
	}
	rsp.Body.Close()
	if rsp.StatusCode != 200 {
		return fmt.Errorf("Non 200 response %d", rsp.StatusCode)
	}
	return nil
}

func Subscribe(topic string) (<-chan []byte, error) {
	conn, _, err := websocket.DefaultDialer.Dial(
		fmt.Sprintf("ws://%s/sub?topic=%s", mqueueServer, topic), make(http.Header))

	if err != nil {
		return nil, err
	}

	ch := make(chan []byte, 100)

	go func() {
		for {
			t, p, err := conn.ReadMessage()
			if err != nil {
				log.Println("Could not read message, closing channel", err)
				conn.Close()
				return
			}
			switch t {
			case websocket.CloseMessage:
				log.Println("Close message, clossing channel")
				conn.Close()
				close(ch)
				return
			default:
				ch <- p
			}
		}
	}()
	return ch, nil
}
