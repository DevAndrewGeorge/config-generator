package plugins

import (
  "os"
  "os/user"
  "reflect"
  "strconv"
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
  if shell, ok := settings["Shell"]; ok {
    s.shell = shell.(string)
  }

  if u, ok := settings["User"]; ok && !reflect.ValueOf(u).IsNil() {
    var script_user *user.User
    var err error

    user_string := *u.(*string)
    if _, parse_error := strconv.ParseInt(user_string, 10, 32); parse_error == nil {
      script_user, err = user.LookupId(user_string)
    } else {
      script_user, err = user.Lookup(user_string)
    }

    if err != nil {
      return err
    }

    s.user = string(script_user.Uid)
  }

  if g, ok := settings["Group"]; ok {
    var script_group *user.Group
    var err error

    group_string := *g.(*string)
    if _, parse_error := strconv.ParseInt(group_string, 10, 32); parse_error == nil {
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
