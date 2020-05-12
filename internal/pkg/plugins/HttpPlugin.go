package plugins

type HttpPlugin struct {}

func NewHttpPlugin() Plugin {
  return &HttpPlugin{}
}

func (h HttpPlugin) Equal(o Plugin) bool {
  return Plugin(h) == o
}
