package p2p

import "net"

// RPC represents the 
type RPC struct {
	From    net.Addr
	Payload []byte
}
