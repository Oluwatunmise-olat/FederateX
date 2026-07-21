package main

import (
	"fmt"
	"log"
)

func main() {
	tun, err := CreateTUN("federatex0")

	if err != nil {
		log.Fatal(err)
	}

	defer tun.Close()

	fmt.Println("Created TUN interface: federatex0")

	fmt.Println("waiting for packets... (press Ctrl+C to stop)")

	maxPacketSize := 1500

	packetBuffer := make([]byte, maxPacketSize)

	for {
		n, err := tun.Read(packetBuffer)

		if err != nil {
			log.Fatalf("read from tun: %v", err)
		}

		packet := packetBuffer[:n]

		version := packet[0] >> 4

		protocol := packet[9]

		srcIP := fmt.Sprintf("%d.%d.%d.%d", packet[12], packet[13], packet[14], packet[15])

		dstIP := fmt.Sprintf("%d.%d.%d.%d", packet[16], packet[17], packet[18], packet[19])

		fmt.Printf("IPv%d | From: %s -> To: %s | protocol=%d | %d bytes\n", version, srcIP, dstIP, protocol, n)
	}
}
