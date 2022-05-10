package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: TZ=<timezone> go run clockWall -port <port>")
		os.Exit(1)
	}

	open_channel := make(chan int)
	for _, arg := range os.Args[1:] {
		connection, err := net.Dial("tcp", arg[strings.LastIndex(arg, "=")+1:])

		if err != nil {
			log.Fatal(err)
		}

		defer connection.Close()
		go io.Copy(os.Stdout, connection)
	}

	a := 1
	a = <-open_channel

	log.Println("Open channel closed. code: ", a)
	close(open_channel)
}
