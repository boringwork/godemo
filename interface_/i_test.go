package interface_

import (
	"fmt"
	"testing"
)

type Dimen interface {
	Mode() int32
	Value() float32
}

type dimen struct {
	mode  int32
	value float32
}

func (me dimen) Mode() int32 {
	return me.mode

}

func (me dimen) Value() float32 {
	return me.value
}

// func (me *dimen) Mode() int32 {
// 	return me.mode

// }

// func (me *dimen) Value() float32 {
// 	return me.value
// }

func p1(d Dimen) {
	fmt.Println(d)
}

func p2(d *Dimen) {
	fmt.Println(d)
}

func TestInterface(t *testing.T) {
	a := dimen{1, 2}
	fmt.Println(a.Mode())

	// p2(&a)
	p1(a)
	p1(&a)
}
