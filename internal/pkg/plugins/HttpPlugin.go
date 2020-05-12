package plugins

type HttpPlugin struct {}

func NewHttpPlugin() Plugin {
  return &HttpPlugin{}
}
