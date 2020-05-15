package plugins

type MultiPlugin struct {}

func NewMultiPlugin() Plugin {
  return &MultiPlugin{}
}

func (m *MultiPlugin) Equal(o Plugin) bool {
  return Plugin(m) == o
}

func (m *MultiPlugin) Configure(settings map[string]interface{}) error {
  return nil
}
