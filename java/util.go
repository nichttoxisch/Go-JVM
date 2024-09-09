package java

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

var __nbi = 0

func NextBytes(bytes []byte, n int, reset ...bool) []byte {
	if len(reset) > 0 {
		__nbi = 0
	}

	__nbi += n
	return bytes[__nbi-n : __nbi]
}

func ReadFile(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func ToInt(bytes []byte) int {
	switch len(bytes) {
	case 1:
		return int(bytes[0])
	case 2:
		return int(binary.BigEndian.Uint16(bytes))
	case 4:
		return int(binary.BigEndian.Uint32(bytes))
	case 8:
		return int(binary.BigEndian.Uint64(bytes))
	}

	panic(fmt.Sprintf("ERROR: Unknown byte array length: %v", len(bytes)))
}
