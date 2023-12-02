// utils/tlvfilereader.go

package utils

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

type TLVFileReader struct {
    FilePath string
    Map     map[string]interface{}
}

func CreateTLVFileReader(filePath string) *TLVFileReader {
    return &TLVFileReader{
        FilePath: filePath,
        Map:     make(map[string]interface{}),
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
            continue // Skip malformed lines
        }

		key := parts[3]
        dataType := parts[4]
        value := parts[5]

        switch dataType {
        case "String":
            t.Map[key] = value
        case "UInt":
            intValue, err := strconv.Atoi(value)
            if err != nil {
                return err
            }
            t.Map[key] = intValue
        case "BigUInt":
            bigUIntValue, err := strconv.ParseUint(value, 10, 64)
            if err != nil {
                return err
            }
            t.Map[key] = bigUIntValue
        // Add more data types here as needed
        }
    }

    return scanner.Err()
}
