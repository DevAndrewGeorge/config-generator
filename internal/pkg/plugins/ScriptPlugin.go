package plugins

type ScriptPlugin struct {
  shell string
  user string
  group string
}

func NewScriptPlugin() Plugin {
  return &ScriptPlugin{}
}

func (s ScriptPlugin) Equal(o Plugin) bool {
  return Plugin(s) == o
}
