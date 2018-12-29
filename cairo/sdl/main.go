package main

/*
#define CAIROSDL_ASHIFT 24
#define CAIROSDL_RSHIFT 16
#define CAIROSDL_GSHIFT  8
#define CAIROSDL_BSHIFT  0
#define CAIROSDL_AMASK (255U << CAIROSDL_ASHIFT)
#define CAIROSDL_RMASK (255U << CAIROSDL_RSHIFT)
#define CAIROSDL_GMASK (255U << CAIROSDL_GSHIFT)
#define CAIROSDL_BMASK (255U << CAIROSDL_BSHIFT)
*/
import "C"

import (
	"fmt"
	"image"
	"os"

	cairo "github.com/ungerik/go-cairo"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	// defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)
	rect := sdl.Rect{0, 0, 800, 600}
	surface.FillRect(&rect, 0xffffffff)
	window.UpdateSurface()

	for {
		event := sdl.WaitEvent()
		switch t := event.(type) {
		case *sdl.WindowEvent:
			switch t.Event {
			case sdl.WINDOWEVENT_EXPOSED:
				fmt.Println("WINDOWEVENT_EXPOSED ###")

				initCairo(surface)
				window.UpdateSurface()
			case sdl.WINDOWEVENT_CLOSE:
				fmt.Println("WINDOWEVENT_CLOSE")
			}
		case *sdl.QuitEvent:
			println("Quit")
			sdl.Quit()
			os.Exit(0)
			break
		}
	}
}

func initCairo(sdlSurface *sdl.Surface) {
	// surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 440, 80)
	surface := NewSurfaceFromSDL(sdlSurface)
	surface.SelectFontFace("Consolas", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surface.SetFontSize(15.0)
	surface.SetSourceRGB(0.2, 0.5, 0.1)
	// surface.SetSourceRGB(0.2*255, 0.5*255, 0.1*255)
	surface.MoveTo(10.0, 50.0)
	surface.ShowText("Hello World")
	// surface.WriteToPNG("hello.png")
	surface.Flush()
	// surface.Finish()
}

func NewSurfaceFromSDL(surface *sdl.Surface) *cairo.Surface {
	var target *cairo.Surface
	var format cairo.Format
	is_dirty := false

	/* Cairo only supports a limited number of pixels formats.  Make
	 * sure the surface format is compatible. */
	if surface.Format.BytesPerPixel != 4 ||
		surface.Format.BitsPerPixel != 32 {
		goto unsupported_format
	}

	if surface.Format.Rmask != C.CAIROSDL_RMASK ||
		surface.Format.Gmask != C.CAIROSDL_GMASK ||
		surface.Format.Bmask != C.CAIROSDL_BMASK {
		goto unsupported_format
	}

	switch surface.Format.Amask {
	case C.CAIROSDL_AMASK:
		format = cairo.FORMAT_ARGB32
	case 0:
		format = cairo.FORMAT_RGB24
	default:
		goto unsupported_format
	}

	/* Make the target point to either the SDL_Surface's data itself
	 * or a shadow image surface if we need to unpremultiply pixels. */
	if true || format == cairo.FORMAT_RGB24 {
		/* The caller is expected to have locked the surface (_if_ it
		 * needs locking) so that sdl_surface->pixels is valid and
		 * constant for the lifetime of the cairo_surface_t.  However,
		 * we're told not to call any OS functions when a surface is
		 * locked, so we really shouldn't call
		 * cairo_image_surface_create () as it will malloc, so really
		 * if the surface needs locking this shouldn't be used.
		 *
		 * However, it turns out malloc is actually safe on many (all?)
		 * platforms so we'll just go ahead anyway. */
		// unsigned char *data = (unsigned char*)(sdl_surface->pixels);
		data := surface.Data()
		target = cairo.NewSurfaceFromData(data, format, int(surface.W), int(surface.H), int(surface.Pitch))
		is_dirty = false
	} else {
		/* Need a shadow image surface. */
		rgba := &image.RGBA{
			Rect: image.Rectangle{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{X: int(surface.W), Y: int(surface.H)},
			},
		}
		target = cairo.NewSurfaceFromImage(rgba)
		is_dirty = true
	}

	fmt.Println(is_dirty)

	if target.Status() == cairo.STATUS_SUCCESS {
		surface.RefCount++
		// target.SetUserData(unsafe.Pointer(surface), surface. )
		// 	// }

		//     // if (cairo_surface_status (target) == CAIRO_STATUS_SUCCESS) {
		//         // sdl_surface->refcount++;
		//         cairo_surface_set_user_data (target,
		//                                      CAIROSDL_TARGET_KEY,
		//                                      sdl_surface,
		//                                      sdl_surface_destroy_func);

		//         if  is_dirty
		//             cairosdl_surface_mark_dirty (target);
	}

	return target

unsupported_format:
	/* Nasty kludge to get a cairo surface in CAIRO_INVALID_FORMAT
	 * state. */
	//  cairo.NewSurfaceFromImage(cairo.Format )
	// return cairo_image_surface_create(
	// 	(cairo_format_t)-1, 0, 0)
	return nil
}
