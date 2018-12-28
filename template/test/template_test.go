package test

import (
	"os"
	"testing"
	"text/template"
)

type Todo struct {
	Width   int
	Height  int
	Text    string
	Gravity string
}

func TestTemplate(t *testing.T) {
	td := Todo{
		Width:   100,
		Height:  100,
		Text:    "Hello Template",
		Gravity: "center",
	}

	tt, err := template.New("todos").Parse(Template())
	if err != nil {
		panic(err)
	}
	err = tt.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
}

func Template() string {
	return `
import "gitlab.com/tenon/widget"
import whizz "gitlab.com/whizz/widget"
Layout{
	Width: {{.Width}},
	Height: {{.Height}},
	Children: [
		Button{
			Width: {{.Width}},
			Height: {{.Height}},
			Text: {{.Text}},
			Color: #ff0000,
		},
		Label{
			Style: "home-label"
			Text: {{.Text}},
        },
        whizz.Home{
            
        }],
}
`
}
