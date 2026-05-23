package plugins

func NewRegistry() *Registry {
	return &Registry{plugins: make(map[string]Plugin)}
}

func (r *Registry) Register(p Plugin) {
	r.plugins[p.ID()] = p
}

func (r *Registry) SetFocus(id string) {
	if current, ok := r.plugins[r.focused]; ok {
		current.OnBlur()
	}
	r.focused = id
	if next, ok := r.plugins[id]; ok {
		next.OnFocus()
	}
}

func (r *Registry) Focused() Plugin {
	return r.plugins[r.focused]
}

func (r *Registry) All() []Plugin {
	result := make([]Plugin, 0, len(r.plugins))
	for _, p := range r.plugins {
		result = append(result, p)
	}
	return result
}
