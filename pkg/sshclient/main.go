package sshclient

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func SetupConnection(user UserConfig) {
	passwordPtr := user.Password
	var auth []ssh.AuthMethod

	if passwordPtr != nil {
		auth = append(auth, ssh.Password(*passwordPtr))
	}

	config := &ssh.ClientConfig{
		User:            *user.Username,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", user.Host, user.Port), config)

	if err != nil {
		panic(err)
	}

	defer client.Close()

	run(client, "ls")
	run(client, "cd Downloads")
	run(client, "ls")
}
