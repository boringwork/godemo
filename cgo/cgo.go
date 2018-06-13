package main

import (
	"fmt"
	"unsafe"
)

type V struct {
    i int32
	j int64
}

func (this V) PutI() {
    fmt.Printf("i=%d\n", this.i)
}

func (this V) PutJ() {
    fmt.Printf("j=%d\n", this.j)
}

func main(){
	var v *V = new(V)
	fmt.Println("align i = ", unsafe.Alignof(v.i), "  align j =", unsafe.Alignof(v.j))
	// return
	fmt.Println("v=", v, "  v* = ", unsafe.Pointer(v), "   vp = ", uintptr( unsafe.Pointer(v) ))
	fmt.Println("vpp = ", unsafe.Pointer(uintptr( unsafe.Pointer(v) ))  )
	
    var i *int32 = (*int32)(unsafe.Pointer(v))
	*i = int32(98)
	fmt.Println("sizeof = " , uintptr( unsafe.Sizeof(int32(0))) )
    var j *int64 = (*int64)(      unsafe.Pointer(     uintptr( unsafe.Pointer(v) ) + uintptr( unsafe.Sizeof(0) )         )       )
    *j = int64(90)
    v.PutI()
	v.PutJ()
	
	var b []byte = []byte{55, 66, 77}
	var c *byte = &b[0]
	*c = 88
	fmt.Println("b=",b , " &b=", &b, "  &b[0]=", &b[0], "  unsafe.Pointer(&b)=", unsafe.Pointer(&b), "  unsafe.Pointer(&b[0])=", unsafe.Pointer(&b[0]))
	
    fmt.Println(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1))))
}