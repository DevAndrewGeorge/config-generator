package plugins

type KubernetesPlugin struct {
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

func (k *KubernetesPlugin) Configure(settings map[string]interface{}) error {
  if config_path, ok := settings["kubeconfig"]; ok {
    k.kubeconfig_path = config_path.(string)
  }

  return nil
}
