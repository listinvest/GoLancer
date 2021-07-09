package sysinfo

import (
	"log"
	"os/exec"
	"runtime"
)

type SystemInfo struct {
	Os       string
	Arch     string
	Hostname string
}

func GetInfo() SystemInfo {
	sysInfo := SystemInfo{
		Os:       runtime.GOOS,
		Arch:     runtime.GOARCH,
		Hostname: getHostName(),
	}
	return sysInfo
}

func getHostName() string {
	out, err := exec.Command("hostname").Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
