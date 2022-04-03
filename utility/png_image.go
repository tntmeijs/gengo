package utility

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
)

// Contains all functionality needed to export raw RGBA pixel data as a .png file
type PngImage struct {
	data     *image.RGBA
	fileName string
}

// Create a new output image
func NewPngImage(width int, height int, fileName string) PngImage {
	return PngImage{image.NewRGBA(image.Rect(0, 0, width, height)), fileName}
}

// Assign the specified pixel red, green, and blue components.
// Each pixel is assumed to use the RGBA uint8 color format. A color can be defined
// by specifying a value between 0 and 255 (inclusive) for each color component.
func (p *PngImage) SetPixelColor(x int, y int, color Color) {
	index := p.data.PixOffset(x, y)
	p.data.Pix[index] = color.Red
	p.data.Pix[index+1] = color.Green
	p.data.Pix[index+2] = color.Blue
	p.data.Pix[index+3] = color.Alpha
}

// Write the PNG image to a file on disk
func (p *PngImage) WritePngToFile() error {
	file, error := os.Create(p.fileName)

	// Unable to create a new file on disk
	if error != nil {
		return errors.New(fmt.Sprintf("Unable to write %s to disk: %s", p.fileName, error.Error()))
	}

	png.Encode(file, p.data)
	error = file.Close()

	// File handle failed to close
	if error != nil {
		return errors.New(fmt.Sprintf("Unable to close file handle: %s", error.Error()))
	}

	// Succes
	return nil
}
