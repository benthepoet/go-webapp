package templates

import (
	"errors"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/cbroglie/mustache"
)

type TM struct {
	templates map[string]*template.Template
}

type TemplateManager struct {
	templates map[string]*mustache.Template
}

func NewTM(r string, m map[string][]string) *TM {
	t := &TM{make(map[string]*template.Template)}

	for k, v := range m {
		t.templates[k] = template.Must(template.ParseFiles(v...))
	}

	return t
}

func New(p string, e string) *TemplateManager {
	files, err := os.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}

	man := &TemplateManager{map[string]*mustache.Template{}}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), e) {
			tpl, err := mustache.ParseFile(p + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			} else {
				i := strings.Index(file.Name(), e)
				man.templates[file.Name()[:i]] = tpl
			}
		}
	}

	return man
}

func (m *TemplateManager) Render(f string, c interface{}) (string, error) {
	tpl, found := m.templates[f]
	if !found {

		return "", errors.New("template was not found")
	}

	return tpl.Render(c)
}

func (m *TemplateManager) RenderInLayout(f string, l string, c interface{}) (string, error) {
	tpl, found := m.templates[f]
	if !found {
		return "", errors.New("template was not found")
	}

	lay, found := m.templates[l]
	if !found {
		return "", errors.New("layout was not found")
	}

	return tpl.RenderInLayout(lay, c)
}
