package cryptutil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

func GenerateRSA() *rsa.PrivateKey {
	rsaKeys, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return rsaKeys
}

func GenerateAES() []byte {
	aesBytes := make([]byte, 32)

	if _, err := rand.Read(aesBytes); err != nil {
		log.Fatal(err)
	}

	return aesBytes
}

func EncryptAESKey(aesBytes []byte, pubkey rsa.PublicKey) []byte {
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&pubkey,
		aesBytes,
		nil)
	if err != nil {
		log.Fatal(err)
	}

	return encryptedBytes
}

func DecryptAESKey(encryptedAESKey []byte, privateKey rsa.PrivateKey) []byte {
	decyptedAESKey, err := privateKey.Decrypt(nil, encryptedAESKey, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		log.Fatal(err)
	}

	return decyptedAESKey
}
