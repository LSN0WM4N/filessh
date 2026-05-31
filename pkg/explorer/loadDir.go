package explorer

import (
	"os"
	"strings"
)

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
