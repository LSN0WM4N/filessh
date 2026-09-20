package cli

import tea "charm.land/bubbletea/v2"

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		list, err := m.explorer.List(true, true)
		if err != nil {
			return "Unable to load the current dir"
		}

		return entriesLoadedMsg{
			entries: list,
		}
	}
}
