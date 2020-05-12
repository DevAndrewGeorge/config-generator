package plugins

type VaultPlugin struct {}

func NewVaultPlugin() Plugin {
  return &VaultPlugin{}
}
