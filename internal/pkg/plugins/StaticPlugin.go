package plugins

type StaticPlugin struct {
  name string
}

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

func (s *StaticPlugin) Configure(name string, settings map[string]interface{}) error {
  s.name = name
  return nil
}
