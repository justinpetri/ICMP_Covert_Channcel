// Works referenced:
// https://blog.logrocket.com/learn-golang-encryption-decryption/
// https://stackoverflow.com/questions/2937123/implementing-icmp-ping-in-go
// https://stackoverflow.com/questions/55631569/go-mod-cannot-find-module-providing-package

// first 8 bits of ICMP packet is reserved. 1 bit for type, 1 for code, 2 for checksum, 2 for identifier, 2 for sequence number

package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const targetIP = "ENTER_TARGET_IP_HERE"

func main() {
	// Creates listener for return packets
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil { // if error, return listen error message
		log.Fatalf("listen error, %s", err)
	}
	defer c.Close()

	// constructs packet
	StringToEncode := "HELLO_CAN_YOU_READ_ME_NOAH"
	Encoding := base64.StdEncoding.EncodeToString([]byte(StringToEncode))
	fmt.Println("Sent Encoded string:", Encoding)

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(Encoding),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}

	// sends packet to targetIP, returns error message if error
	if _, err := c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(targetIP)}); err != nil {
		log.Fatalf("WriteTo err, %s", err)
	}

	// captures return message, returns error message if error
	rb := make([]byte, 1500)
	n, peer, err := c.ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}

	rm, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), rb[:n])
	if err != nil {
		log.Fatal(err)
	}

	// display result
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Printf("got response from %v", peer)
	default:
		log.Printf("got %+v; want echo reply", rm)
	}
}
