package ssh

import (
	"context"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
	cg "ssh-server/modules/config"
)

type Server struct {
	sshConfig ssh.ServerConfig
	ctx       context.Context
}

func NewServer(ctx context.Context) (server *Server, err error) {
	server = &Server{
		sshConfig: ssh.ServerConfig{
			NoClientAuth: !cg.Config.SSH.ClientAuth,
			PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
				// Should use constant-time compare (or better, salt+hash) in a production setting.
				if c.User() == cg.Config.SSH.Username && string(pass) == cg.Config.SSH.Password {
					return nil, nil
				}
				return nil, fmt.Errorf("password rejected for %q", c.User())
			},
		},
		ctx: ctx,
	}
	privateBytes, err := ioutil.ReadFile(cg.Config.SSH.HostKeyFile)
	if err != nil {
		log.Fatalf("Failed to load private key (%s): %s", cg.Config.SSH.HostKeyFile, err)
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %s", err)
	}
	server.sshConfig.AddHostKey(private)
	return
}

func (server *Server) Listen() {
	ip := cg.Config.SSH.Host
	port := cg.Config.SSH.Port

	var lc net.ListenConfig
	listener, err := lc.Listen(server.ctx, "tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		log.Fatalf("Can't open listener at tcp port %s:%d -  %s", ip, port, err)
	}
	log.Printf("Listening on %s:%d", ip, port)
	for {
		tcpConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed accepting ssh (%s)", err)
		} else {
			go server.accept(tcpConn)
		}
	}
}

type ServerConn struct {
	Conn       *ssh.ServerConn
	NewChannel <-chan ssh.NewChannel
	Reqs       <-chan *ssh.Request
}

func (server *Server) accept(tcpConn net.Conn) {
	// Before use, a handshake must be performed on the incoming net.Conn.
	sshConn, chans, reqs, err := ssh.NewServerConn(tcpConn, &server.sshConfig)
	if err != nil {
		log.Printf("Failed handshake (%s)", err)
		return
	}
	log.Printf("New SSH connection from %s (%s)", sshConn.RemoteAddr(), sshConn.ClientVersion())

	pl := &ServerConn{sshConn, chans, reqs}

	// Set up terminal modes
	shellSession(server.ctx, pl)
}
