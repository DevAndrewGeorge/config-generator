package plugins

type EnvironmentPlugin struct {
  file string
}

func NewEnvironmentPlugin() Plugin {
  return &EnvironmentPlugin{}
}

func (e *EnvironmentPlugin) Equal(p Plugin) bool {
  a, ok := p.(*EnvironmentPlugin)
  if ok {
    return *e == *a
  }

  return false
}

func (e *EnvironmentPlugin) Configure(settings map[string]interface{}) error {
  if file, ok := settings["file"]; ok {
    e.file = file.(string)
  }

  return nil
}
