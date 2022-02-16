package ssh

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/abiosoft/ishell"
	"github.com/kballard/go-shellquote"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"ssh-server/modules/shell"
	"ssh-server/modules/ssh/rfc_4254"
)

func discardOutOfBounds(reqs <-chan *ssh.Request) {
	for r := range reqs {
		if r.WantReply {
			r.Reply(false, []byte("Discarded"))
		}
		log.Printf("Shell Out of bounds request: '%s'\n", r.Type)
	}
}

func shellSession(ctx context.Context, pl *ServerConn) {
	serverConn := pl.Conn
	go discardOutOfBounds(pl.Reqs)

	ch := <-pl.NewChannel
	if ch.ChannelType() != "session" {
		ch.Reject(ssh.UnknownChannelType, "Untagged session can only be used with 'session's.")
		pl.Conn.Close()
		return
	}

	chn, req, err := ch.Accept()
	if err != nil {
		log.Printf("Error occured in global session: %s", err)
	}

	var sh *ishell.Shell

	var term *terminal.Terminal
	var stdout io.Writer = chn

	isPty := false
	for r := range req {
		pl, err := rfc_4254.ParseRequest(r)
		if err != nil {
			r.Reply(false, []byte(fmt.Sprintf("Couldn't parse payload %s", err)))
			continue
		}

		switch p := pl.(type) {
		case rfc_4254.PtyReq:
			isPty = true
			if r.WantReply {
				r.Reply(true, nil)
			}
			term = terminal.NewTerminal(chn, p.TermModes)
			term.SetSize(int(p.WidthInChars), int(p.HeightInRows))
			stdout = term
		case rfc_4254.Shell:
			sh = shell.Shell(&shell.Session{
				Stdin:      chn,
				Stdout:     stdout,
				Stderr:     chn.Stderr(),
				IsPty:      isPty,
				Context:    ctx,
				ServerConn: serverConn,
			})
			go func() {
				sh.Run()
				chn.SendRequest("exit-status", false, ssh.Marshal(&rfc_4254.ExitStatus{ExitStatus: 0}))
				chn.Close()

			}()
		case rfc_4254.X11Req:
			r.Reply(false, []byte(fmt.Sprintf("%s-not-supported", r.Type)))
		case rfc_4254.Exec:
			sh = shell.Shell(&shell.Session{
				Stdin:      chn,
				Stdout:     stdout,
				Stderr:     chn.Stderr(),
				IsPty:      isPty,
				Context:    ctx,
				ServerConn: serverConn,
			})

			cmd, err := shellquote.Split(p.Command)
			if err != nil {
				r.Reply(false, []byte(err.Error()))
			}
			sh.Process(cmd...)
			chn.SendRequest("exit-status", false, ssh.Marshal(&rfc_4254.ExitStatus{ExitStatus: 0}))
			chn.Close()
		case rfc_4254.WindowChange:
			if term != nil {
				term.SetSize(int(p.WidthColumns), int(p.HeightRows))
			}
		case rfc_4254.ExitSignal:
			sh.Close()
		default:
			r.Reply(false, []byte(fmt.Sprintf("%s-not-supported", r.Type)))
		}
	}
}
