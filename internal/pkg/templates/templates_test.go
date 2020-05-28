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
        var i int
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

        actual, err := New("", map[string]interface{}{"file": temp.Name()})
        expected := &Template{text: &data}

        if err != nil || !actual.Equal(expected) { t.Error(expected, actual, err) }

        os.Remove(temp.Name())
    })

    t.Run("data is nested template", func(t *testing.T) {
        str := "test"
        data := map[string]interface{}{
            "keys": map[string]interface{}{ "test": str },
        }

        actual, err := New("", data)
        expected := &Template{
            templates: map[string]*Template{
                "test": &Template{name: "test", text: &str},
            },
        }

        if err != nil || !expected.Equal(actual) { t.Error(expected, actual, err) }
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

func TestTemplateIsNested(t *testing.T) {
    t.Run("untested", func(t *testing.T) {
        t1 := &Template{}
        if t1.IsNested() { t.Fail() }
    })

    t.Run("nested", func(t *testing.T) {
        t1 := &Template{
            templates: map[string]*Template{},
        }
        if !t1.IsNested() { t.Fail() }
    })
}

func TestTemplateRender(t *testing.T) {
    t.Run("variables is nil", func(t *testing.T) {
        text := ""
        obj := &Template{text: &text}
        actual, err := obj.Render(nil)
        if err != nil || actual != "" { t.Fail() }
    })

    t.Run("variable not found", func (t *testing.T) {
        text := "{{ .test }}"
        obj := &Template{text: &text}
        if _, err := obj.Render(nil); err == nil { t.Fail() }
    })

    t.Run("nested template", func(t *testing.T) {
        obj := &Template{
            templates: map[string]*Template{},
        }

        if _, err := obj.Render(nil); err == nil { t.Fail() }
    })

    t.Run("nil text", func(t *testing.T) {
        obj := &Template{}
        actual, err := obj.Render(nil)
        if err != nil || actual != "" { t.Fail() }
    })

    t.Run("empty text", func(t *testing.T) {
        text := ""
        obj := &Template{text: &text}
        actual, err := obj.Render(nil)
        if err != nil || actual != "" { t.Fail() }
    })

    t.Run("static text", func(t *testing.T) {
        text := "test"
        obj := &Template{text: &text}
        actual, err := obj.Render(nil)
        if err != nil || actual != "test" { t.Fail() }
    })

    t.Run("valid template text", func(t *testing.T) {
        text := "{{ .test }}"
        obj := &Template{text: &text}

        actual, err := obj.Render(map[string]string{"test": "test"})
        if err != nil || actual != "test" { t.Fail() }
    })

    t.Run("invalid template text", func(t *testing.T) {
        text := "{{ .test "
        obj := &Template{text: &text}
        if _, err := obj.Render(nil); err == nil { t.Fail() }
    })
}

func TestTemplateRenderYaml(t *testing.T) {

}

func TestTemplateRenderJson(t *testing.T) {

}

func TestTemplateRenderMap(t *testing.T) {

}
