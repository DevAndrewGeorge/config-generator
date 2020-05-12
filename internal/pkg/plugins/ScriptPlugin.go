package plugins

type ScriptPlugin struct {}

func NewScriptPlugin() Plugin {
  return &ScriptPlugin{}
}
