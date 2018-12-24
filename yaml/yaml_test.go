package main

import (
	"fmt"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

var template = `
import:
  - gitlab.com/mustodo/tenon/ui"
  - gitlab.com/mustodo/tenon/app"

Layout: {
  Widget: ui.Box,
  Width: 100%,
  Height: 200px,
  Align: ui.Center|ui.Left,
  Padding: [10, 10, 10, 10],
  Margin: {Top: 10, Bottom: 20},
  Children: [
    {
      Widget: ui.Box,
      Width: 100%,
      Height: 200px,
    },
  ],
}

`

type Artx struct {
	Import []string
}

func TestYaml(t *testing.T) {
	a := make(map[string]interface{}, 1)

	yaml.UnmarshalStrict([]byte(template), &a)

	for k, v := range a {
		fmt.Printf("%T ", v)
		fmt.Println(k, ":", v)

		if m, ok := v.(map[interface{}]interface{}); ok {
			for k1, v1 := range m {
				fmt.Printf("%T ", v1)
				fmt.Println(k1, ":", v1)
			}
		}
	}
}
