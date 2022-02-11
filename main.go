package main

import (
	"context"
	"log"
	cg "ssh-server/modules/config"
	ssh "ssh-server/modules/ssh"
)

/*
export GOPRIVATE=ssh-server
go mod init ssh-server
*/

type program struct {
	mainContext context.Context
	cancel      context.CancelFunc
}

func LoadConfig() (p program) {
	cg.LoadConfig("config.json") // load config
	p.mainContext, p.cancel = context.WithCancel(context.Background())
	return
}

func main() {
	config := LoadConfig()
	server, err := ssh.NewServer(config.mainContext)
	if err != nil {
		log.Panicf("Can't create server: %s", err)
	}
	server.Listen()
}
