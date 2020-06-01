package validators

func New(validator_type string, config interface{}) (Validator, error) {
    var v Validator
    switch validator_type {
    case "regex":
    default:

    }

    return v, nil
}

type Validator interface {
    Equal(o Validator) bool
    Validate(test string) bool
    configure(config interface{}) error
}
