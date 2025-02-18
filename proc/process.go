package proc

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"math"
	"os"
)

// RGB range 0-255 == 0-100% intensity uses 8-bit color
type RGBHist struct {
	R, G, B [256]int
	AlphaHist
	profile string
}

// Cmyk range 0-100 because it measures ink coverage but
// Not sure if ill have 256 vals or 100 vals, when I convert
// the values.
type CMYKHist struct {
	C, M, Y, K [256]int
	profile    string
}

// May delete but could find something interesting to use
// for opacity information
type AlphaHist struct {
	A [256]int
}

func Process(img string) error {
	fmt.Println("Processing test")
	reader, err := os.Open("proc/src/test.jpg")
	if err != nil {
		return err
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	bounds := m.Bounds()
	alphaCol := []color.Color{}
	// TODO: Keep eye open for potential optimization.
	// Doubt theres one for what im doing. Since I want
	// to see every pixel.
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			alphaCol = append(alphaCol, m.At(x, y))
		}
	}

	//TODO: Open two go routines or build out seperate function
	rgbHist := populateRGBHist(alphaCol)
	cmykHist := populateCMYKHist(alphaCol)

	return nil
}

func populateCMYKHist(data []color.Color) *CMYKHist {
	cmykH := new(CMYKHist)
	cmykH.profile = "SWOP"

	for _, col := range data {
		// normalize r' = r\total
		r, g, b, _ := col.RGBA()
		total := r + g + b
		sTotal := math.Sqrt(float64(r ^ 2 + g ^ 2 + b ^ 2))
		fmt.Println(total, sTotal)
		//get cmy
		// get k
		// adjust
	}
	return cmykH
}

func populateRGBHist(data []color.Color) *RGBHist {
	rgbH := new(RGBHist)
	rgbH.profile = "sRBG--Swop"
	for _, col := range data {
		r, g, b, a := col.RGBA()
		rgbH.R[r>>8]++
		rgbH.G[g>>8]++
		rgbH.B[b>>8]++
		rgbH.A[a>>8]++
	}
	return rgbH
}

func NormRGB() {}
