package mask

import (
	"fmt"
	"testing"
	"unsafe"
)

type _TestDimen int32

type MeasuredDimen int32

const MaxMeasuredDimen int32 = 0x3FFFFFFF
const measuredDimenModeMask int32 = -1073741824 //0xC0000000
const measuredDimenValueMask int32 = 0x3FFFFFFF

func AMeasuredDimen(value int32, mode int32) MeasuredDimen {
	if mode < 0 || mode > 2 {
		panic("error for MeasuredDimen mdoe")
	}
	if value < 0 || value > MaxMeasuredDimen {
		panic("error for MeasuredDimen value")
	}
	return MeasuredDimen(mode<<30 | value)
}

func (me *MeasuredDimen) Mode() int32 {
	return (int32(*me) & measuredDimenModeMask) >> 30 & 3
}

func (me *MeasuredDimen) Value() int32 {
	return int32(*me) & measuredDimenValueMask
}

func TestDimen(t *testing.T) {
	var v uint32 = 0xC0000000
	v1 := *(*int32)(unsafe.Pointer(&v))
	fmt.Println(v1)

	m := measuredDimenModeMask
	fmt.Printf("2 = %.32b， %.8X\n", -3<<30, -3<<30)
	fmt.Printf("m = %.32b， %.8X\n", uint32(m), uint32(m))
	// mode := 2
	// value := 10
	// fmt.Printf("2 = %.32b， %.8X\n", mode<<30|value, mode<<30|value)
	fmt.Printf("2 = %.32b， %.8X\n", 0xC0000000, 0xC0000000)
	fmt.Printf("2 = %.32b， %.8X\n", 0x3FFFFFFF, 0x3FFFFFFF)
	fmt.Printf("2 = %.32b， %.32b\n", AMeasuredDimen(2, 2), uint32(AMeasuredDimen(2, 2)))
	a := AMeasuredDimen(MaxMeasuredDimen, 0)
	fmt.Printf("value = %.32b   mode = %.32b\n", a.Value(), a.Mode())
	mode := a.Mode()
	fmt.Printf("value = %.32b   mode = %.32b\n", a.Value(), *(*uint32)(unsafe.Pointer(&mode)))

}

func TestDimen2(t *testing.T) {
	var a int32 = -99
	var b byte = 2
	var d int32 = -10 >> 3
	c := a | int32(b)
	e := a & d
	fmt.Printf("2 = %.32b， %.32b\n", a, uint32(a))
	fmt.Printf("2 = %.32b， %.32b\n", b, uint32(b))
	fmt.Printf("2 = %.32b， %.32b\n", c, uint32(c))
	fmt.Printf("2 = %.32b， %.32b\n", d, uint32(d))
	fmt.Printf("2 = %.32b， %.32b\n", e, uint32(e))
	var f uint32 = 0xC0000000
	g := int32(f)
	fmt.Printf("2 = %d， %.32b\n", g, uint32(g))
}
