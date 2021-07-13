package sysinfo

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type SystemInfo struct {
	Os       string
	Arch     string
	Hostname string
}

func GetInfo() SystemInfo {
	hostname, _ := os.Hostname()
	sysInfo := SystemInfo{
		Os:       runtime.GOOS,
		Arch:     runtime.GOARCH,
		Hostname: hostname,
	}
	return sysInfo
}

func GetFileList() []string {
	var fileList []string

	rootDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	rootDir = filepath.Dir(rootDir)

	err = filepath.WalkDir(rootDir,
		func(path string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() && d.Type() != os.ModeSymlink {

				fileList = append(fileList, path)
			}

			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

	return fileList
}

func SaveFileList(fileList []string, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, files := range fileList {
		_, err = fmt.Fprintln(file, files)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

}

func LoadFileList(fileName string) []string {
	var files []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := bufio.NewScanner(file)

	for buf.Scan() {
		files = append(files, buf.Text())
	}

	return files
}
