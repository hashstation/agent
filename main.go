// main.go

package main

import (
	"fmt"
	"github.com/hashstation/agent/utils"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files, err := filepath.Glob(filepath.Join(cwd, "hashcat_files*.zip"))
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "No hashcat files found.")
		os.Exit(1)
	}

	// zipFile := utils.ZipFile{Path: files[0]}
	// zipFile.UnpackTo("test") // cwd

	var hashcatPattern string
	if runtime.GOOS == "windows" {
		hashcatPattern = "hashcat*.exe"
	} else if runtime.GOOS == "linux" {
		hashcatPattern = "hashcat*.bin"
	} else if runtime.GOOS == "darwin" {
		hashcatPattern = "hashcat*.app"
	}

	files, err = filepath.Glob(filepath.Join(cwd, hashcatPattern))
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "No hashcat binary found.")
		os.Exit(1)
	}

	hashcatBinary := files[0]
	fmt.Println(hashcatBinary)

	configFile := "config"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	tlvFileReader := utils.CreateTLVFileReader(configFile)
	err = tlvFileReader.Parse()
	if err != nil {
		panic(err)
	}

	for k, v := range tlvFileReader.Map {
		fmt.Printf("%s: %v\n", k, v)
	}

}
