package netty_poll

import (
	"fmt"
	"net"
)

type tcpListen struct {
	NettyPoll
	*net.TCPListener
	connHandel func(conn net.Conn)
}

func NewTcpListen(address string, hanFel func(conn net.Conn)) *tcpListen {
	return &tcpListen{NettyPoll: NettyPoll{address},connHandel: hanFel}
}

func (t *tcpListen) Start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", t.address)
	if err != nil {
		return err
	}
	t.TCPListener, err = net.ListenTCP("tcp", tcpAddr)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("listener err:",err)
			}
		}()

		for {
			conn, err := t.TCPListener.Accept()
			if err != nil {
				fmt.Println("client channel exit",err)
			}
			go t.connHandel(conn)
		}
	}()
	fmt.Println("tcp listen success")
	return nil
}