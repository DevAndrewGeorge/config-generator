package plugins

type FilePlugin struct {}

func NewFilePlugin() Plugin {
  return &FilePlugin{}
}

func (f *FilePlugin) Equal(o Plugin) bool {
  return Plugin(f) == o
}

func (f *FilePlugin) Configure(settings map[string]interface{}) error {
  return nil
}
