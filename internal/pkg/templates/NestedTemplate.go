package templates

type NestedTemplate struct {}

func (n NestedTemplate) Equal(o Template) bool {
  return o != nil && Template(n) == o
}
