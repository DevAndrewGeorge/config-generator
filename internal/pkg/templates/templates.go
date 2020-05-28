package templates

import(
    "fmt"
    "bytes"
    "text/template"
    "io/ioutil"
    "encoding/json"
    "github.com/devandrewgeorge/config-generator/internal/pkg/errors"
)

func New(name string, data interface{}) (*Template, error) {
    template := &Template{name: name}
    switch data.(type) {
    case nil:
    case string:
        converted := data.(string)
        template.text = &converted
    case map[string]interface{}:
        converted := data.(map[string]interface{})
        if filename, found := converted["file"]; found {
            content, err := ioutil.ReadFile(filename.(string))
            if err != nil { return nil, err }
            temp := string(content)
            template.text = &temp
            break
        } else if _, found := converted["keys"]; found {
            template.templates = map[string]*Template{}
            for k, v := range converted["keys"].(map[string]interface{}) {
                child, err := New(k, v)
                if err != nil { return nil, err }
                template.templates[k] = child
            }
        } else {
            return nil, &errors.TemplateError{}
        }


    default:
        return nil, &errors.TemplateError{}
    }


    return template, nil
}

type parse_func func(string) (interface{}, error)
type Template struct {
    name string
    text *string
    templates map[string]*Template
}

func (t *Template) IsNested() bool {
    return t.templates != nil
}

func (t *Template) Render(variables map[string]string) (string, error) {
    if t.IsNested() {
        return "", &errors.TemplateError{}
    }

    if t.text == nil {
        return "", nil
    }

    renderer := template.New(t.name)
    renderer.Option("missingkey=error")
    if _, err := renderer.Parse(*t.text); err != nil {
        return "", err
    }

    result := &bytes.Buffer{}
    if variables == nil {
        variables = map[string]string{}
    }
    if err := renderer.Execute(result, variables); err != nil {
        return "", err
    }

    return result.String(), nil
}

func (t *Template) parse_json(raw string) (interface{}, error) {
    var i interface{}
    err := json.Unmarshal([]byte(raw), &i)
    fmt.Println(i)

    if err != nil {
        str := new(string)
        err = json.Unmarshal([]byte(fmt.Sprintf("\"%s\"", raw)), str)

        if err != nil { return nil, err }
        return str, nil
    }
    return i, nil
}

func (t *Template) parse_yaml(raw string) (interface{}, error) {
    return nil, nil
}

func (t *Template) parse(rendered map[string]interface{}, parser parse_func) (map[string]interface{}, error) {
    var err error
    parsed := map[string]interface{}{}
    for key, child := range rendered {
        raw, isChild := child.(string)
        if isChild {
            parsed[key], err = parser(raw)
            if err != nil { return nil, err }
        } else {
            parsed[key], err = t.parse(child.(map[string]interface{}), parser)
            if err != nil { return nil, err }
        }
    }

    return parsed, nil
}

func (t *Template) RenderYaml(variables map[string]string) (string, error) {
    return "", nil
}

func (t *Template) RenderJson(variables map[string]string) (string, error) {
    if !t.IsNested() { return "", &errors.TemplateError{} }

    rendered, render_error := t.RenderMap(variables)
    if render_error != nil { return "", render_error }

    parsed, parse_error := t.parse(rendered, t.parse_json)
    if parse_error != nil { return "", parse_error }

    encoded, encoding_error := json.MarshalIndent(parsed, "", "    ")
    if encoding_error != nil { return "", encoding_error }
    return string(encoded), nil
}

func (t *Template) RenderMap(variables map[string]string) (map[string]interface{}, error) {
    if !t.IsNested() { return nil, &errors.TemplateError{} }
    if len(t.templates) == 0 { return map[string]interface{}{}, nil }

    rendered := map[string]interface{}{}
    for key, child := range t.templates {
        var err error
        if child.IsNested() {
            rendered[key], err = child.RenderMap(variables)
            if err != nil { return nil, err }
        } else {
            rendered[key], err = child.Render(variables)
            if err != nil { return nil, err }
        }
    }

    return rendered, nil
}

func (t *Template) Equal(o *Template) bool {
    if t.name != o.name { return false }

    if t.text != nil || o.text != nil {
        if t.text == nil || o.text == nil {
            return false
        } else if *t.text != *o.text {
            return false
        }
    }

    if len(t.templates) != len(o.templates) { return false }
    for k, t_child := range t.templates {
        o_child, found := o.templates[k]
        if !found || !t_child.Equal(o_child) { return false }
    }
    return true
}
