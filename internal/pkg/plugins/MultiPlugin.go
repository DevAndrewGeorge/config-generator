package plugins

type MultiPlugin struct {}

func NewMultiPlugin() Plugin {
  return &MultiPlugin{}
}
