package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

const Width = 1280
const Height = 720
const ComponentsPerPixel = 4

// setPixel is used to assign a pixel specific red, green, and blue components.
// Each pixel is assumed to use the RGBA uint8 color format. A color can be defined
// by specifying a value between 0 and 255 (inclusive) for each color component.
func setPixel(x int, y int, red uint8, green uint8, blue uint8, pixels *[]uint8) {
	firstColorIndex := (x * ComponentsPerPixel) + (Width * y * ComponentsPerPixel)

	(*pixels)[firstColorIndex] = red
	(*pixels)[firstColorIndex+1] = green
	(*pixels)[firstColorIndex+2] = blue
	(*pixels)[firstColorIndex+3] = 255
}

func main() {
	// Create a blank image
	data := image.NewRGBA(image.Rect(0, 0, Width, Height))

	// Set all pixels to baby blue
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			setPixel(x, y, 0, 204, 255, &data.Pix)
		}
	}

	// Create a new file
	file, error := os.Create("output.png")
	if error != nil {
		fmt.Println("Error while creating output image: ", error)
	}

	// Write the image data as a .png file
	png.Encode(file, data)

	// Close the file handle
	file.Close()
}
