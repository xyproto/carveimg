package main

import (
	"fmt"
	"os"

	"github.com/esimov/caire"
	"github.com/xyproto/preview"
	"github.com/xyproto/vt100"
)

const versionString = "preview 1.0.2"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(versionString)
		fmt.Fprintln(os.Stderr, "Please supply a PNG filename")
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
	nImage, err := preview.LoadImage(filename)
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
	if err := preview.Draw(c, resizedImage); err != nil {
		vt100.Close()
		fmt.Fprintln(os.Stderr, "Could not draw image: %s\n", err)
		os.Exit(1)
	}

	// Output the filename on top of the image
	//c.Write(10, 10, vt100.LightBlue, vt100.BackgroundDefault, filename)

	// Draw the contents of the canvas to the screen
	c.Draw()

	// Wait for a keypress
	vt100.WaitForKey()

	// Reset the vt100 terminal settings
	vt100.Close()

}
