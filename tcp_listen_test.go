package netty_poll

import (
	"fmt"
	"net"
	"testing"
)

func TestNewTcpListen(t *testing.T) {
	tcpListen := NewTcpListen(":7777", func(conn net.Conn) {
		go read(conn)
	})
	tcpListen.Start()

	<- make(chan struct{})
}

func read(conn net.Conn)  {
	bytes := make([]byte,1024)

	for  {
		readLen, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("tcp read exits")
			conn.Close()
			return
		}
		if readLen == 0 {
			fmt.Println("tcp read exits")
			conn.Close()
			return
		}
		fmt.Println("收到消息：",string(bytes[:readLen]))
	}
}
