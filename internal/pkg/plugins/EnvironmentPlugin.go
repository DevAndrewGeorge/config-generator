package plugins

type EnvironmentPlugin struct {
  file string
}

func NewEnvironmentPlugin() Plugin {
  return &EnvironmentPlugin{}
}

func (e EnvironmentPlugin) Equal(o Plugin) bool {
  return Plugin(e) == o
}
