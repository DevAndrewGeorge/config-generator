package templates

func New(name string, data interface{}) (*Template, error) {
    return &Template{}, nil
}

type Template struct {
    name string
    text *string
    templates map[string]*Template
}

func (t *Template) Render(variables map[string]string) (string, error) {
    return "", nil
}

func (t *Template) RenderYaml(variables map[string]string) (string, error) {
    return "", nil
}

func (t *Template) RenderJson(variables map[string]string) (string, error) {
    return "", nil
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
