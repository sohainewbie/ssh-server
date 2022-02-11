package shell

import (
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
	"golang.org/x/crypto/ssh"

	cg "ssh-server/modules/config"
)

type Session struct {
	Stdin       io.ReadCloser
	Stdout      io.Writer
	Stderr      io.Writer
	IsPty       bool
	HistoryFile string
	ServerConn  *ssh.ServerConn

	Context context.Context
}

func Shell(session *Session) *ishell.Shell {

	shell := ishell.NewWithConfig(&readline.Config{
		Prompt:              fmt.Sprintf("%s$ ", session.ServerConn.User()),
		HistoryFile:         session.HistoryFile,
		StdinWriter:         session.Stdout,
		Stdin:               session.Stdin,
		Stdout:              session.Stdout,
		Stderr:              session.Stderr,
		ForceUseInteractive: session.IsPty,
	})

	if session.IsPty {
		shell.Println(cg.Config.SSH.TextDisplay)
	}

	shell.AddCmd(&ishell.Cmd{
		Name: "ls",
		Help: "",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			cmd := exec.Command("ls", "-al")
			output, _ := cmd.CombinedOutput()

			// c.Print("Please insert the password: ")
			// password, err := c.ReadLineErr()
			// if err != nil {
			// 	c.Println(err.Error())
			// 	return
			// }
			c.Printf("%s\n", string(output))

		},
	})

	// addKeyCmds(shell, session)
	// addUserCmds(shell, session)

	shell.Interrupt(
		func(c *ishell.Context, count int, line string) {
			if count >= 2 {
				c.Println("Interrupted")
				c.Stop()
			}
			c.Println("Input Ctrl-c once more to exit")
		})

	return shell
}
