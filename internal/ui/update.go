package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/LSN0WM4N/filessh/internal/events"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.filelist, cmd = m.filelist.Update(msg)
	cmds = append(cmds, cmd)

	m.preview, cmd = m.preview.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case events.WorkingDirectoryChanged:
		childName := msg.ChildName
		actualBasename := m.explorer.PathBasename()

		ok, err := m.explorer.ChangeDir(childName)

		if err != nil || !ok {
			break
		}

		cursor := 0
		newList, err := m.explorer.List(true, true) // TODO: Change hardcoded flags for real config paramters

		if childName == ".." { // Was a goBack
			cursor = firstWithName(newList, actualBasename)
		}

		// Emit message to update the file list
		m.filelist, cmd = m.filelist.Update(events.DirectoryListChanged{
			DirectoryList:  &newList,
			CursorPosition: cursor,
		})

		cmds = append(cmds, cmd)
	case tea.KeyPressMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}
