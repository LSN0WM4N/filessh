package sshclient

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func run(client *ssh.Client, cmd string) {
	session, err := client.NewSession()

	if err != nil {
		panic(err)
	}

	defer session.Close()

	var out bytes.Buffer
	session.Stdout = &out

	err = session.Run(cmd)

	if err != nil {
		panic(err)
	}

	fmt.Println(out.String())
}
