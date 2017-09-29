package main

import (
	"github.com/master-pfa-info/imgsrv"
)

func main() {
	img := NewImage(256, 256)
	// call 'fill' on img and f1
	// ...

	// then, send the image's content to be displayed
	imgsrv.Paint(img.title, img.pixels)
	// and, finally, the path to a file where to store the image.
	imgsrv.Print("./img.png")
}

func f1(x, y, int) uint8 { return uint8(x+y) / 2 }
