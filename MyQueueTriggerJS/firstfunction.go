package main

import (
	"log"
	_ "net/http/pprof"
	"os"
	"strconv"
)

func main() {
	inputMessage := os.Getenv("inputMessage")
	outputMessageLocation := os.Getenv("outputQueueItem")
	if inputMessage != "" {
		cacheFile, err := os.OpenFile(outputMessageLocation, os.O_WRONLY|os.O_CREATE, 0)
		if err != nil {
			log.Fatalf("Unable to open file %s: %s\n", outputMessageLocation, err)
		}
		s := strconv.Itoa(len(inputMessage))
		cacheFile.WriteString(s)
	}
}
