package templates

func New(name string, data interface{}) (*Template, error) {
    return nil, nil
}

type Template struct {
    name string
    text *string
    templates map[string]*template
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
    if t.text != o.text { return false }
    if length(t.templates) != length(o.templates) { return false }
    for k,v := range t.templates {
        if v != o.templates[k] { return false }
    }
    return true
}
