package templates

import (
    "testing"
)

func TestNew(t *testing.T) {

}

func TestTemplateEqual(t *testing.T) {
    t.Run("all things equal", func(t *testing.T) {
        text1, text2 := "test", "test"
        templates1 := map[string]*Template{"child": &Template{name: "child"}}
        templates2 := map[string]*Template{"child": &Template{name: "child"}}
        t1 := &Template{
            name: "test",
            text: &text1,
            templates: templates1,
        }

        t2 := &Template{
            name: "test",
            text: &text2,
            templates: templates2,
        }

        if !t1.Equal(t2) { t.Fail() }
    })

    t.Run("different names", func(t *testing.T) {
        t1, t2 := &Template{}, &Template{name: "test"}
        if t1.Equal(t2) { t.Fail() }
    })

    t.Run("different text", func(t *testing.T) {
        text := "test"
        t1, t2 := &Template{}, &Template{text: &text}
        if t1.Equal(t2) { t.Fail() }
    })

    t.Run("different templates", func(t *testing.T) {
        child1, child2 := &Template{name: "child1"}, &Template{name: "child2"}
        t1 := &Template{templates: map[string]*Template{"child": child1}}
        t2 := &Template{templates: map[string]*Template{"child": child2}}

        if t1.Equal(t2) { t.Fail() }
    })
}
