package main

import (
	"fmt"
	"log"
	"os"
)
const frameHeaderSize = 10

func getByteSlice(size int) []byte {
	return make([]byte, size)
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
}