package main

import (
	"crypto/rsa"
	"fmt"

	"github.com/WesEfird/GoLancer/cryptutil"
	"github.com/WesEfird/GoLancer/sysinfo"
)

var privateKey *rsa.PrivateKey
var publicKey rsa.PublicKey
var aesKey []byte
var encryptedAES []byte
var decryptedAES []byte

func main() {
	fmt.Println(sysinfo.GetInfo())
	privateKey = cryptutil.GenerateRSA()
	publicKey = privateKey.PublicKey
	aesKey = cryptutil.GenerateAES()
	encryptedAES = cryptutil.EncryptAESKey(aesKey, publicKey)
	decryptedAES = cryptutil.DecryptAESKey(encryptedAES, *privateKey)

	//DEMO
	fmt.Println(*privateKey)
	fmt.Println("Pub:")
	fmt.Println(publicKey)
	fmt.Println("AES:")
	fmt.Println(aesKey)
	fmt.Println("Encrypted AES:")
	fmt.Println(encryptedAES)
	fmt.Println("Decrypted AES:")
	fmt.Println(decryptedAES)
	cryptutil.EncryptFile("test.txt", ".lcr", aesKey)
	fmt.Println("Encrypted: test.txt")
	cryptutil.DecryptFile("test.txt.lcr", ".lcr", aesKey)
	fmt.Println("Decrypted file")
}
