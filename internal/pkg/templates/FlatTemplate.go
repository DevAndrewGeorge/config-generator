package templates

type FlatTemplate struct{}

func (f FlatTemplate) Equal(o Template) bool {
  return o != nil && FlatTemplate(f) == o
}
