package plugins

type ConsolePlugin struct {
}

func NewConsolePlugin() Plugin {
  return &ConsolePlugin{}
}
