package netty_poll

import "net"

type udpListen struct {
	NettyPoll
	*net.UDPConn
	addrHandel func(conn *net.UDPConn)
}

func NewUdpListen(address string, addrHandel func(conn *net.UDPConn)) *udpListen {
	return &udpListen{NettyPoll: NettyPoll{address},addrHandel: addrHandel}
}

func (u *udpListen) Start() error {
	udpAddr, err := net.ResolveUDPAddr("udp", u.address)
	if err != nil {
		return err
	}
	u.UDPConn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	go u.addrHandel(u.UDPConn)
	return nil
}