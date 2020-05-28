package templates

import (
    "os"
    "encoding/json"
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
    t.Run("flat template", func(t *testing.T) {
        obj := &Template{}
        if _, err := obj.RenderJson(nil); err == nil { t.Fail() }
    })

    t.Run("empty", func(t *testing.T) {
        obj := &Template{templates: map[string]*Template{}}
        actual, err := obj.RenderJson(nil)
        if err != nil || actual != "{}" { t.Fail() }
    })

    t.Run("strings", func(t *testing.T) {
        text := "child"
        child := &Template{text: &text}
        obj := &Template{templates: map[string]*Template{"child": child}}

        encoded, err := obj.RenderJson(nil)
        if err != nil  { t.Error(err); return; }

        decoded  := map[string]string{}
        err = json.Unmarshal([]byte(encoded), &decoded)
        if err != nil { t.Error(err); return; }

        //making sure text is not wrapped
        if decoded["child"] != "child" { t.Error(decoded["child"]); return; }
    })

    t.Run("numbers", func(t *testing.T) {
        text := "1"
        child := &Template{text: &text}
        obj := &Template{templates: map[string]*Template{"child": child}}

        encoded, err := obj.RenderJson(nil)
        if err != nil { t.Fail() }

        decoded := &map[string]int{}
        err = json.Unmarshal([]byte(encoded), decoded)
        if err != nil { t.Error(err); return; }
    })

    t.Run("in-line objects", func(t *testing.T) {
        text := "{\"grandchild\": \"grandchild\"}"
        child := &Template{text: &text}
        obj := &Template{templates: map[string]*Template{"child": child}}

        encoded, err := obj.RenderJson(nil)
        if err != nil { t.Fail() }

        decoded := &map[string]map[string]string{}
        err = json.Unmarshal([]byte(encoded), decoded)
        if err != nil { t.Error(err); return; }
    })

    t.Run("nested objects", func(t *testing.T) {
        text := "grandchild"
        grandchild := &Template{text: &text}
        child := &Template{templates: map[string]*Template{"grandchild": grandchild}}
        obj := &Template{templates: map[string]*Template{"child": child}}

        encoded, err := obj.RenderJson(nil)
        if err != nil { t.Fail() }

        decoded := &map[string]map[string]string{}
        err = json.Unmarshal([]byte(encoded), decoded)
        if err != nil { t.Error(err); return; }
    })

    t.Run("arrays", func(t *testing.T) {
        text := "[1, 2, 3]"
        child := &Template{text: &text}
        obj := &Template{templates: map[string]*Template{"child": child}}

        encoded, err := obj.RenderJson(nil)
        if err != nil { t.Fail() }

        decoded := &map[string][]int{}
        err = json.Unmarshal([]byte(encoded), decoded)
        if err != nil { t.Error(err); return; }
    })

    t.Run("booleans", func(t *testing.T) {
        text := "true"
        child := &Template{text: &text}
        obj := &Template{templates: map[string]*Template{"child": child}}

        encoded, err := obj.RenderJson(nil)
        if err != nil { t.Fail() }

        decoded := &map[string]bool{}
        err = json.Unmarshal([]byte(encoded), decoded)
        if err != nil { t.Error(err); return; }
    })

    t.Run("explicit null", func(t *testing.T) {
        text := "null"
        child := &Template{text: &text}
        obj := &Template{templates: map[string]*Template{"child": child}}

        encoded, err := obj.RenderJson(nil)
        if err != nil { t.Fail() }

        decoded := &map[string]*string{}
        err = json.Unmarshal([]byte(encoded), decoded)
        if err != nil { t.Error(err); return; }
        if (*decoded)["child"] != nil { t.Fail(); return; }
    })
}

func TestTemplateRenderMap(t *testing.T) {
    t.Run("flat template", func(t *testing.T) {
        obj := &Template{}
        if _, err := obj.RenderMap(nil); err == nil { t.Fail() }
    })

    t.Run("no templates", func(t *testing.T) {
        obj := &Template{templates: map[string]*Template{}}
        actual, err := obj.RenderMap(nil)
        if err != nil || len(actual) > 0 { t.Fail() }
    })

    t.Run("one static template", func(t *testing.T) {
        text := "hello"
        child := &Template{text: &text}
        obj := &Template{templates: map[string]*Template{"child": child}}
        actual, err := obj.RenderMap(nil)
        if err != nil { t.Fail(); return; }
        if len(actual) != 1 { t.Fail(); return; }
        if value, found := actual["child"]; !found || value.(string) != "hello" { t.Fail(); return; }
    })

    t.Run("one dynamic template", func (t *testing.T) {
        text := "{{ .test }}"
        child := &Template{text: &text}
        obj := &Template{templates: map[string]*Template{"child": child}}

        actual, err := obj.RenderMap(map[string]string{"test": "test"})
        if err != nil { t.Fail(); return; }
        child_rendered_value, converted := actual["child"].(string)
        if !converted || child_rendered_value != "test" { t.Fail(); return; }
    })

    t.Run("invalid child template", func(t *testing.T) {
        text := "{{ .test "
        child := &Template{text: &text}
        obj := &Template{templates: map[string]*Template{"child": child}}

        if _, err := obj.RenderMap(map[string]string{"test": "test"}); err == nil { t.Fail() }
    })

    t.Run("multiple templates", func(t *testing.T) {
        hello_text, world_text := "hello", "world"
        child1, child2 := &Template{text: &hello_text}, &Template{text: &world_text}
        obj := &Template{templates: map[string]*Template{
            "hello": child1,
            "world": child2,
        }}

        actual, err := obj.RenderMap(nil)
        if err != nil { t.Fail(); return; }
        if len(actual) != 2 { t.Fail(); return; }
        if actual["hello"] != "hello" || actual["world"] != "world" { t.Fail() }
    })

    t.Run("nested template", func(t *testing.T) {
        text := "grandchild"
        grandchild := &Template{text: &text}
        child := &Template{templates: map[string]*Template{"grandchild": grandchild}}
        obj := &Template{templates: map[string]*Template{"child": child}}

        actual, err := obj.RenderMap(nil)

        if err != nil { t.Fail(); return; }
        if c, found := actual["child"]; found {
            if gc, found := c.(map[string]interface{})["grandchild"]; !found || gc.(string) != "grandchild" {
                t.Fail()
                return
            }
        } else {
            t.Fail();
            return;
        }

    })
}
