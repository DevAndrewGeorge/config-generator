package plugins

type StaticPlugin struct {}

func NewStaticPlugin() Plugin {
  return &StaticPlugin{}
}

func (s *StaticPlugin) Equal(o Plugin) bool {
  return Plugin(s) == o
}

func (s *StaticPlugin) Configure(settings map[string]interface{}) error {
  return nil
}
