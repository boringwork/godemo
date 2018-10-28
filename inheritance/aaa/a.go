package aaa

type T int

func (me *T) m() {
	println("ok")
}

func F(i interface{ m() }) { i.m() }
