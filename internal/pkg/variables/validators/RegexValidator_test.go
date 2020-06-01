package validators

import(
    "testing"
)

func TestRegexEqual(t *testing.T) {
    t.Run("other is generic nil (false)", func(t *testing.T) {
        r := &RegexValidator{}
        if r.Equal(nil) { t.Fail() }
    })

    t.Run("other is same type, nil (false)", func(t *testing.T) {
        r := &RegexValidator{}
        var o_ptr *RegexValidator = nil
        if r.Equal(o_ptr) { t.Fail() }
    })

    t.Run("other is same type, unequal (false)", func(t *testing.T) {
        r1, r2 := &RegexValidator{}, &RegexValidator{}
        r1.configure("test")
        if r1.Equal(r2) { t.Fail() }
    })

    t.Run("other is same type, equal (true)", func(t *testing.T) {
        r1, r2 := &RegexValidator{}, &RegexValidator{}
        r1.configure("test")
        r2.configure("test")
        if !r1.Equal(r2) { t.Fail() }
    })


    t.Run("other is different validator type (false)", func(t *testing.T) {
        t.Error("no other type exists yet")
    })
}

func TestRegexConfigure(t *testing.T) {
    t.Run("nil config type (valid)", func (t *testing.T) {
        r := &RegexValidator{}
        err := r.configure(nil)
        if err != nil { t.Error(err) }
    })

    t.Run("string config type (valid)", func (t *testing.T) {
        r := &RegexValidator{}
        err := r.configure("test")
        if err != nil { t.Error(err, r) }
    })

    t.Run("invalid config type (invalid)", func (t *testing.T) {
        r := &RegexValidator{}
        if err := r.configure(2); err == nil { t.Error(r) }
    })
}

func TestRegexValidate(t *testing.T) {
    t.Run("custom regex: matching", func (t *testing.T) {
        r := &RegexValidator{}
        r.configure("test")
        if !r.Validate("testing") { t.Fail() }
    })

    t.Run("custom regex: unmatching", func (t *testing.T) {
        r := &RegexValidator{}
        r.configure("tset")
        if r.Validate("testing") { t.Fail() }
    })
}
