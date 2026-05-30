package explorer

import (
	"fmt"
	"os"
	"strings"

	"github.com/LSN0WM4N/filessh/pkg/bus"
	"github.com/LSN0WM4N/filessh/pkg/plugins"
)

type ExplorerPlugin struct {
	ctx     plugins.PluginContext
	entries []os.DirEntry
	cwd     string
	cursor  int
}

type remoteEntry struct {
	name string
}

func (r *remoteEntry) Name() string               { return r.name }
func (r *remoteEntry) IsDir() bool                { return false }
func (r *remoteEntry) Type() os.FileMode          { return 0 }
func (r *remoteEntry) Info() (os.FileInfo, error) { return nil, nil }

func (e *ExplorerPlugin) ID() string { return "explorer" }

func (e *ExplorerPlugin) Init(ctx plugins.PluginContext) {
	fmt.Print("\033[H\033[2J")

	e.ctx = ctx

	cwd := e.ctx.Data["cwd"].(string)
	if len(cwd) == 0 {
		cwd = "/" // TODO: Fix later, !UNSAFE
	}
	e.cwd = cwd
}

func (e *ExplorerPlugin) OnFocus() {
	e.loadDir(e.cwd)
	fmt.Printf("[Explorer Focus]\n")
	e.Render(plugins.Viewport{Height: 0, Width: 0})
}

func (e *ExplorerPlugin) loadDir(path string) {
	session, _ := e.ctx.Session.NewSession()
	defer session.Close()

	out, _ := session.Output("ls -1a " + path)

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	e.entries = make([]os.DirEntry, 0, len(lines))
	for _, line := range lines {
		if len(line) > 0 {
			e.entries = append(e.entries, &remoteEntry{name: line})
		}
	}
}

func (e *ExplorerPlugin) OnKey(ev bus.Event) {
	key := ev.Payload.(bus.KeyInfo)
	switch key.Seq {
	case "down":
		e.cursor = min(e.cursor+1, len(e.entries))
		e.Render(plugins.Viewport{Height: 0, Width: 0})
	case "up":
		e.cursor = max(e.cursor-1, 0)
		e.Render(plugins.Viewport{Height: 0, Width: 0})
	case "enter":
		selected := e.entries[e.cursor]
		if selected.IsDir() {
			e.cwd = e.cwd + "/" + selected.Name()
			e.loadDir(e.cwd)
		}
		e.Render(plugins.Viewport{Height: 0, Width: 0})
	}
}

func (e *ExplorerPlugin) Render(vp plugins.Viewport) {
	// fmt.Println("[Explorer] Render executed")
	// fmt.Println("\033[H\033[2J")
	for i, entry := range e.entries {
		if e.cursor == i {
			fmt.Printf(" -> ")
		} else {
			fmt.Printf(" -  ")
		}
		fmt.Printf("%s\n", entry.Name())
	}
}

func (e *ExplorerPlugin) OnEvent(ev bus.Event) {
	// fmt.Printf("[Explorer] OnEvent triggered: [%s]\n", ev)
}

func (e *ExplorerPlugin) OnBlur()  {}
func (e *ExplorerPlugin) Destroy() {}
