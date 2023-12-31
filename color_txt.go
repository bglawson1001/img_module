package img_module

import (
	"fmt"
	_ "image/jpeg"

	"os"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/gofont/goregular"
)

func Text(txt, outputDestination string) error {
	const W = 500
	const H = 300

	// Create a temporary file and write the byte slice to it
	tempFile, err := os.CreateTemp("", "font-*.ttf")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(goregular.TTF); err != nil {
		panic(err)
	}

	dc := gg.NewContext(W, H)

	if err := dc.LoadFontFace(tempFile.Name(), 72); err != nil {
		panic(err)
	}

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(.5, 0, 0)
	dc.DrawStringAnchored(txt, W/2, H/2, 0.5, 0.5)
	dc.Stroke()

	dc.SavePNG(outputDestination)

	fmt.Println("Image with added text saved.")

	return nil
}
