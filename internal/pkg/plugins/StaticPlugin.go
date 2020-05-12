package plugins

type StaticPlugin struct {}

func NewStaticPlugin() Plugin {
  return &StaticPlugin{}
}
