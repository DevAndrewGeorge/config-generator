package plugins

import(
  "reflect"
)

type KubernetesPlugin struct {
  name string
  kubeconfig_path string
}

func NewKubernetesPlugin() Plugin {
  return &KubernetesPlugin{
    kubeconfig_path: "~/.kube/config",
  }
}

func (k *KubernetesPlugin) Equal(p Plugin) bool {
  a, ok := p.(*KubernetesPlugin)
  if ok {
    return *k == *a
  }

  return false
}

func (k *KubernetesPlugin) Configure(name string, settings map[string]interface{}) error {
  k.name = name
  if config_path, ok := settings["Kubeconfig"]; ok && !reflect.ValueOf(config_path).IsNil() {
      k.kubeconfig_path = *config_path.(*string)
  }

  return nil
}
