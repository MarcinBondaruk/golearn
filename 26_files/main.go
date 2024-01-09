package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"a.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	bs := []byte{97, 98, 99}

	bytesWritten, err := bufferedWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("bytes written to file: %d\n", bytesWritten)

	bytesAvailable := bufferedWriter.Available()

	log.Printf("bytes available: %d\n", bytesAvailable)

	_, err = bufferedWriter.WriteString("my random string")

	if err != nil {
		log.Fatal(err)
	}

	unflushedBufferSize := bufferedWriter.Buffered()

	log.Printf("unflushed bytes: %d\n", unflushedBufferSize)
	bufferedWriter.Flush()

	file2, err := os.Open("readfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
}
