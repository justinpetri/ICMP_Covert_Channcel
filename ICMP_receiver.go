package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"golang.org/x/net/icmp"
)

func main() {

	fmt.Println("Waiting for ICMP packets...")
	messageDecoder()

}

func packetCapture() (int, time.Time) {

	// listens for packets
	icmpListener, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatalf("ERROR with targetIP %s", err)
	}

	dataForPackets := make([]byte, 1500)
	capturedPacketSize, sentFromAddress, err := icmpListener.ReadFrom(dataForPackets)
	if err != nil {
		log.Fatalf("ERROR capturing packet data %s", err)
	}

	receivedTime := time.Now()

	if capturedPacketSize == 15 {
		fmt.Printf("%s received START/STOP packet from %s\n", time.Now().Format("15:04:05"), sentFromAddress)
	} else if capturedPacketSize == 8 {
		fmt.Printf("%s received packet from %s\n", time.Now().Format("15:04:05"), sentFromAddress)
	}

	return capturedPacketSize, receivedTime

}

func secondsToBinary() string {

	capturedPacketSize, receivedTime := packetCapture()

	var binaryEncodedString string = ""

	for 1 == 1 {

		if capturedPacketSize == 15 {
			nextCapturedPacketSize, nextReceivedTime := packetCapture()

			if nextCapturedPacketSize == 8 {
				timeDifference := nextReceivedTime.Sub(receivedTime)
				receivedTime = nextReceivedTime

				timeBetweenPackets := int(timeDifference / time.Millisecond)
				if timeBetweenPackets%100 != 0 {
					remainder := timeBetweenPackets % 100
					timeBetweenPackets -= remainder
				}

				if timeBetweenPackets == 500 {
					binaryEncodedString += "0"
				} else if timeBetweenPackets == 1000 {
					binaryEncodedString += "1"
				}

			} else if nextCapturedPacketSize == 15 {
				break

			} else {
				log.Fatalf("ERROR with received ICMP packet after initial startStopMessage packet")
			}

		} else {
			fmt.Println("Received ICMP packet but still waiting for initial startStopMessage packet")
			continue
		}
	}

	return binaryEncodedString
}

func splitBinaryMessage() []string {
	// breaks up secondsToBinary() string and turn it into list
	receivedBinaryMessage := secondsToBinary()
	split := 8
	formattedBinary := []string{}

	for i := 0; i < len(receivedBinaryMessage); i += split {

		end := i + split
		if end > len(receivedBinaryMessage) {
			end = len(receivedBinaryMessage)
		}

		formattedBinary = append(formattedBinary, receivedBinaryMessage[i:end])
	}

	return formattedBinary
}

func messageDecoder() {

	formattedBinary := splitBinaryMessage()

	var decodedMessage string

	for instance := 0; instance < len(formattedBinary); instance += 1 {

		if asciiNumber, err := strconv.ParseInt(formattedBinary[instance], 2, 64); err != nil {
			fmt.Println(err)
		} else {
			decodedMessage += string(asciiNumber)
		}
	}

	fmt.Println("Received decoded message:", decodedMessage)

}
