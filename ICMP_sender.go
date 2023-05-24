package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"reflect"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const targetIP = "127.0.0.1"

func main() {
	// Creates listener for return packets
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil { // if error, return listen error message
		log.Fatalf("listen error, %s", err)
	}
	defer c.Close()

	// constructs packet
	stringToEncode := "Hi"

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(stringToEncode),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Converts string to binary
	var binaryString string
	for _, c := range stringToEncode {
		binaryString = fmt.Sprintf("%s%.8b", binaryString, c)
	}

	fmt.Println(binaryString)
	fmt.Println(reflect.TypeOf(binaryString))

	// for _ in range binaryString -> if 0 then wait 2 seconds and second BUT if 1 then wait 1
	for value := 0; value < len(binaryString); value++ {
		fmt.Println(binaryString[value])
		if binaryString[value] == 48 { // 48 is ASCII value of 0
			time.Sleep(2 * time.Second)
			fmt.Println("SENT 0")
		} else {
			time.Sleep(1 * time.Second)
			fmt.Println("SENT 1")
		}

		// sends packet to targetIP
		c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(targetIP)}) // NEED TO CONDENSE THIS

	}

	fmt.Println("Sent encoded string:", stringToEncode)

}
