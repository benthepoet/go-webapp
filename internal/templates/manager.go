package templates

import (
	"bytes"
	"errors"
	"html/template"
)

type TM struct {
	templates map[string]*template.Template
}

func NewTM(r string, m map[string][]string) *TM {
	t := &TM{make(map[string]*template.Template)}

	for k, v := range m {
		var fs []string
		for _, j := range v {
			fs = append(fs, r+"/"+j)
		}
		t.templates[k] = template.Must(template.ParseFiles(fs...))
	}

	return t
}

func (m *TM) RenderInLayout(f string, l string, c interface{}) (string, error) {
	tpl, found := m.templates[f]
	if !found {
		return "", errors.New("template was not found")
	}

	var buf bytes.Buffer

	err := tpl.ExecuteTemplate(&buf, l, c)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
