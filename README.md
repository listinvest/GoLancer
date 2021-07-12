# GoLancer
Ransomware PoC implementation built in Golang.

Encrypts files using 256-bit AES-CTR with random nonces
Encryption (and decryption) runs in parallel if enough logical cores are availible. Files are split into chunks and each chunk is encrypted (or decrypted) in it's own goRoutine.
AES key is encrypted with 2048 OAEP RSA public key


<h2>Usage</h2>
-g : Generates RSA keypair and saves them to the disk.
-kill : Starts the encryption process, crawls user directories and encrypts (and deletes) all files
-d : Starts the decryption process. The private key must be in the same directory as the GoLancer binary. Will decrypt files, then remove the encrypted files.
<h2>Disclaimer</h2>

The purpose of GoLancer is to allow the study of malware and enable security researchers to have access to a live malware implementation; this implementation is to be used as a tool for educational purposes, and for the development of defensive rules, tactics, and techniques. This program is intended to only be used in environments that the user owns and controls, or in environments where the user has explicit permission to run offensive security tools. The user must adhere to all laws in their jurisdiction and must conduct themselves ethically when using this tool.

This tool may cause irreversible changes to your system(s), be extremely careful when running this tool.
