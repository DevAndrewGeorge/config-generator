package plugins

type StaticPlugin struct {}

func NewStaticPlugin() Plugin {
  return &StaticPlugin{}
}

func (s *StaticPlugin) Equal(p Plugin) bool {
  a, ok := p.(*StaticPlugin)
  if ok {
    return *s == *a
  }

  return false
}

func (s *StaticPlugin) Configure(settings map[string]interface{}) error {
  return nil
}
