package plugins

import (
  "testing"
)

func TestNewKubernetesPlugin(t *testing.T) {
  k, ok := NewKubernetesPlugin().(*KubernetesPlugin)
  if !ok {
    t.Fail()
  }

  if k.kubeconfig_path != "~/.kube/config" {
    t.Fail()
  }
}

func TestKubernetesPluginConfigure(t *testing.T) {
  k := NewKubernetesPlugin().(*KubernetesPlugin)
  k.Configure("test", map[string]interface{}{"Kubeconfig": "testing"})
  if k.name != "test" || k.kubeconfig_path != "testing" {
    t.Fail()
  }
}
