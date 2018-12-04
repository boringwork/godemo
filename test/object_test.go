package test

import (
	"fmt"
	"testing"
)

func TestObject(t *testing.T) {
	fmt.Println("hello")
	t.Fail()
}
