package plugins

type ConsolePlugin struct {
  name string
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

func (c *ConsolePlugin) Configure(name string, settings map[string]interface{}) error {
  c.name = name
  return nil
}
