package main

import (
	"fmt"
	"net"
	"os"

	com "github.com/betty2310/redigo/pkg/command"
)

const (
	BACKSPACE_KEY = 8
	CARET_KEY     = 13
	NEW_LINE_KEY  = 10
	CTRL_C_KEY    = 3
)

var _ = net.Listen
var _ = os.Exit

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Got connection from ", conn.LocalAddr().String())
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	command := make([]byte, 0, 128)
	kv := make(map[string]string, 0)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from client", err)
			return
		}

		for i := range n {
			b := buffer[i]
			switch b {
			case BACKSPACE_KEY:
				if len(command) > 0 {
					command = command[0 : len(command)-1]
				}
			case CARET_KEY:
			case NEW_LINE_KEY:
				com.ProcessCommand(conn, string(command), kv)
				command = command[:0]
			case CTRL_C_KEY:
				conn.Write([]byte("+Ok Goodbye\r\n"))
				conn.Close()
				return
			default:
				command = append(command, b)
			}
		}
	}
}
