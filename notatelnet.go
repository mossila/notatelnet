package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	host, port := parseHostAndPort(args)
	address := fmt.Sprintf("%s:%s", host, port)
	con, err := net.Dial("tcp", address)

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Connected to %s (%s)", address, con.RemoteAddr())

	defer func(con net.Conn) {
		errC := con.Close()
		if errC != nil {
			log.Fatal(errC)
		}
	}(con)
}
func parseHostAndPort(args []string) (host, port string) {
	host = args[0]
	port = "80"
	if len(args) > 1 {
		port = args[1]
	} else {
		hostAndPort := strings.Split(args[0], ":")
		host = hostAndPort[0]
		if len(hostAndPort) > 1 {
			port = hostAndPort[1]
		}
	}
	return
}
