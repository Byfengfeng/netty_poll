package netty_poll

import (
	"net"
	"testing"
)

func TestNewWebSocketListen(t *testing.T) {
	webSocketListen := NewWebSocketListen(":7777",func(conn net.Conn) {
		func(){
			go read(conn)
			close := make(chan bool)
			select {
			case <- close:
				return
			}
		}()

	})
	webSocketListen.Start()
}