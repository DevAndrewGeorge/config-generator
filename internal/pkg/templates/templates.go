package templates

type Template interface {
  Equal(Template) bool
}
