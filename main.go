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

func main() {
	ip := "127.0.0.1"
	port := "503"

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

	for {
		conn, err := net.Dial("tcp", ip+":"+port)
		if err != nil {
			fmt.Println(err)
			time.Sleep(2)
			continue
		}

		fmt.Printf("Connected %s:%s\n", ip, port)

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
}
