package plugins

type HttpPlugin struct {
  name string
}

func NewHttpPlugin() Plugin {
  return &HttpPlugin{}
}

func (h *HttpPlugin) Equal(p Plugin) bool {
  a, ok := p.(*HttpPlugin)
  if ok {
    return *h == *a
  }

  return false
}

func (h *HttpPlugin) Configure(name string, settings map[string]interface{}) error {
  h.name = name
  return nil
}
