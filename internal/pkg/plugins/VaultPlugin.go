package plugins

type VaultPlugin struct {
  token string
  address string
  cacert_path string
  capath_path string
  verify bool
}

func NewVaultPlugin() Plugin {
  return &VaultPlugin{}
}

func (v VaultPlugin) Equal(o Plugin) bool {
  return Plugin(v) == o
}
