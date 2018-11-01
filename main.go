package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func inputLoop(conn net.Conn) {
	for {
		messageReader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message: ")
		messageString, _ := messageReader.ReadString('\n')

		message := strings.TrimSpace(messageString)

		if message != "" {
			conn.Write([]byte("**" + message + "**"))
			conn.Write([]byte("\r\n"))

			fmt.Printf("Send: %s\n", "**"+message+"**")

			buff := make([]byte, 1024)
			n, _ := conn.Read(buff)
			log.Printf("Receive: %s", buff[:n])
		} else {
			fmt.Printf("Error: %s\n", "Can't send empty message")
		}
	}
}

func main() {
	ip := "127.0.0.1"
	port := "503"
	yn := ""
	conType := "client"

	fmt.Printf("KUPER, default: %s:%s\n", ip, port)

	ipReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ip: ")
	ipString, _ := ipReader.ReadString('\n')

	if strings.TrimSpace(ipString) != "" {
		ip = strings.TrimSpace(ipString)
	}

	portReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter port: ")
	portString, _ := portReader.ReadString('\n')

	if strings.TrimSpace(portString) != "" {
		port = strings.TrimSpace(portString)
	}

	fmt.Printf("Using %s:%s \n", ip, port)

	fmt.Printf("Use as server? (y/N): ")
	_, err := fmt.Scan(&yn)

	if err != nil {
		panic(err)
	}

	yn = strings.TrimSpace(yn)
	yn = strings.ToLower(yn)

	if yn == "y" || yn == "yes" {
		conType = "server"
	}

	if conType == "client" {
		for {
			conn, err := net.Dial("tcp", ip+":"+port)
			if err != nil {
				fmt.Println(err)
				time.Sleep(2)
				continue
			}

			fmt.Printf("Connected %s:%s\n", ip, port)

			go inputLoop(conn)
		}
	} else if conType == "server" {
		fmt.Printf("Server listening on %s:%s\n", ip, port)

		conn, err := net.Listen("tcp", ip+":"+port)
		if err != nil {
			panic(err)
		}

		for {
			conn, err := conn.Accept()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}

			log.Printf("Client connected")

			go inputLoop(conn)
		}
	}
}
