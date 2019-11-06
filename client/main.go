package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func main() {
	fmt.Println("Hello world")

	opts := &socketio_client.Options{}
	uri := "http://localhost:4001/"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On("newMessage", func(msg string) {
		log.Printf("on message:%v\n", msg)
	})
	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)
		client.Emit("newMessage", command)
		log.Printf("send message:%v\n", command)
		if len(command) > 0 {
			fmt.Println()
		}
	}
}
