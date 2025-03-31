package command

import (
	"fmt"
	"net"
)

func ProcessCommand(conn net.Conn, command string, kv map[string]string) {
	fmt.Printf("Processing command: %v\n", command)
	args := SplitArguments(command)
	switch args[0] {
	case "ping":
		conn.Write([]byte("+PONG\r\n"))
	case "quit":
		conn.Write([]byte("+OK\r\n"))
		conn.Close()
	case "echo":
		conn.Write([]byte(args[1]))
		conn.Write([]byte("\r\n"))
	case "set":
		kv[args[1]] = args[2]
		conn.Write([]byte("+OK\r\n"))
	case "get":
		val, ok := kv[args[1]]
		if !ok {
			conn.Write([]byte("-1\r\n"))
		} else {
			conn.Write([]byte(val))
			conn.Write([]byte("\r\n"))
		}
	default:
		conn.Write([]byte("-ERR unknown command '" + command + "'\r\n"))
	}
}

func SplitArguments(command string) []string {
	args := make([]string, 0)
	openQuotation := -1
	token := make([]byte, 0)
	for i := range command {
		switch command[i] {
		case ' ':
			if openQuotation == -1 {
				args = append(args, string(token))
				token = token[:0]
			}
		case '"':
			if openQuotation == -1 {
				openQuotation = i
			} else {
				token = []byte(command[openQuotation+1 : i])
				args = append(args, string(token))
				openQuotation = -1
				token = token[:0]
			}
		default:
			token = append(token, command[i])
		}
	}
	args = append(args, string(token))
	return args
}
