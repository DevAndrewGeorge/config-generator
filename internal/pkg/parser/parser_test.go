package parser

import (
  "fmt"
  "testing"
  "strings"
  "github.com/devandrewgeorge/config-generator/internal/pkg/generator"
)

func TestParse (t *testing.T) {
  createConfig := func(plugins []string, variables []string, templates []string, outputs []string) []byte {
    return []byte(
      fmt.Sprintf(
        "plugins:\n%s\nvariables\n%s\ntemplates:\n%s\noutputs:\n%s\n",
        strings.Join(plugins,  "\n"),
        strings.Join(variables, "\n"),
        strings.Join(templates, "\n"),
        strings.Join(outputs, "\n"),
      ),
    )
  }

  default_generator := *generator.New()
  examples := map[string]string {
    "plugins_default_override": "\tenvironment:\n\t\tenvironment:\n\t\t\tfile: test",
    "plugins_custom_one": "\ttest1:\n\t\tenvironment:\n\t\t\tfile: hello",
    "plugins_custom_multiple": "\ttest1:\n\t\tenvironment\n\t\t\tfile: hello\n\ttest2:\n\t\tenvironment",
  }

  t.Run("config: nil", func(t *testing.T) {
    g, e := Parse(nil)
    if e != nil || g == nil || (*g).Equal(default_generator) {
      t.Fail()
    }
  })

  t.Run("config: empty", func(t *testing.T) {
    g, e := Parse([]byte{})
    if e != nil || g == nil || (*g).Equal(default_generator) {
      t.Fail()
    }
  })

  t.Run("config: basic", func(t *testing.T) {
    g, e := Parse(createConfig([]string{}, []string{}, []string{}, []string{}))
    if e != nil || g == nil || (*g).Equal(default_generator) {
      t.Fail()
    }
  })

  t.Run("config: invalid yaml (formatting)", func (t *testing.T) {
    _, e := Parse([]byte("\toffset"))
    if e == nil {
      t.Fail()
    }
  })

  t.Run("config: invalid yaml (duplicate keys)", func (t *testing.T) {
    _, e := Parse([]byte("test:\ntest:"))
    if e == nil {
      t.Fail()
    }
  })

  t.Run("config: invalid top-level key", func(t *testing.T) {
    _, e := Parse([]byte("extra:"))
    if e == nil {
      t.Fail()
    }
  })

  t.Run("plugins: default override", func(t *testing.T) {
    g, e := Parse(
      createConfig(
        []string{examples["plugins_default_override"]},
        []string{},
        []string{},
        []string{},
      ),
    )

    if e != nil || g.Plugins["environment"].Equal(default_generator.Plugins["environment"]) {
      t.Fail()
    }
  })

  t.Run("plugins: inject custom", func(t *testing.T) {
    g, _ := Parse(
      createConfig(
        []string{examples["plugins_custom_one"]},
        []string{},
        []string{},
        []string{},
      ),
    )

    if _, ok := g.Plugins["test1"]; !ok {
      t.Fail()
    }
  })

  t.Run("plugins: inject multiple", func(t *testing.T) {
    g, _ := Parse(
      createConfig(
        []string{examples["plugins_custom_multiple"]},
        []string{},
        []string{},
        []string{},
      ),
    )

    if _, ok := g.Plugins["test1"]; !ok {
      t.Fail()
    }

    if _, ok := g.Plugins["test2"]; !ok {
      t.Fail()
    }
  })
}
