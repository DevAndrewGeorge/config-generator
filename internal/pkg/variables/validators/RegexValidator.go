package validators

import(
    "errors"
    "regexp"
)

type RegexValidator struct {
    regexp *regexp.Regexp
}

func (r *RegexValidator) configure(config interface{}) error {
    var reg *regexp.Regexp
    var err error
    switch config.(type) {
    case nil:
        reg, err = regexp.Compile(".+")
        if err != nil { return err }
    case string:
        reg, err = regexp.Compile(config.(string))
        if err != nil { return err }
    default:
        return errors.New("invalid regex")
    }

    r.regexp = reg
    return nil
}

func (r *RegexValidator) Validate(test string) bool {
    return r.regexp.MatchString(test)
}

func (r1 *RegexValidator) Equal(o Validator) bool {
    r2, ok := o.(*RegexValidator)
    if !ok { return false }
    if r2 == nil { return false }

    if r1.regexp == r2.regexp { return true }
    if r1.regexp == nil || r2.regexp == nil { return false }
    if r1.regexp.String() != r2.regexp.String() { return false }

    return true
}
