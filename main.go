package main

import (
	"log"

	"github.com/Pranav-Agarwal0612/distributed_fs/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		Decoder:       p2p.GOBDecoder{},
		HandshakeFunc: p2p.NOPHandshakeFunc,
	}
	
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal((err))
	}
	select {}
	// fmt.Println("Hello World!")
}
