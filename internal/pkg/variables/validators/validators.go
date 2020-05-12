package validators

type Validator struct {}

func (v Validator) Equal(o Validator) bool {
  return v == o
}
