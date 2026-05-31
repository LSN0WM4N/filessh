package explorer

import (
	"os"
	"path/filepath"

	"github.com/LSN0WM4N/filessh/pkg/bus"
	"github.com/LSN0WM4N/filessh/pkg/plugins"
)

type ExplorerPlugin struct {
	ctx     plugins.PluginContext
	entries []os.DirEntry
	cwd     string
}

func (e *ExplorerPlugin) ID() string { return "explorer" }

func (e *ExplorerPlugin) Init(ctx plugins.PluginContext) {
	e.ctx = ctx

	cwd := e.ctx.Data["cwd"].(string)
	if len(cwd) == 0 {
		cwd = "/" // TODO: Fix later, !UNSAFE
	}
	e.cwd = cwd
}

func (e *ExplorerPlugin) OnFocus() {
	e.loadDir(e.cwd)
}

func (e *ExplorerPlugin) OnEvent(ev bus.Event) {
	switch ev.Type {
	case EventEnterDir:
		folder := ev.Payload.(string)

		e.cwd = e.cwd + "/" + folder
		e.loadDir(e.cwd)
	case EventGoBack:
		e.cwd = filepath.Dir(e.cwd)
		e.loadDir(e.cwd)
	}
}

func (e *ExplorerPlugin) OnKey(ev bus.Event) {}
func (e *ExplorerPlugin) OnBlur()            {}
func (e *ExplorerPlugin) Destroy()           {}

// Specific methods for explorer plugin

func (e *ExplorerPlugin) Cwd() string            { return e.cwd }
func (e *ExplorerPlugin) Entries() []os.DirEntry { return e.entries }
