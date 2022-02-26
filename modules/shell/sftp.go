package shell

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"ssh-server/modules/sftp"
)

func sftpCmd(shell *ishell.Shell, session *Session) {
	shell.AddCmd(&ishell.Cmd{
		Name: "sftp",
		Help: "sftp service",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			c.Print("Please input host: ")
			host, err := c.ReadLineErr()
			if err != nil {
				c.Println(err.Error())
				return
			}
			c.Print("Please input port: ")
			port, err := c.ReadLineErr()
			if err != nil {
				c.Println(err.Error())
				return
			}
			c.Print("Please input username: ")
			user, err := c.ReadLineErr()
			if err != nil {
				c.Println(err.Error())
				return
			}
			c.Print("Please input password: ")
			pass, err := c.ReadLineErr()
			if err != nil {
				c.Println(err.Error())
				return
			}
			c.Print("Please input location save: ")
			fileLocation, err := c.ReadLineErr()
			if err != nil {
				c.Println(err.Error())
				return
			}
			c.Print("Please input fileName: ")
			fileName, err := c.ReadLineErr()
			if err != nil {
				c.Println(err.Error())
				return
			}

			payload := sftp.SFTPConfig{
				User:     user,
				Pass:     pass,
				Address:  fmt.Sprintf("%s:%s", host, port),
				Location: fmt.Sprintf("%s", fileLocation),
				Filename: fileName,
			}

			if err := sftp.UploadFile(payload); err != nil {
				c.Println("Error: " + err.Error())
				return
			}
			c.Printf("File copied.")
		},
	})
}
