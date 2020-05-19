package plugins

type MultiPlugin struct {
  name string
}

func NewMultiPlugin() Plugin {
  return &MultiPlugin{}
}

func (m *MultiPlugin) Equal(p Plugin) bool {
  a, ok := p.(*MultiPlugin)
  if ok {
    return *m == *a
  }

  return false
}

func (m *MultiPlugin) Configure(name string, settings map[string]interface{}) error {
  m.name = name
  return nil
}
