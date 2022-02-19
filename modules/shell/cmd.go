package shell

import (
	"github.com/abiosoft/ishell"
	"os/exec"
	"strings"
)

func executeCMD(shell *ishell.Shell, session *Session) {
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
}
