package plugins

type ConsolePlugin struct {
}

func NewConsolePlugin() Plugin {
  return &ConsolePlugin{}
}

func (c *ConsolePlugin) Equal(o Plugin) bool {
  return Plugin(c) == o
}

func (c *ConsolePlugin) Configure(map[string]interface{}) error {
  return nil
}
