package p2p

import "net"

type HandshakeFunc func(net.Conn) error

func NOPHandshakeFunc(net.Conn) error { return nil }
