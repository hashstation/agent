// utils/tlvfilereader.go

package utils

import (
	"bufio"
	"os"
	"strings"
)

type TLVFileReader struct {
	FilePath string
	Map      map[string]string
}

func NewTLVFileReader(filePath string) *TLVFileReader {
	return &TLVFileReader{
		FilePath: filePath,
		Map:      make(map[string]string),
	}
}

func (t *TLVFileReader) Parse() error {
	file, err := os.Open(t.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		if len(parts) < 8 {
			continue
		}

		key := parts[3]
		value := parts[5]

		t.Map[key] = value
	}

	return scanner.Err()
}
