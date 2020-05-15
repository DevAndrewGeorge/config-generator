package plugins

type MultiPlugin struct {}

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

func (m *MultiPlugin) Configure(settings map[string]interface{}) error {
  return nil
}
