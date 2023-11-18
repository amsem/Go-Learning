package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	var msg string
	//for msg == "" {
		//establish connection
		connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
		defer connection.Close()
        
        if err != nil {
			panic(err)
		}
		// GET USER MESSAGE
		fmt.Print("Message : ")
		fmt.Scanf(`%v`, &msg)
		///send some data

		_, err = connection.Write([]byte(msg))
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("Received: ", string(buffer[:mLen]))
	

}
