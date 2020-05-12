package plugins

type ConsolePlugin struct {
}

func NewConsolePlugin() Plugin {
  return &ConsolePlugin{}
}

func (c ConsolePlugin) Equal(o Plugin) bool {
  return Plugin(c) == o
}
