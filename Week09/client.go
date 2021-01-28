package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("client is running")
	times := 1
	for {
		msg := "hello kitty"
		fmt.Fprintf(conn, msg+"\n")
		fmt.Printf("send: %v\n", msg)
		reply, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("receive: %v\n", reply)
		times++
		time.Sleep(time.Second)
	}
}
