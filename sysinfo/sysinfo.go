package sysinfo

import (
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
	/*
		rootDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		rootDir = filepath.Dir(rootDir)
	*/
	// for testing only
	rootDir := "C:\\test"

	err := filepath.WalkDir(rootDir,
		func(path string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() {
				fileList = append(fileList, path)
			}

			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

	return fileList
}
