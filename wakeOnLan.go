package main

import (
	"bytes"
	"fmt"
	"net"
	"strings"
)

// WakeOnLAN sends a magic packet to the specified MAC address.
func WakeOnLAN(macAddress string) error {

	macAddress = strings.ReplaceAll(macAddress, "-", ":") // Normalize to colon-separated
	macBytes, err := net.ParseMAC(macAddress)
	if err != nil {
		return fmt.Errorf("invalid MAC address: %w", err)
	}

	magicPacket := bytes.Repeat([]byte{0xff}, 6) // 6 bytes of 0xff
	for i := 0; i < 16; i++ {
		magicPacket = append(magicPacket, macBytes...) // 16 repetitions of the MAC address
	}

	broadcastAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9") // Port 9 is the standard WOL port.
	if err != nil {
		return fmt.Errorf("failed to resolve broadcast address: %w", err)
	}

	conn, err := net.DialUDP("udp", nil, broadcastAddr)
	if err != nil {
		return fmt.Errorf("failed to create UDP connection: %w", err)
	}
	defer conn.Close()

	_, err = conn.Write(magicPacket)
	if err != nil {
		return fmt.Errorf("failed to send magic packet: %w", err)
	}

	return nil
}
