package main

import (
	"context"
	"os"

	"github.com/LSN0WM4N/filessh/pkg/bus"
	"github.com/LSN0WM4N/filessh/pkg/sshclient"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	host := os.Getenv("SSH_HOST")
	port := os.Getenv("SSH_PORT")
	user := os.Getenv("SSH_USER")
	pass := os.Getenv("SSH_PASS")

	// Stablish a connection
	client, err := sshclient.SetupConnection(sshclient.UserConfig{
		Host:     host,
		Port:     port,
		Username: &user,
		Password: &pass,
	})

	if err != nil {
		panic(err)
	}

	defer client.Close()

	session, err := sshclient.OpenSession(client)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Setup main bus and plugins
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventBus := bus.NewEventBus()

	go eventBus.Run(ctx)

	session, _ = sshclient.PTYMode(session, ctx, eventBus)

	session.Close()
	os.Exit(0)
}
