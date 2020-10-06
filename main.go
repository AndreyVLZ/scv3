package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)
const frameHeaderSize = 10
const (
	// id3SizeLen is length of ID3v2 size format (4 * 0bxxxxxxxx).
	id3SizeLen = 4

	synchSafeMaxSize  = 268435455            // == 0b00001111 11111111 11111111 11111111
	synchSafeSizeBase = 7                    // == 0b01111111
	synchSafeMask     = uint(254 << (3 * 8)) // 11111110 000000000 000000000 000000000

	synchUnsafeMaxSize  = 4294967295           // == 0b11111111 11111111 11111111 11111111
	synchUnsafeSizeBase = 8                    // == 0b11111111
	synchUnsafeMask     = uint(255 << (3 * 8)) // 11111111 000000000 000000000 000000000
)
var ErrInvalidSizeFormat = errors.New("invalid format of tag's/frame's size")

func getByteSlice(size int) []byte {
	return make([]byte, size)
}
func parseSize(data []byte, synchSafe bool) (int64, error) {
	var sizeBase uint
	if synchSafe {
		sizeBase = synchSafeSizeBase
	} else {
		sizeBase = synchUnsafeSizeBase
	}

	var size int64
	for _, b := range data {
		if synchSafe && b&128 > 0 { // 128 = 0b1000_0000
			return 0, ErrInvalidSizeFormat
		}

		size = (size << sizeBase) | int64(b)
	}

	return size, nil
}
func parseSize2(data []byte) int32 {
	size := int32(0)
	for i, b := range data {
		if b&0x80 > 0 {
			fmt.Println("Size byte had non-zero first bit")
		}

		shift := uint32(len(data)-i-1) * 7
		size |= int32(b&0x7f) << shift
	}
	return size
}
func main(){
	file, err := os.Open("file")
	if err != nil {
		log.Fatal(err)
	}
	buf := getByteSlice(32 * 1024)
	fhBuf := buf[:frameHeaderSize]
	if _, err := file.Read(fhBuf); err != nil {
		log.Fatal(err)
	}

	id := string(fhBuf[:4])
	fmt.Println(id)
	fmt.Println(fhBuf[5]&1<<7)

	bodySize, err := parseSize(fhBuf[6:], true)
	fmt.Println(bodySize)
	size := parseSize2(fhBuf[6:])
	fmt.Println(size)
}