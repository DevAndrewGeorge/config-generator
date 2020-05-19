package plugins

import(
  "os"
  "reflect"
  "io/ioutil"
)

type VaultPlugin struct {
  name string
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

func (v *VaultPlugin) Equal(p Plugin) bool {
  a, ok := p.(*VaultPlugin)
  if ok {
    return *v == *a
  }

  return false
}

func (v *VaultPlugin) Configure(name string, settings map[string]interface{}) error {
  v.name = name

  if token, ok := settings["Token"]; ok && !reflect.ValueOf(token).IsNil() {
    v.token = *token.(*string)
  }

  if address, ok := settings["Address"]; ok && !reflect.ValueOf(address).IsNil() {
    v.address = *address.(*string)
  }

  if ca_cert, ok := settings["CAcert"]; ok && !reflect.ValueOf(ca_cert).IsNil() {
    v.cacert_path = *ca_cert.(*string)
  }

  if ca_path, ok := settings["CApath"]; ok && !reflect.ValueOf(ca_path).IsNil() {
    v.capath = *ca_path.(*string)
  }

  if skip_verify, ok := settings["SkipVerify"]; ok && !reflect.ValueOf(skip_verify).IsNil() {
    v.verify = *skip_verify.(*bool)
  }

  return nil
}
