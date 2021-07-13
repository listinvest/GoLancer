# GoLancer

Ransomware PoC implementation built in Golang.<nl>
<nl>
- Encrypts files using 256-bit AES-CTR with random initialization vectors per file.<nl>
- Encryption (and decryption) runs in parallel if enough logical cores are availible. List of files are split into chunks and each chunk is encrypted (or decrypted) in it's own goRoutine.<nl>
- AES key is encrypted with 2048 OAEP RSA public key and sent to a webserver via form POST request.

  
  

<h2>Usage</h2>
<h3>Flags</h3>

`-g` : Generates RSA keypair and saves them to the disk.

`-kill` : Starts the encryption process, crawls user directories and encrypts (and deletes) all files

`-d` : Starts the decryption process. The private key must be in the same directory as the GoLancer binary. Will decrypt files, then remove the encrypted files.

`-a` : Decrypts AES key using RSA public key.

<h3>Operation</h3>

The 'attacker' will need a web-server setup that accepts form posts requests. GoLancer will send out a form POST request (``application/x-www-form-urlencoded``) to a defined web-address that contains the fields `hostname` and `key` . Any test web-server will work, you could even use netcat (`nc -lvnp 80`) . The 'attacker' can also use something like webhook.site to accept these requests. (Although security controls may block this site)

On the 'attacker' machine:

 1. Build binary for attacker's OS and arch
	 1. `go build github.com/WesEfird/GoLancer` 
 2. Generate RSA key-pair (This will create files `private.pem` and `public.pem`)
	 1. `./GoLancer -g` 
 3.  Save the private key somewhere safe
 4. Edit `addr` var in `main.go` to the domain or IP address of your webserver. (Or use something like webhook.site to accept POST requests)
 5. Build binary for target's OS and arch
	 1. `env GOOS=target-OS GOARCH=target-architecture go build github.com/WesEfird/GoLancer` 
	 2. https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
 6. Deliver binary and `public.pem` to target machine
 7. Wait for binary to be executed on target machine (see below)
	 1. Once the encryption has completed, a POST request will be made to the web-address defined in step 4
 8. Grab AES key from POST request made to your web-server
	 1. Copy key to file named `golancer-e.key` (Make sure file is written with ANSI encoding)
 9. Decrypt AES key
	 1. `./GoLancer -a ` (Make sure the file is in the same directory as GoLancer binary)
	 2. This will create the file `golancer.key` which will contain the decrypted AES key
 10. Now the 'attacker' has the key that will allow the 'target' to decrypt their file system
	 1. Deliver decrypted AES key to the 'target'
 
<nl>

On the 'target' machine:

 1. Once the GoLancer binary and generated RSA public key has been delivered to the 'target', start the encryption process.
	 1. `./GoLancer -kill`
	 2. This will generate a list of all files in the 'target' machine's home directories.
	 3. The AES key will be encrypted using the RSA public key, and will be sent to the defined web-server via POST request.
 2. The encryption process will start, and all files (except for GoLancer files) in the 'target' machine's home directories will be encrypted using 256-bit AES-CTR
 3. Have the decrypted AES key delivered to the 'target' machine
	 1. Move the AES key file `golancer.key` to the directory where the GoLancer binary is located
 4. Decrypt the file system
	 1.  `./GoLancer -d`
	 2. The filesystem will be decrypted, and all encrypted artifacts will be removed


<h3>Build</h3>

 1. `git clone https://github.com/WesEfird/GoLancer.git`
 2. `cd GoLancer`
 3. `go build github.com/WesEfird/GoLancer`



<h2>TODO</h2>

 - Clean AES key from memory after encryption has completed
 - Stop being lazy about error handling
 - Generate ransom note (Or even a cool webpage??)
 - Remove all traces of GoLancer once decryption has taken place
 - Implement data exfiltration

 
<h2>Disclaimer</h2>

The purpose of GoLancer is to allow the study of malware and enable security researchers to have access to a live malware implementation; this implementation is to be used as a tool for educational purposes, and for the development of defensive rules, tactics, and techniques. This program is intended to only be used in environments that the user owns and controls, or in environments where the user has explicit permission to run offensive security tools. The user must adhere to all laws in their jurisdiction and must conduct themselves ethically when using this tool.

**This tool may cause irreversible changes to your system(s), be extremely careful when running this tool.**
