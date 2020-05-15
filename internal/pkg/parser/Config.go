package parser

type GeneratorConfig struct {
  Plugins map[string]map[string]interface{}
  Variables map[string]map[string]interface{}
  Templates map[string]interface{}
  Outputs map[string]map[string]interface{}
}
