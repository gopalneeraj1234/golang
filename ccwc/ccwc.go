package main

import (
	"ccwc/wordcounter"
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
	wordCounter := wordcounter.NewWordCounter(filename)
	allCounts, err := wordCounter.Count()
	if err != nil {
		log.Fatal(err)
	}

	if userFlags.IsCharCount() {
		fmt.Printf("\t%d %s\n", allCounts.GetCharCount(), filename)
	}

	if userFlags.IsLineCount() {
		fmt.Printf("\t%d %s\n", allCounts.GetLineCount(), filename)
	}

	if userFlags.IsWordCount() {
		fmt.Printf("\t%d %s\n", allCounts.GetWordCount(), filename)
	}
}
