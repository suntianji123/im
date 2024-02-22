package codec

import "encoding/binary"

// BytesToInt decode packet data length byte to int(Big end)
func BytesToInt(b []byte) int {
	return int(binary.BigEndian.Uint32(b))
}

func IntToBytes(num int) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(num))
	return bytes
}
