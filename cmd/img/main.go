package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/xyproto/carveimg"
	"github.com/xyproto/vt100"
	"golang.org/x/image/draw"
)

const versionString = "img 1.2.2"

// var imageResizeFunction = draw.NearestNeighbor
// var imageResizeFunction = draw.ApproxBiLinear
// var imageResizeFunction = draw.BiLinear
var imageResizeFunction = draw.CatmullRom

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(versionString)
		fmt.Fprintln(os.Stderr, "Please supply a GIF, JPEG or PNG image filename")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Initialize vt100 terminal settings
	vt100.Init()

	// Prepare a canvas
	c := vt100.NewCanvas()

	// Find the width and height of the canvas
	width := int(c.Width())
	height := int(c.Height())

	// Load the given filename
	nImage, err := carveimg.LoadImage(filename)
	if err != nil {
		vt100.Close()
		fmt.Fprintf(os.Stderr, "Could not load %s: %s\n", filename, err)
		os.Exit(1)
	}

	// Set the desired size to the size of the current terminal emulator
	resizedImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Resize the image
	imageResizeFunction.Scale(resizedImage, resizedImage.Rect, nImage, nImage.Bounds(), draw.Over, nil)

	// Draw the image to the canvas, using only the basic 16 colors
	if err := carveimg.Draw(c, resizedImage); err != nil {
		vt100.Close()
		fmt.Fprintln(os.Stderr, "Could not draw image: %s\n", err)
		os.Exit(1)
	}

	// Output the filename on top of the image
	baseFilename := filepath.Base(filename)
	c.Write(uint((width-len(baseFilename))/2), 0, vt100.LightBlue, vt100.BackgroundDefault, baseFilename)

	// Draw the contents of the canvas to the screen
	c.Draw()

	// Wait for a keypress
	vt100.WaitForKey()

	// Reset the vt100 terminal settings
	vt100.Close()
}
