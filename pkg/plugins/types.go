package plugins

import (
	"github.com/LSN0WM4N/filessh/pkg/bus"
	"golang.org/x/crypto/ssh"
)

type PluginContext struct {
	Bus     *bus.EventBus
	Session *ssh.Client
	Fd      int
	Data    map[string]interface{}
}

type Viewport struct {
	Height int
	Width  int
}

// Contract of plugins
type Plugin interface {
	ID() string

	Init(ctx PluginContext)
	Render(vp Viewport)

	OnKey(e bus.Event)
	OnEvent(e bus.Event)

	OnFocus()
	OnBlur()

	Destroy()
}

type Registry struct {
	plugins map[string]Plugin
	focused string
}
