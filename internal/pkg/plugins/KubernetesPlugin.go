package plugins

type KubernetesPlugin struct {
  kubeconfig_path string
}

func NewKubernetesPlugin() Plugin {
  return &KubernetesPlugin{}
}

func (k KubernetesPlugin) Equal(o Plugin) bool {
  return Plugin(k) == o
}
