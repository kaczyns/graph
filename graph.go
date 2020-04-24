package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// TODO: Figure out how to put origin off-center
// TODO: Abstract graph out so main can just make it and then call
//       the function.

// HLine draws a horizontal line
func HLine(img *image.RGBA, x1 int, y int, x2 int, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(img *image.RGBA, x int, y1 int, y2 int, col color.Color) {
    for ; y1 <= y2; y1++ {
        img.Set(x, y1, col)
    }
}

// Rect draws a rectangle utilizing HLine() and VLine()
/*
func Rect(img *image.RGBA, x1 int, y1 int, x2 int, y2 int, col color.Color) {
    HLine(img, x1, y1, x2, col)
    HLine(img, x1, y2, x2, col)
    VLine(img, x1, y1, y2, col)
    VLine(img, x2, y1, y2, col)
}
*/

var black = color.Gray16{0}
var grey = color.Gray16{0x8000}

const (
	scale = 100
	hmin = -5 * scale
	hmax = 5 * scale
	vmin = -5 * scale
	vmax = 5 * scale
)

func equation(x float32) float32 {
	return (0.5 * x) + 2.0
}

func main() {
	// Create an image and draw the graph on it.
	img := image.NewRGBA(image.Rect(hmin, vmin, hmax, vmax))

	// For the time being, we will multiply all values by 100 to get
	// the coordinates on the image (ie x=10 would be 1000 on the image).
	HLine(img, hmin, 0, hmax, black)
	VLine(img, 0, vmin, vmax, black)

	for x := 0; x <= hmax; x += scale {
		VLine(img, x, -10, 10, black)
	}
	
	for x := 0; x >= hmin; x -= scale {
		VLine(img, x, -10, 10, black)
	}

	for y	:= 0; y <= vmax; y += scale {
		HLine(img, -10, y, 10, black)
	}

	for y := 0; y >= vmin; y -= scale {
		HLine(img, -10, y, 10, black)
	}

	// TODO: Add numbers to the graph
	
	// OK now lets draw our line: y = 2x + 2
  red := color.RGBA{255, 0, 0, 255} // Red
	for x := hmin; x <= hmax; x++ {
		// The x values are the horizontal pixels of the image.
		// First we convert the value to floating point, and divide by the
		// scaling factor of the graph.  Then we compute y.  Finally we
		// multiply y by the scaling factor and convert back to an integer.
		xfloat := float32(x) / float32(scale)
		yfloat := equation(xfloat)
		y := int(yfloat * float32(scale))

		// Need to "flip" y because negative is at the bottom.
		y = y * -1
		
		// Don't plot the point unless y is within the rectangle.
		// We should be reasonably confident that x is within the rectangle
		if (y >= vmin) && (y <= vmax) {
			img.Set(x, y, red)
		}
	}
	
    //HLine(img, 10, 20, 80, grey)
    //col = color.RGBA{0, 255, 0, 255} // Green
    //Rect(img, 10, 10, 80, 50, black)

    f, err := os.Create("draw.png")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    png.Encode(f, img)
}
