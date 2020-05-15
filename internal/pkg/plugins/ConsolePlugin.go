package plugins

type ConsolePlugin struct {
}

func NewConsolePlugin() Plugin {
  return &ConsolePlugin{}
}

func (c *ConsolePlugin) Equal(p Plugin) bool {
  a, ok := p.(*ConsolePlugin)
  if ok {
    return *c == *a
  }

  return false
}

func (c *ConsolePlugin) Configure(map[string]interface{}) error {
  return nil
}
