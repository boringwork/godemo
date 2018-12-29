package reflect5

import "fmt"

type Button struct {
	Text   string
	Width  int32
	Height int32
}

func (me *Button) Mounted() {
	fmt.Println("Button mounted", me.Text, me.Width, me.Height)
}
