package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

func Listen(port int) error {
	server, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	defer server.Close()

	if err != nil {
		return err
	}

	for {
		connection, err := server.Accept()
		if err != nil {
			return err
		}
		go handle(connection)
	}

	return nil
}

func handle(connection net.Conn) {
	defer connection.Close()
	fmt.Printf("Serving %s\n", connection.RemoteAddr().String())

	line, _, err := bufio.NewReader(connection).ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Print(connection.LocalAddr())
	log.Println(string(line))

}
