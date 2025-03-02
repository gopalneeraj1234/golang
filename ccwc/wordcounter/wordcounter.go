package counter

import (
	"io"
	"os"
)

const BufferSize = 4096

type WordCounter struct {
	filePath string
}

func (wc *WordCounter) CountChars() (int, error) {
	filePtr, err := os.Open(wc.filePath)
	if err != nil {
		return 0, err
	}
	defer filePtr.Close()

	data := make([]byte, 4096)

	charCount := 0
	for {
		n, err := filePtr.Read(data)
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}
		charCount += n
	}
	return charCount, nil
}

func NewWordCounter(fPath string) *WordCounter {
	return &WordCounter{
		filePath: fPath,
	}
}
