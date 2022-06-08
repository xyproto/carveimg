package main

import (
	"bytes"
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

var (
	//go:embed grumpy-cat.png
	grumpyCat []byte
)

func convertToNRGBA(img image.Image) (*image.NRGBA, error) {
	nImage := image.NewNRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c, ok := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			if !ok {
				return nil, errors.New("could not convert color to NRGBA")
			}
			nImage.Set(x, y, c)
			// 			r := int(c.R)
			// 			g := int(c.G)
			// 			b := int(c.B)
			// 			alpha := int(c.A) // TODO: assume a black background and multiply in the alpha
			// 			fmt.Println(r, g, b, alpha)
		}
	}
	return nImage, nil
}

func LoadEmbeddedImage() (*image.NRGBA, error) {
	if len(grumpyCat) == 0 {
		return nil, errors.New("embed error: no data")
	}
	// Decode the image
	reader := bytes.NewReader(grumpyCat)
	img, err := png.Decode(reader)
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
	fmt.Printf("img after converting to 16-colors: %T\n", img)

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c, ok := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			if !ok {
				errors.New("could not convert color to NRGBA")
			}
			_ = c
			// Draw things on the canvas
			canvas.PlotColor(uint(x), uint(y), vt100.LightRed, '.') // TODO: Use the correct vt100 color based on perhaps the img.At(x, y) color index?
		}
	}
	return nil
}

func main() {
	// Initialize vt100 terminal settings
	vt100.Init()

	// Prepare a canvas
	c := vt100.NewCanvas()

	// Find the width and height of the canvas
	width := int(c.Width())
	height := int(c.Height())

	p := &caire.Processor{
		BlurRadius:     4, // or ie. 1
		SobelThreshold: 2, // or ie. 4
		NewWidth:       width,
		NewHeight:      height,
		Percentage:     false,
		Square:         false,
		Debug:          false,
		//Preview:        true, // Show GUI window, may cause errors related to context or ctx.Px
		//FaceDetect:     false,
		//FaceAngle:      0.0,
		//MaskPath:       "",
		//RMaskPath:      "",
		//ShapeType:      "circle",
		//SeamColor:      "#ff0000",
	}

	nImage, err := LoadEmbeddedImage()
	if err != nil {
		vt100.Close()
		fmt.Fprintf(os.Stderr, "Could not decode grumpy-cat.png: %s\n", err)
		os.Exit(1)
	}

	// 	var buf bytes.Buffer
	// 	reader := bytes.NewReader(grumpyCat)
	// 	err := proc.Process(&buf, reader)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Could not process grumpy-cat.png: %s\n", err.Error())
	// 		os.Exit(1)
	// 	}
	// 	p.Process(in, out)

	fmt.Printf("IMAGE IS %T\n", nImage)

	resizedImage, err := p.Resize(nImage)
	if err != nil {
		vt100.Close()
		fmt.Fprintf(os.Stderr, "Could not resize grumpy-cat.png: %s\n", err)
		os.Exit(1)

	}

	fmt.Printf("RESIZED: %T\n", resizedImage)

	if err := Draw(c, resizedImage); err != nil {
		vt100.Close()
		fmt.Fprintln(os.Stderr, "Could not draw the image")
		os.Exit(1)
	}

	c.Write(12, 12, vt100.LightGreen, vt100.BackgroundDefault, "grumpy-cat.png")

	// 	for y := 0; y < height; y++ {
	// 		for x := 0; x < width; x++ {
	// 		}
	// 	}

	// 	nImage := image.NewNRGBA(img.Bounds())
	// 	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
	// 		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
	// 			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
	// 			r := int(c.R)
	// 			g := int(c.G)
	// 			b := int(c.B)
	// 			alpha := int(c.A) // TODO: assume a black background and multiply in the alpha
	// 			fmt.Println(r, g, b, alpha)
	// 		}
	// 	}

	// Draw the contents of the canvas
	c.Draw()

	// Wait for a keypress
	vt100.WaitForKey()

	// Reset the vt100 terminal settings
	vt100.Close()

}
