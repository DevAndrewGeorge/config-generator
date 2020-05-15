package plugins

type HttpPlugin struct {}

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

func (h *HttpPlugin) Configure(settings map[string]interface{}) error {
  return nil
}
