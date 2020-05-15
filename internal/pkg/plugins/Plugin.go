package plugins

type ConstructorFunction = func() Plugin
var Plugins = map[string]ConstructorFunction {
  "console": NewConsolePlugin,
  "environment": NewEnvironmentPlugin,
  "file": NewFilePlugin,
  "http": NewHttpPlugin,
  "kubernetes": NewKubernetesPlugin,
  "multi": NewMultiPlugin,
  "script": NewScriptPlugin,
  "static": NewStaticPlugin,
  "vault": NewVaultPlugin,
}

type Plugin interface {
  Equal(Plugin) bool
  Configure(map[string]interface{}) error
}
