package main

import (
	"fmt"
	"log"
	"net"
)

func SendMessage() {

	//Connect to Server
	conn, err1 := net.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer conn.Close()

	//Send Data to Server
	a, err2 := conn.Write([]byte("Hello there!"))
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(a) //prints out number of bytes sent to the server

	//Receive response from Server
	buffer := make([]byte, 1024)
	b, err3 := conn.Read(buffer)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(b)

	response := string(buffer[:b])
	fmt.Println("Server Response: ", response)

}

func main() {
	//Listen for incoming TCP connections on a scientific port
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	a, err := conn.Read(buffer)
	if err != nil {
		log.Println(err)
		return
	}

	data := string(buffer[:a])
	fmt.Println("Received data: ", data)

	response := "Hello there!!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Println(err)
		return
	}
}
