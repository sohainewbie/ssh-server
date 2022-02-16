package sftp

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

func NewScp() {
	const SSH_ADDRESS = "10.16.5.162:22"
	const SSH_USERNAME = "app"
	const SSH_PASSWORD = "Smartfren@2020"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(SSH_PASSWORD),
		},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}
	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Failed create client sftp client. " + err.Error())
	}

	fDestination, err := sftpClient.Create("/home/app/test-file.txt")
	if err != nil {
		log.Fatal("Failed to create destination file. " + err.Error())
	}

	fSource, err := os.Open("test.txt")
	if err != nil {
		log.Fatal("Failed to read source file. " + err.Error())
	}

	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		log.Fatal("Failed copy source file into destination file. " + err.Error())
	}

	log.Println("File copied.")
}
