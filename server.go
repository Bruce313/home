package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/websocket"
	"github.com/Bruce313/home/agent"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	confPath := "./config.toml"
	conf, err := ParseConfig(confPath)
	if err != nil {
		log.Println(err)
		return
	}

	_ = agent.NewLightSensor(conf.LightSensors[0].HttpURL)

	http.Handle("/light", &LightHandler{})

	http.Handle("/websocket", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
