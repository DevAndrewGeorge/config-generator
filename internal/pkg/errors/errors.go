package errors

import(

)

type ConfigError struct {
   Message string
}

func (c *ConfigError) Error() string {
  return c. Message
}

type OutputError struct {
   Message string
}

func (o *OutputError) Error() string {
  return o. Message
}

type PluginError struct {
   Message string
}

func (p *PluginError) Error() string {
  return p. Message
}

type TemplateError struct {
   Message string
}

func (t *TemplateError) Error() string {
  return t. Message
}

type VariableError struct {
   Message string
}

func (v *VariableError) Error() string {
  return v. Message
}
