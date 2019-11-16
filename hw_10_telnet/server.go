package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const retry = 3

func handleConnection(conn net.Conn) {
	writeHandler(conn, "Welcome to %s, friend from %s\n", []interface{}{conn.LocalAddr(), conn.RemoteAddr()})

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("RECEIVED: %s", text)
		if text == "quit" || text == "exit" {
			break
		}

		writeHandler(conn, "I have received '%s'\n", []interface{}{text})
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error happend on connection with %s: %v", conn.RemoteAddr(), err)
	}

	log.Printf("Closing connection with %s", conn.RemoteAddr())

	func() {
		_ = conn.Close()
	}()
	return

}

func writeHandler(conn net.Conn, pattern string, text []interface{}) {
	var err error
	for i := 0; i < retry; i++ {
		_, err = conn.Write([]byte(fmt.Sprintf(pattern, text...)))
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Fatal("WRITE: Basic server functionality doesnt work, panic")
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:3302")
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Cannot accept: %v", err)
		}
		go handleConnection(conn)
		break
	}
	func() {
		_ = l.Close()
	}()
}
