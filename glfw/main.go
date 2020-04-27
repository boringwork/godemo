package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func CursorPosCallback(w *glfw.Window, xpos float64, ypos float64) {
	fmt.Println(xpos, ypos)
}

func MouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	fmt.Println(button, action, mod)
}

func ScrollCallback(w *glfw.Window, xoff float64, yoff float64) {
	fmt.Println(xoff, yoff)
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.SetCursorPosCallback(CursorPosCallback)
	window.SetMouseButtonCallback(MouseButtonCallback)
	window.SetScrollCallback(ScrollCallback)

	window.SetInputMode(glfw.StickyKeysMode, 1)

	// window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		// glfw.PollEvents()
		glfw.WaitEvents()
	}
}
