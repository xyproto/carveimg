package main

import (
	_ "embed"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/esimov/caire"
	"github.com/xyproto/palgen"
	"github.com/xyproto/vt100"
)

const drawRune = '▒'

func convertToNRGBA(img image.Image) (*image.NRGBA, error) {
	nImage := image.NewNRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c, ok := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			if !ok {
				return nil, errors.New("could not convert color to NRGBA")
			}
			nImage.Set(x, y, c)
		}
	}
	return nImage, nil
}

func LoadImage(filename string) (*image.NRGBA, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// Read and decode the image
	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	if nImage, ok := img.(*image.NRGBA); ok {
		return nImage, nil
	}
	return convertToNRGBA(img)
}

func Draw(canvas *vt100.Canvas, m image.Image) error {
	// Convert the image to only use the basic 16-color palette
	img, err := palgen.ConvertBasic(m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// img is now an indexed image
	var vc vt100.AttributeColor
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			for i, rgb := range palgen.BasicPalette16 {
				if rgb[0] == c.R && rgb[1] == c.G && rgb[2] == c.B {
					switch i {
					case 0:
						vc = vt100.Black
					case 1:
						vc = vt100.Red
					case 2:
						vc = vt100.Green
					case 3:
						vc = vt100.Yellow
					case 4:
						vc = vt100.Blue
					case 5:
						vc = vt100.Magenta
					case 6:
						vc = vt100.Cyan
					case 7:
						vc = vt100.LightGray
					case 8:
						vc = vt100.DarkGray
					case 9:
						vc = vt100.LightRed
					case 10:
						vc = vt100.LightGreen
					case 11:
						vc = vt100.LightYellow
					case 12:
						vc = vt100.LightBlue
					case 13:
						vc = vt100.LightMagenta
					case 14:
						vc = vt100.LightCyan
					case 15:
						vc = vt100.White
					default:
						vc = vt100.White
					}
					break
				}
			}
			// Draw the "pixel" on the canvas using the vc color and the draw rune
			canvas.PlotColor(uint(x), uint(y), vc, drawRune)
		}
	}
	return nil
}

func main() {
	if len(os.Args) <= 1 {
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
	nImage, err := LoadImage(filename)
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
	if err := Draw(c, resizedImage); err != nil {
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
