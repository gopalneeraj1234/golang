package main

import (
	"ccwc/wordcounter"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	userFlags := ProcessFlags()
	filenames := flag.Args()
	if len(filenames) == 0 {
		fmt.Println("No files provided")
		return
	}

	sumAllCounts := &wordcounter.CountResult{}
	for _, filename := range filenames {
		wordCounter := wordcounter.NewWordCounter(filename)
		allCounts, err := wordCounter.Count()
		sumAllCounts.SetCharCount(sumAllCounts.GetCharCount() + allCounts.GetCharCount())
		sumAllCounts.SetWordCount(sumAllCounts.GetWordCount() + allCounts.GetWordCount())
		sumAllCounts.SetLineCount(sumAllCounts.GetLineCount() + allCounts.GetLineCount())
		sumAllCounts.SetLongestLineLen(max(sumAllCounts.GetLongestLineLen(), allCounts.GetLongestLineLen()))
		if err != nil {
			log.Fatal(err)
		}

		outputStr := buildOutput(userFlags, filename, allCounts)
		fmt.Println(outputStr)
	}
	fmt.Println(buildOutput(userFlags, "total", sumAllCounts))
}

func buildOutput(userFlags *UserFlags, filename string, allCounts *wordcounter.CountResult) any {
	var sb strings.Builder

	if userFlags.IsLineCountFlag() {
		sb.WriteString(fmt.Sprintf("%8d", allCounts.GetLineCount()))
	}

	if userFlags.IsWordCountFlag() {
		sb.WriteString(fmt.Sprintf("%8d", allCounts.GetWordCount()))
	}

	if userFlags.IsCharCountFlag() {
		sb.WriteString(fmt.Sprintf("%8d", allCounts.GetCharCount()))
	}

	if userFlags.IsLongestLineLenFlag() {
		sb.WriteString(fmt.Sprintf("%8d", allCounts.GetLongestLineLen()))
	}

	sb.WriteString(" ")
	sb.WriteString(filename)
	return sb.String()
}
