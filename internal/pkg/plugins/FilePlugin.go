package plugins

type FilePlugin struct {}

func NewFilePlugin() Plugin {
  return &FilePlugin{}
}

func (f *FilePlugin) Equal(p Plugin) bool {
  a, ok := p.(*FilePlugin)
  if ok {
    return *f == *a
  }

  return false
}

func (f *FilePlugin) Configure(settings map[string]interface{}) error {
  return nil
}
