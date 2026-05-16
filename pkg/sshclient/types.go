package sshclient

type UserConfig struct {
	Host string
	Port string

	Username *string
	Password *string
}
