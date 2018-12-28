package main

import (
	"runtime"

	"github.com/golang-ui/cairo"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
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

	window.MakeContextCurrent()

	// Display * x11_display = glfwGetX11Display();
	// GLXContext glx_context = glfwGetGLXContext(window);
	// Window x11_window = glfwGetX11Window(window);

	// cairo_device_t * cairo_device = cairo_glx_device_create(x11_display, glx_context);
	// cairo_surface_t * cairo_surface = cairo_gl_surface_create_for_window(cairo_device, x11_window, 800, 600);
	// cairo_device_destroy(cairo_device);
	// cairo_t * ctx = cairo_create(cairo_surface);

	context := window.GetNSGLContext()

	cairo.NewDevice

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
