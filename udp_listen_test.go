package netty_poll

import (
	"net"
	"testing"
)

func TestNewUdpListen(t *testing.T) {
	NewUdpListen(":7777", func(conn *net.UDPConn) {
		go func() {
			bytes := make([]byte,1024)
			for  {
				readLen, addr, err := conn.ReadFrom(bytes)

			}

		}()
	})
}
