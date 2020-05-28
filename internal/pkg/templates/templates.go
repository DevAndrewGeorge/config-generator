package templates

import(
    "bytes"
    "text/template"
    "io/ioutil"
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

func (t *Template) RenderYaml(variables map[string]string) (string, error) {
    return "", nil
}

func (t *Template) RenderJson(variables map[string]string) (string, error) {
    return "", nil
}

func (t *Template) RenderMap(variables map[string]string) (map[string]interface{}, error) {
    return map[string]interface{}{}, nil
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
