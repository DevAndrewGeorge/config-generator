package plugins

type EnvironmentPlugin struct {
}

func NewEnvironmentPlugin() Plugin {
  return &EnvironmentPlugin{}
}
