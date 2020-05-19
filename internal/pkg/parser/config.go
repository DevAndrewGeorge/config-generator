package parser

import(
  "fmt"
  "strings"
  "reflect"
  "github.com/fatih/structs"
  "github.com/devandrewgeorge/config-generator/internal/pkg/errors"
  "github.com/devandrewgeorge/config-generator/internal/pkg/plugins"
)

type GeneratorConfig struct {
  Plugins   map[string]*namedPluginConfig
  Variables map[string]map[string]interface{} // *NamedVariable
  Templates map[string]interface{} // *NamedTemplate
  Outputs   map[string]map[string]interface{} // *NamedOutput
}

type namedPluginConfig struct {
  Console     *consolePluginConfig
  Environment *environmentPluginConfig
  File        *filePluginConfig
  Http        *httpPluginConfig
  Kubernetes  *kubernetesPluginConfig
  Multi       *multiPluginConfig
  Script      *scriptPluginConfig
  Static      *staticPluginConfig
  Vault       *vaultPluginConfig
}

type consolePluginConfig struct { Test *string }

type environmentPluginConfig struct {
  File *string
}

type filePluginConfig struct {}

type httpPluginConfig struct {}


type kubernetesPluginConfig struct {
  Kubeconfig *string
}

type multiPluginConfig struct {}

type scriptPluginConfig struct {
  Shell *string
  User  *string
  Group *string
}

type staticPluginConfig struct {}

type vaultPluginConfig struct {
  Token       *string
  Address     *string
  CAcert      *string `yaml:"ca_cert"`
  CApath      *string `yaml:"ca_path"`
  SkipVerify  *bool   `yaml:"skip_verify"`
}

func (named_plugin *namedPluginConfig) getPlugin(plugin_name string) (plugins.Plugin, error) {
  named_plugin_config := structs.New(named_plugin).Map()

  num_configurations := 0
  var plugin plugins.Plugin
  for plugin_type, plugin_config := range named_plugin_config {
    if reflect.ValueOf(plugin_config).IsNil() { continue }

    num_configurations++
    if num_configurations > 1 {
      return nil, &errors.PluginError{Message: "more than one plugin type declared"}
    }

    newFunc, found := plugins.Plugins[strings.ToLower(plugin_type)]

    if !found {
      return nil, &errors.PluginError{Message: fmt.Sprintf("%s is not a valid plugin type", plugin_type)}
    }

    plugin = newFunc()
    var err error
    if input_config, ok := plugin_config.(map[string]interface{}); ok {
      err = plugin.Configure(plugin_name, input_config)
    } else {
      err = plugin.Configure(
        plugin_name,
        structs.New(plugin_config).Map(),
      )
    }

    if err != nil { return nil, err }
  }

  if num_configurations == 0 {
    return nil, &errors.PluginError{Message: "no plugin type declared"}
  }

  return plugin, nil
}
