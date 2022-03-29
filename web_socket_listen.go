package netty_poll

import (
	"net"
	"net/http"
	"golang.org/x/net/websocket"
)

type webSocketListen struct {
	NettyPoll
	connHandel func(conn net.Conn)
}

func NewWebSocketListen(address string,hanDel func(conn net.Conn)) *webSocketListen {
	return &webSocketListen{NettyPoll{address},hanDel}
}

func (w *webSocketListen) Start()  {
	http.Handle("/",websocket.Handler(func(conn *websocket.Conn) {
		w.connHandel(conn)
	}))
	http.ListenAndServe(w.address,nil)
}