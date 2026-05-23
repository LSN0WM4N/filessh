package explorer

import (
	"fmt"
	"os"

	"github.com/LSN0WM4N/filessh/pkg/bus"
	"github.com/LSN0WM4N/filessh/pkg/plugins"
)

type ExplorerPlugin struct {
	ctx     plugins.PluginContext
	entries []os.DirEntry
	cwd     string
	cursor  int
}

func (e *ExplorerPlugin) ID() string { return "explorer" }

func (e *ExplorerPlugin) Init(ctx plugins.PluginContext) {
	e.ctx = ctx
	e.cwd = "/"
}

func (e *ExplorerPlugin) OnFocus() {
	e.loadDir(e.cwd)
}

func (e *ExplorerPlugin) loadDir(path string) {
	session, _ := e.ctx.Session.NewSession()
	defer session.Close()

	out, _ := session.Output("ls -1a " + path)

	fmt.Printf("%s\n", out)
}

func (e *ExplorerPlugin) OnKey(ev bus.Event) {
	key := ev.Payload.(bus.KeyInfo)
	switch key.Seq {
	case "down":
		e.cursor++
	case "up":
		e.cursor--
	case "enter":
		selected := e.entries[e.cursor]
		if selected.IsDir() {
			e.cwd = e.cwd + "/" + selected.Name()
			e.loadDir(e.cwd)
		}
	}
}

func (e *ExplorerPlugin) Render(vp plugins.Viewport) {
	for i, entry := range e.entries {
		// dibujar cada entrada en su línea dentro del viewport
		_ = i
		_ = entry
	}
}

func (e *ExplorerPlugin) OnBlur()              {}
func (e *ExplorerPlugin) OnEvent(ev bus.Event) {}
func (e *ExplorerPlugin) Destroy()             {}
