package plugins

type EnvironmentPlugin struct {
  file string
}

func NewEnvironmentPlugin() Plugin {
  return &EnvironmentPlugin{}
}

func (e *EnvironmentPlugin) Equal(o Plugin) bool {
  return Plugin(e) == o
}

func (e *EnvironmentPlugin) Configure(settings map[string]interface{}) error {
  if file, ok := settings["file"]; ok {
    e.file = file.(string)
  }

  return nil
}
