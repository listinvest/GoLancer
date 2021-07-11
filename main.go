package main

import (
	"fmt"

	"github.com/WesEfird/GoLancer/sysinfo"
)

func main() {
	//DEMO save and load keys
	fmt.Println(sysinfo.GetInfo())
	//sysinfo.GetFileList()
	fmt.Println(sysinfo.GetFileList())
	//fmt.Println(os.PathSeparator)
}
