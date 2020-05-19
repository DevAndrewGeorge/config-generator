package plugins

type FilePlugin struct {
  name string
}

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

func (f *FilePlugin) Configure(name string, settings map[string]interface{}) error {
  f.name = name
  return nil
}
