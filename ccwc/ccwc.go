package main

import (
	counter "ccwc/wordcounter"
	"flag"
	"fmt"
	"log"
)

func main() {
	userFlags := ProcessFlags()
	filenames := flag.Args()
	if len(filenames) == 0 {
		fmt.Println("No files provided")
		return
	}
	filename := filenames[0]
	wordCounter := counter.NewWordCounter(filename)
	if userFlags.IsCharCount() {
		charCount, err := wordCounter.CountChars()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\t%d %s\n", charCount, filename)
	}
}
