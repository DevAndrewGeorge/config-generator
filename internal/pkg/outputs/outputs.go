package outputs

type Output struct {
}

func (a Output) Equal(b Output) bool {
  return a == b
}
