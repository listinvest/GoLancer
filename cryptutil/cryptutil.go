package cryptutil

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"io"
	"log"
	"os"
	"strings"
)

// Generate 2048-bit RSA keypair
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

//Generate random 32-byte AES key
func GenerateAES() []byte {
	aesBytes := make([]byte, 32)

	if _, err := rand.Read(aesBytes); err != nil {
		log.Fatal(err)
	}

	return aesBytes
}

//Encrypt the AES key with OAEP RSA public key
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

//Decrypt the AES key with the RSA private key
func DecryptAESKey(encryptedAESKey []byte, privateKey rsa.PrivateKey) []byte {
	decyptedAESKey, err := privateKey.Decrypt(nil, encryptedAESKey, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		log.Fatal(err)
	}

	return decyptedAESKey
}

func EncryptFile(filename string, extension string, aeskey []byte) {
	// Open target file
	infile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Close file once function has finished execution
	defer infile.Close()

	// Create cipher block using an AES key
	block, err := aes.NewCipher(aeskey)
	if err != nil {
		log.Fatal(err)
	}

	// Create random initialization vector
	initvector := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, initvector); err != nil {
		log.Fatal(err)
	}

	outfile, err := os.OpenFile(filename+extension, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	buf := make([]byte, 1024)
	stream := cipher.NewCTR(block, initvector)

	for {
		n, err := infile.Read(buf)

		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			// Write buffer to file
			outfile.Write(buf[:n])
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
		}

	}
	// Append the initialization vector to end of file
	outfile.Write(initvector)
}

func DecryptFile(filename string, extension string, aeskey []byte) {
	// Open target file
	infile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Close file once function has finished execution
	defer infile.Close()

	// Create cipher block using an AES key
	block, err := aes.NewCipher(aeskey)
	if err != nil {
		log.Fatal(err)
	}

	fileinfo, err := infile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	initvector := make([]byte, block.BlockSize())
	msgLen := fileinfo.Size() - int64(len(initvector))

	_, err = infile.ReadAt(initvector, msgLen)
	if err != nil {
		log.Fatal(err)
	}

	outfile, err := os.OpenFile(strings.Trim(filename, extension), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	buf := make([]byte, 1024)
	stream := cipher.NewCTR(block, initvector)

	for {
		n, err := infile.Read(buf)
		if n > 0 {
			// Account for the initialization vector at the end of the file
			if n > int(msgLen) {
				n = int(msgLen)
			}
			msgLen -= int64(n)

			stream.XORKeyStream(buf, buf[:n])
			outfile.Write(buf[:n])

		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}
}
