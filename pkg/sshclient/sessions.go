package sshclient

/*
PTY Mode - Shell mode, in this mode you will able to work
exactly as a normal shell, no TUI, no render, just a shell
*/
func PTYMode() {}

/*
Pipe Mode - TUI Mode, in this mode the view content is all
rendered by the TUI engine, in this mode I plan to give you
access to some basic commands, not a whole shell

!IMPORTANT: In this mode, programs such as `vim`, `htop` or
these like will no work, for these use PTY mode
*/
func PipeMode() {}
