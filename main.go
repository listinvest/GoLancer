package main

import (
	"crypto/rsa"
	"fmt"

	"github.com/WesEfird/GoLancer/cryptutil"
	"github.com/WesEfird/GoLancer/sysinfo"
)

var pubKey rsa.PublicKey

func main() {
	//DEMO save and load keys
	fmt.Println(sysinfo.GetInfo())
	privateKey := cryptutil.GenerateRSA()
	fmt.Println(*privateKey)

	fmt.Println("Print pub")
	fmt.Println(privateKey.PublicKey)
	fmt.Println("Saving pub...")
	cryptutil.SaveRSAPublicKey(privateKey.PublicKey, "public.pem")
	fmt.Println("Loading pub...")
	pubKey = cryptutil.LoadRSAPublicKey("public.pem")
	fmt.Println(pubKey)

}
