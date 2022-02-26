package sftp

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

type SFTPConfig struct {
	User     string
	Pass     string
	Address  string
	Location string
	Filename string
}

func UploadFile(payload SFTPConfig) (err error) {
	sshConfig := &ssh.ClientConfig{
		User:            payload.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(payload.Pass),
		},
	}

	client, err := ssh.Dial("tcp", payload.Address, sshConfig)
	if client != nil {
		defer client.Close()
	}
	if err != nil {
		log.Printf("Failed to dial. " + err.Error())
		return
	}
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Printf("Failed create client sftp client. " + err.Error())
		return
	}

	log.Printf("payload.Location: " + payload.Location)
	fDestination, err := sftpClient.Create(payload.Location)
	if err != nil {
		log.Printf("Failed to create destination file. " + err.Error())
		return
	}
	fSource, err := os.Open(payload.Filename)
	if err != nil {
		log.Printf("Failed to read source file. " + err.Error())
		return
	}

	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		log.Printf("Failed copy source file into destination file. " + err.Error())
		return
	}
	return
}
