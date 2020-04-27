package main

type Canvas interface {
	DrawColor(color uint32)
}

type canvas struct {
	ptr interface{}
}

func NewCanvas() Canvas {
	return &canvas{}
}

func (me *canvas) DrawColor(color uint32) {

}
