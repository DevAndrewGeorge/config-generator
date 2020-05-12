package plugins

type KubernetesPlugin struct {}

func NewKubernetesPlugin() Plugin {
  return &KubernetesPlugin{}
}
