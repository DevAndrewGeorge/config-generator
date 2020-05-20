package templates

import (
    "os"
    "io/ioutil"
    "testing"
)

func TestNew(t *testing.T) {
    t.Run("data is nil", func(t *testing.T) {
        actual, err := New("test", nil)
        expected := &Template{name: "test"}
        if err != nil || !actual.Equal(expected) { t.Fail() }
    })

    t.Run("data is incorrect type", func(t *testing.T) {
        var i interface{}
        if _, err := New("test", i); err == nil { t.Fail() }
    })

    t.Run("data is string", func(t *testing.T) {
        data := "test"
        actual, err := New("", data)
        expected := &Template{text: &data}
        if err != nil || !actual.Equal(expected) { t.Fail() }
    })

    t.Run("data is file", func(t *testing.T) {
        filename := "testfile"
        data := "test"

        temp, _ := ioutil.TempFile(".", filename)
        temp.Write([]byte(data))
        temp.Close()

        actual, err := New("", map[string]interface{}{"file": filename})
        expected := &Template{text: &data}

        if err != nil || !actual.Equal(expected) { t.Fail() }

        os.Remove(filename)
    })

    t.Run("data is nested template", func(t *testing.T) {
        str := "test"
        data := map[string]interface{}{
            "keys": map[string]interface{}{ "test": &str },
        }

        actual, err := New("", data)
        expected := &Template{
            templates: map[string]*Template{
                "test": &Template{text: &str},
            },
        }

        if err != nil || !actual.Equal(expected) { t.Fail() }
    })
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
