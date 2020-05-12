package plugins

type FilePlugin struct {}

func NewFilePlugin() Plugin {
  return &FilePlugin{}
}
