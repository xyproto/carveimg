package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/esimov/caire"
	"github.com/xyproto/carveimg"
	"github.com/xyproto/vt100"
)

const versionString = "carve 1.3.0"

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

	// Prepare the content-aware image resizing
	p := &caire.Processor{
		BlurRadius:     1, // or ie. 4
		SobelThreshold: 4, // or ie. 2
		NewWidth:       width,
		NewHeight:      height,
		Percentage:     false,
		Square:         false,
		Debug:          false,
		//Preview:        true, // Show GUI window, may cause errors related to context or ctx.Px
		FaceDetect: true,
		//FaceAngle:      0.0,
		//MaskPath:       "",
		//RMaskPath:      "",
		ShapeType: "circle",
		//SeamColor:      "#ff0000",
	}

	// Load the given filename
	nImage, err := carveimg.LoadImage(filename)
	if err != nil {
		vt100.Close()
		fmt.Fprintf(os.Stderr, "Could not load %s: %s\n", filename, err)
		os.Exit(1)
	}

	// Content aware resizing
	resizedImage, err := p.Resize(nImage)
	if err != nil {
		vt100.Close()
		fmt.Fprintf(os.Stderr, "Could not resize image: %s\n", err)
		os.Exit(1)

	}

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
