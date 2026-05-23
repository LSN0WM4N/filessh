package pty

import (
	"os"

	"github.com/LSN0WM4N/filessh/pkg/bus"
	"github.com/LSN0WM4N/filessh/pkg/plugins"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

// For simplicity i will treat as PTY and Pipe mode as plugins
// That means that they can be disabled and make the whole
// program useless (no one is that dumb, right?)

type TerminalPlugin struct {
	ctx      plugins.PluginContext
	session  *ssh.Session
	rawState *term.State
	active   bool
}

func (t *TerminalPlugin) ID() string { return "terminal" }

func (t *TerminalPlugin) Init(ctx plugins.PluginContext) {
	t.ctx = ctx
}

func (t *TerminalPlugin) OnFocus() {
	t.session, _ = t.ctx.Session.NewSession()

	termType := os.Getenv("TERM")
	if termType == "" {
		termType = "xterm-256color"
	}
	w, h, _ := term.GetSize(t.ctx.Fd)

	t.session.RequestPty(termType, h, w, ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	})

	t.session.Stdout = os.Stdout
	t.session.Stderr = os.Stderr
	t.session.Stdin = os.Stdin

	term.Restore(t.ctx.Fd, t.rawState)

	t.session.Shell()
	t.active = true
}

func (t *TerminalPlugin) OnBlur() {
	if t.session != nil {
		t.session.Close()
		t.session = nil
	}

	t.rawState, _ = term.MakeRaw(t.ctx.Fd)
	t.active = false
}

func (t *TerminalPlugin) OnKey(e bus.Event) {}
func (t *TerminalPlugin) OnEvent(e bus.Event) {
	if e.Type == bus.EventResize {
		info := e.Payload.(bus.ResizeInfo)
		if t.session != nil {
			t.session.WindowChange(info.Height, info.Width)
		}
	}
}
func (t *TerminalPlugin) Render(vp plugins.Viewport) {}
func (t *TerminalPlugin) Destroy()                   { t.OnBlur() }
