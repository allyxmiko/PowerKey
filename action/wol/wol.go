package wol

import (
	"fmt"
	"net"
)

func WakeOnLAN(mac string) error {
	mp, err := NewMagicPacket(mac)
	if err != nil {
		return fmt.Errorf("failed to create magic packet: %w", err)
	}

	udpAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
	if err != nil {
		return fmt.Errorf("failed to resolve UDP address: %w", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return fmt.Errorf("failed to dial UDP: %w", err)
	}
	defer conn.Close()

	if _, err := conn.Write(mp.Bytes()); err != nil {
		return fmt.Errorf("failed to send magic packet: %w", err)
	}
	return nil
}
