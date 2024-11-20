package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client Test...Start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return 
	}

	for {
		conn.Write([]byte("hello Zinx"))
		buf := make([]byte, 512)
		if cnt, _ := conn.Read(buf); cnt != 0 {
			fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)
		}
		time.Sleep(time.Second * 1)
	}
}