package main

import (
	"fmt"
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

	session = sshclient.SetupShell(session)

	// fmt.Fprintln(os.Stdin, "pwd")
	// fmt.Fprintln(stdin, "cd /tmp")
	// fmt.Fprintln(stdin, "pwd")
	// fmt.Fprintln(stdin, "ls")

	// keep alive
	fmt.Println("ENTER para salir")
	fmt.Scanln()

	session.Close()
	os.Exit(0)
}
