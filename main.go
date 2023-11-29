package main

import (
	"fmt"
	"net"
	"log"
)

func isServer() bool {
	fmt.Println("Choose program type: 1 - server | 2 - client (default = 2)")
	var program_type int
	fmt.Scanln(&program_type)
	if program_type == 1 {
		return true
	} else if program_type == 2 {
		return false
	}
	return false
}

func startServer(port string) {
	fmt.Println("Starting server on port: " + port)
	server, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal(err) // Server could not be created
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Print(err)
	}

	fmt.Print(string(buf))
}

func connectToServer(ip string, port string) {
	conn, err := net.Dial("tcp", ip + ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	var username string
	fmt.Println("Please enter username:")
	fmt.Scanln(&username)
	conn.Write([]byte(username + " joined the chat room!"))

	for {
		var message string
		fmt.Scanln(&message)
		conn.Write([]byte(message))
	}
}

func main() {
	if isServer() {
		fmt.Println("Please select server port (default = 8080):")
		var port string
		fmt.Scanln(&port)
		if port == "" {
			port = "8080"
		}
		startServer(port)
	} else {
		var ip string
		var port string
		fmt.Println("Please enter server IP address:")
		fmt.Scanln(&ip)
		fmt.Println("Please enter server port:")
		fmt.Scanln(&port)
		if ip == "" {
			ip = "localhost"
		}
		if port == "" {
			port = "8080"
		}
		connectToServer(ip, port)
	}
}