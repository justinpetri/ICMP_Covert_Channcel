package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	// listens for packets
	fmt.Println("Waiting for ICMP packet...")
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatalf("listen error, %s", err)
	}

	rb := make([]byte, 1500)
	n, peer, err := c.ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}

	rm, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), rb[:n])
	if err != nil {
		log.Fatal(err)
	}

	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		fmt.Printf("got response from %v", peer)
	default:
		break
	}

	// closes ListenPacket
	defer c.Close()

	// decode the message
	encodedICMP := string(rb[8:n])
	fmt.Println("Received encoded string:", encodedICMP)
	decodedICMP, err := base64.StdEncoding.DecodeString(encodedICMP)
	if err != nil {
		fmt.Println("Decoding error")
	}

	// outputs decoded string from received packet
	fmt.Printf("%s\n", decodedICMP)
}
