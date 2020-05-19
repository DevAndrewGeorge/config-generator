package parser

import (
  log "github.com/sirupsen/logrus"
  "github.com/devandrewgeorge/config-generator/internal/pkg/plugins"
)

func parsePlugins(all_named_plugin_configs map[string]*namedPluginConfig) (map[string]plugins.Plugin, error) {
  all_plugins := map[string]plugins.Plugin{}
  for plugin_name, named_plugin_config := range all_named_plugin_configs {
    plugin, err := named_plugin_config.getPlugin(plugin_name)
    if err != nil {
      log.WithField("scope", "plugins").WithField("name", plugin_name).Error(err.Error())
      return nil, err
    }

    all_plugins[plugin_name] = plugin
  }

  return all_plugins, nil
}
