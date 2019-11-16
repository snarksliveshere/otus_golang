package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	timeout int
	host    string
	port    string
)

const (
	timeoutDefault = 10
)

func init() {
	flag.IntVar(&timeout, "timeout", timeoutDefault, "timeout")
	flag.StringVar(&host, "host", "127.0.0.1", "host")
	flag.StringVar(&port, "port", "3302", "port")
}

func readRoutine(ctx context.Context, cancel context.CancelFunc, conn net.Conn, ch chan<- os.Signal) {
	scanner := bufio.NewScanner(conn)

LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN")
				ch <- os.Signal(syscall.SIGTERM)
				//break LOOP
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)
		}
	}
	log.Printf("Finished readRoutine")
	return
}

func writeRoutine(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
WRITER:
	for {
		select {
		case <-ctx.Done():
			break WRITER
		default:
			if !scanner.Scan() {
				break WRITER
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	log.Printf("Finished writeRoutine")
	return
}

func main() {
	t := time.Duration(timeout) * time.Second
	fmt.Println(t)

	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), t)

	conn, err := dialer.DialContext(ctx, "tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	stopch := make(chan os.Signal, 1)
	signal.Notify(stopch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		readRoutine(ctx, cancel, conn, stopch)
	}()

	go func() {
		writeRoutine(ctx, conn)
	}()

	<-stopch

	err = conn.Close()
	if err != nil {
		log.Fatal("I cannot close the connection")
	}
	fmt.Println("ola3")
}
