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
	charCount      int
	lineCount      int
	wordCount      int
	longestLineLen int
}

func (cr *CountResult) SetCharCount(count int) {
	cr.charCount = count
}

func (cr *CountResult) SetWordCount(count int) {
	cr.wordCount = count
}

func (cr *CountResult) SetLineCount(count int) {
	cr.lineCount = count
}

func (cr *CountResult) SetLongestLineLen(len int) {
	cr.longestLineLen = len
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

func (cr *CountResult) GetLongestLineLen() int {
	return cr.longestLineLen
}

func (cr *CountResult) String() string {
	var sb strings.Builder
	sb.WriteString("charCount=" + strconv.Itoa(cr.charCount) + " ")
	sb.WriteString("lineCount=" + strconv.Itoa(cr.lineCount) + " ")
	sb.WriteString("wordCount=" + strconv.Itoa(cr.wordCount) + " ")
	sb.WriteString("longestLineLen=" + strconv.Itoa(cr.longestLineLen) + " ")
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
	lineLen := 0
	for {
		n, err := reader.Read(data)
		if err != nil && err == io.EOF {
			if !wordComplete {
				countResult.wordCount++
			}
			countResult.longestLineLen = max(countResult.longestLineLen, lineLen)
			countResult.lineCount++
			lineLen = 0
			break
		}
		if err != nil {
			return nil, err
		}
		countResult.charCount += n
		for i := 0; i < n; i++ {
			ch := data[i]
			if ch == '\n' {
				countResult.longestLineLen = max(countResult.longestLineLen, lineLen)
				countResult.lineCount++
				lineLen = 0
				if !wordComplete {
					countResult.wordCount++
				}
				wordComplete = true
			} else if ch == ' ' && !wordComplete {
				countResult.wordCount++
				wordComplete = true
				lineLen++
			} else if ch != ' ' && ch != '\n' {
				wordComplete = false
				lineLen++
			} else if ch == ' ' {
				lineLen++
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
