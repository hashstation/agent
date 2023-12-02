// main.go

package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "github.com/hashstation/agent/utils"
)

func main() {
	hashcatFilesFound := false

    err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if strings.HasPrefix(info.Name(), "hashcat_files") && strings.HasSuffix(info.Name(), ".zip") {
			hashcatFilesFound = true
            //zipFile := utils.ZipFile{Path: path}
            //return zipFile.UnpackTo("test")
        }

        return nil
    })

    if err != nil || !hashcatFilesFound {
        fmt.Fprintln(os.Stderr, "No hashcat files found.")
        os.Exit(1)
    }

    configFile := "config"

    if len(os.Args) > 1 {
        configFile = os.Args[1]
    }

	tlvFileReader := utils.CreateTLVFileReader(configFile)
    err = tlvFileReader.Parse()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error parsing file:", err)
        return
    }

    for k, v := range tlvFileReader.Map {
        fmt.Printf("%s: %v\n", k, v)
    }
}
