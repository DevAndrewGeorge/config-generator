package plugins

import(
  "os"
  "io/ioutil"
)

type VaultPlugin struct {
  token string
  address string
  cacert_path string
  capath string
  verify bool
}

func NewVaultPlugin() Plugin {
  v := VaultPlugin{}

  var e string
  var ok bool
  e, ok = os.LookupEnv("VAULT_TOKEN")
  if ok {
    v.token = e
  } else {
    token, err := ioutil.ReadFile("~/.vault-token")
    if err == nil {
      v.token = string(token)
    }
  }

  e, ok = os.LookupEnv("VAULT_ADDR")
  if ok {
    v.address = e
  }

  e, ok = os.LookupEnv("VAULT_CACERT")
  if ok {
    v.cacert_path = e
  }

  e, ok = os.LookupEnv("VAULT_CAPATH")
  if ok {
    v.capath = e
  }

  e, ok = os.LookupEnv("VAULT_SKIP_VERIFY")
  if ok && len(e) > 0 {
    v.verify = true
  }

  return &v
}

func (v *VaultPlugin) Equal(o Plugin) bool {
  return Plugin(v) == o
}

func (v *VaultPlugin) Configure(settings map[string]interface{}) error {
  if token, ok := settings["token"]; ok {
    v.token = token.(string)
  }

  if address, ok := settings["address"]; ok {
    v.address = address.(string)
  }

  if ca_cert, ok := settings["ca_cert"]; ok {
    v.cacert_path = ca_cert.(string)
  }

  if ca_path, ok := settings["ca_path"]; ok {
    v.capath = ca_path.(string)
  }

  if skip_verify, ok := settings["skip_verify"]; ok {
    v.verify = skip_verify.(bool)
  }

  return nil
}
