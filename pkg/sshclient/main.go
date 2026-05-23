package sshclient

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func SetupConnection(user UserConfig) (*ssh.Client, error) {
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

	return ssh.Dial("tcp", fmt.Sprintf("%s:%s", user.Host, user.Port), config)
}

func OpenSession(client *ssh.Client) (*ssh.Session, error) {
	return client.NewSession()
}
