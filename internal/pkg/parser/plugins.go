package parser

import (
  "errors"
  "github.com/devandrewgeorge/config-generator/internal/pkg/plugins"
)

func parsePlugin(plugin_name string, plugin_type string, config map[string]interface{}) (plugins.Plugin, error) {
  createPluginFunc, ok := plugins.Plugins[plugin_type]
  if !ok {
    msg := "not a valid plugin type"
    log.WithField("name", plugin_type).Error(msg)
    return nil, errors.New(msg)
  }

  plugin := createPluginFunc()
  if plugin.Configure(config) != nil {
    msg := "failed to configure plugin"
    log.WithField("name", plugin_name).Error(msg)
    return nil, errors.New(msg)
  }

  return plugin, nil
}
