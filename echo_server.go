package main

import (
	"fmt"
	"io"
	"log"
	"net"
)
// START OMIT
const listenAddr = "localhost:4000"

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening on localhost port 4000 ...")
	for {
		c, err := l.Accept()
		fmt.Println("accepted a connection")
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(c, c)
	}
}
// END OMIT
