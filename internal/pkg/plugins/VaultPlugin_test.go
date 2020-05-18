package plugins

import (
  "testing"
  "os"
)

func TestNewVaultPlugin(t *testing.T) {
  os.Setenv("VAULT_ADDR", "testing")
  os.Setenv("VAULT_CACERT", "testing")
  os.Setenv("VAULT_CAPATH", "testing")
  os.Setenv("VAULT_SKIP_VERIFY", "1")

  t.Run("type check", func (t *testing.T) {
    if _, ok := NewVaultPlugin().(*VaultPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("VAULT_TOKEN", func (t *testing.T) {
    os.Setenv("VAULT_TOKEN", "testing")
    v := NewVaultPlugin().(*VaultPlugin)
    if v.token != "testing" ||
       v.address != "testing" ||
       v.cacert_path != "testing" ||
       v.capath != "testing" ||
       v.verify != true {
      t.Fail()
    }
  })

  // TODO: find how to test without persisting changes
  t.Run("~/.vault-token", func(t *testing.T) {
    os.Unsetenv("VAULT_TOKEN")
    // v := NewVaultPlugin().(*VaultPlugin)
  })

  // TODO: run cleanly where I can be assured no file exist
  t.Run("no VAULT_TOKEN or ~/.vault-token", func (t *testing.T) {
    // v := NewVaultPlugin().(*VaultPlugin)
  })
}

func TestVaultPluginConfigure(t *testing.T) {
  v := NewVaultPlugin().(*VaultPlugin)
  v.Configure("test", map[string]interface{}{
    "token": "testing",
    "address": "testing",
    "ca_cert": "testing",
    "ca_path": "testing",
    "skip_verify": true,
  })

  if v.name != "test" ||
     v.token != "testing" ||
     v.address != "testing" ||
     v.cacert_path != "testing" ||
     v.capath != "testing" ||
     v.verify != true {
    t.Fail()
  }
}
