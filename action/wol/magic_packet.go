package wol

import (
	"fmt"
	"net"
)

type MacAddr [6]byte

type MagicPacket struct {
	header  [6]byte
	payload [16]MacAddr
}

func NewMagicPacket(mac string) (*MagicPacket, error) {
	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		return nil, fmt.Errorf("invalid MAC format: %w", err)
	}

	if len(hwAddr) != 6 {
		return nil, fmt.Errorf("invalid MAC length: got %d bytes (need 6)", len(hwAddr))
	}

	return &MagicPacket{
		header:  [6]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		payload: createPayload(*(*MacAddr)(hwAddr)),
	}, nil
}

func createPayload(addr MacAddr) (payload [16]MacAddr) {
	for i := range payload {
		payload[i] = addr
	}
	return
}

func (mp *MagicPacket) Bytes() []byte {
	buf := make([]byte, 0, 6+16*6) // 预分配精确长度
	buf = append(buf, mp.header[:]...)
	for _, mac := range mp.payload {
		buf = append(buf, mac[:]...)
	}
	return buf
}
