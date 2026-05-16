package main

import (
	"os"

	"github.com/LSN0WM4N/filessh/pkg/sshclient"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	host := os.Getenv("SSH_HOST")
	port := os.Getenv("SSH_PORT")
	user := os.Getenv("SSH_USER")
	pass := os.Getenv("SSH_PASS")

	sshclient.SetupConnection(sshclient.UserConfig{
		Host:     host,
		Port:     port,
		Username: &user,
		Password: &pass,
	})
}
