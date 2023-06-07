package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const targetIP = "127.0.0.1"

func main() {

	stringToEncode := "hi"
	fmt.Println("\nSending:", stringToEncode)

	binaryString := convertToBinary(stringToEncode)
	sendICMPPackets(binaryString)

}

func createICMPPackets() (*icmp.PacketConn, []byte) {

	icmpListener, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatalf("ERROR with targetIP %s", err)
	}

	startStopMessage := "message"
	icmpPacket := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(startStopMessage),
		},
	}

	packetBinaryEncoding, err := icmpPacket.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}

	return icmpListener, packetBinaryEncoding

}

func convertToBinary(stringToEncode string) string {

	var binaryString string
	for _, c := range stringToEncode {
		binaryString = fmt.Sprintf("%s%.8b", binaryString, c)
	}

	fmt.Println(binaryString)
	return binaryString
}

func sendICMPPackets(binaryString string) {
	icmpListener, packetBinaryEncoding := createICMPPackets()

	// START PACKET
	icmpListener.WriteTo(packetBinaryEncoding, &net.IPAddr{IP: net.ParseIP(targetIP)})

	for value := 0; value < len(binaryString); value++ {
		if binaryString[value] == 48 { // 48 is ASCII value of 0
			time.Sleep(500 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
		}

		sentTime := time.Now().Format("15:04:05")
		fmt.Printf("%s Sent %v\n", sentTime, string(binaryString[value]))

		// sends packet to targetIP
		icmpListener.WriteTo(make([]byte, 8), &net.IPAddr{IP: net.ParseIP(targetIP)}) // Allocates 8 bytes for "Echo Request" Type to send empty ICMP packet

	}

	// STOP PACKET
	time.Sleep(1 * time.Second)
	icmpListener.WriteTo(packetBinaryEncoding, &net.IPAddr{IP: net.ParseIP(targetIP)}) // need to add specific byte number

	fmt.Println("Sent encoded string")
}
