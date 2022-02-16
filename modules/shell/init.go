package shell

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"

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
		Name: "cmd",
		Help: "Sent all command",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			if len(c.Args) > 0 {
				cmd := exec.Command(c.Args[0])
				if len(c.Args) > 1 {
					arg1, arg2 := strings.Join(c.Args[:1], " "), strings.Join(c.Args[1:], " ")
					cmd = exec.Command(arg1, arg2)
				}

				output, _ := cmd.CombinedOutput()
				c.Printf("%s\n", string(output))
			}
		},
	})

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
