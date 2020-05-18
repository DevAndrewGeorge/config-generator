package plugins

import (
  "os"
  "os/user"
)

type ScriptPlugin struct {
  name string
  shell string
  user string
  group string
}

func NewScriptPlugin() Plugin {
  return &ScriptPlugin{
    shell: "/bin/bash",
    user: string(os.Getuid()),
    group: string(os.Getgid()),
  }
}

func (s *ScriptPlugin) Equal(p Plugin) bool {
  a, ok := p.(*ScriptPlugin)
  if ok {
    return *s == *a
  }

  return false
}

func (s *ScriptPlugin) Configure(name string, settings map[string]interface{}) error {
  s.name = name
  if shell, ok := settings["shell"]; ok {
    s.shell = shell.(string)
  }

  if u, ok := settings["user"]; ok {
    var script_user *user.User
    var err error

    user_string, _ := u.(string)
    if _, isInt := u.(int); isInt {
      script_user, err = user.LookupId(user_string)
    } else {
      script_user, err = user.Lookup(user_string)
    }

    if err != nil {
      return err
    }

    s.user = string(script_user.Uid)
  }

  if g, ok := settings["group"]; ok {
    var script_group *user.Group
    var err error

    group_string, _ := g.(string)

    if _, isInt := g.(int); isInt {
      script_group, err = user.LookupGroupId(group_string)
    } else {
      script_group, err = user.LookupGroup(group_string)
    }

    if err != nil {
      return err
    }

    s.group = string(script_group.Gid)
  }

  return nil
}
