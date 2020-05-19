package parser

import(
  "testing"
  "github.com/devandrewgeorge/config-generator/internal/pkg/plugins"
)

func TestNamedPluginConfigGetPlugin(t *testing.T) {
  t.Run("no plugin configuration", func(t *testing.T){
    input := &namedPluginConfig{}
    if _, err := input.getPlugin("test"); err == nil {
      t.Fail()
    }
  })

  t.Run("one plugin configuration", func(t *testing.T){
    input := &namedPluginConfig{
      Environment: &environmentPluginConfig{},
    }
    expected := &plugins.EnvironmentPlugin{}
    expected.Configure("test", nil)

    result, err := input.getPlugin("test")

    if err != nil || !result.Equal(expected) { t.Fail() }
  })

  t.Run("multiple plguin configurations", func(t *testing.T){
    input := &namedPluginConfig{
      Environment: &environmentPluginConfig{},
      Console: &consolePluginConfig{},
    }

    if _, err := input.getPlugin("test"); err == nil { t.Fail() }
  })
}
