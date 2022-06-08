package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/png"
	"os"

	"github.com/xyproto/vt100"
)

var (
	//go:embed grumpy-cat.png
	grumpyCat []byte
)

func Load() error {
	// Decode the image
	reader := bytes.NewReader(grumpyCat)
	img, err := png.Decode(reader)
	if err != nil {
		return err
	}
	fmt.Printf("img %T\n", img)
	return nil
}

func main() {
	if len(grumpyCat) == 0 {
		fmt.Fprintln(os.Stderr, "Embed error")
		os.Exit(1)
	}

	if err := Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Could not decode grumpy-cat.png: %s\n\n", err.Error())
		os.Exit(1)
	}

	// Initialize vt100 terminal settings
	vt100.Init()

	// Prepare a canvas
	c := vt100.NewCanvas()

	// Draw things on the canvas
	c.Plot(10, 10, '!')
	c.Write(12, 12, vt100.LightGreen, vt100.BackgroundDefault, "hi")
	c.Write(15, 15, vt100.White, vt100.BackgroundMagenta, "floating")
	c.PlotColor(12, 17, vt100.LightRed, '*')
	c.PlotColor(10, 20, vt100.LightBlue, 'ø')
	c.PlotColor(11, 20, vt100.LightBlue, 'l')

	c.WriteString(10, 21, vt100.White, vt100.BackgroundRed, "øl")

	// Draw the contents of the canvas
	c.Draw()

	// Wait for a keypress
	vt100.WaitForKey()

	// Reset the vt100 terminal settings
	vt100.Close()
}
