package wordcounter

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const BufferSize = 4096

type WordCounter struct {
	filePath string
}

type CountResult struct {
	charCount int
	lineCount int
	wordCount int
}

func (cr *CountResult) GetCharCount() int {
	return cr.charCount
}

func (cr *CountResult) GetLineCount() int {
	return cr.lineCount
}

func (cr *CountResult) GetWordCount() int {
	return cr.wordCount
}

func (cr *CountResult) String() string {
	var sb strings.Builder
	sb.WriteString("charCount" + strconv.Itoa(cr.charCount) + " ")
	sb.WriteString("lineCount" + strconv.Itoa(cr.lineCount) + " ")
	sb.WriteString("wordCount" + strconv.Itoa(cr.wordCount) + " ")
	return sb.String()
}

func (wc *WordCounter) Count() (*CountResult, error) {
	filePtr, err := os.Open(wc.filePath)
	if err != nil {
		return nil, err
	}
	defer filePtr.Close()

	reader := bufio.NewReader(filePtr)
	data := make([]byte, 4096) // Do not use more than 4096 bytes at any point of time

	countResult := &CountResult{}
	wordComplete := true
	for {
		n, err := reader.Read(data)
		if err != nil && err == io.EOF {
			if !wordComplete {
				countResult.wordCount++
			}
			countResult.lineCount++
			break
		}
		if err != nil {
			return nil, err
		}
		countResult.charCount += n
		for i := 0; i < n; i++ {
			ch := data[i]
			if ch == '\n' {
				countResult.lineCount++
				if !wordComplete {
					countResult.wordCount++
				}
				wordComplete = true
			} else if ch == ' ' && !wordComplete {
				countResult.wordCount++
				wordComplete = true
			} else if ch != ' ' && ch != '\n' {
				wordComplete = false
			}
		}
	}
	return countResult, nil
}

func NewWordCounter(fPath string) *WordCounter {
	return &WordCounter{
		filePath: fPath,
	}
}
