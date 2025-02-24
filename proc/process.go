package proc

import (
	"fmt"
	"image"
	"image/color"
	jpeg "image/jpeg"
	"os"
)

// RGB range 0-255 == 0-100% intensity uses 8-bit color
type RGBHist struct {
	R, G, B [256]uint32
	AlphaHist
	profile string
}

// Cmyk range 0-100 because it measures ink coverage
type CMYKHist struct {
	C, M, Y, K [101]uint32
	profile    string
}

// May delete but could find something interesting to use
// for opacity information
type AlphaHist struct {
	A [256]uint32
}

func Process(img string) error {
	fmt.Println("Processing test")
	reader, err := os.Open("proc/src/test.jpg")
	if err != nil {
		return err
	}

	writer, err := os.Create("proc/src/proc_test.jpg")
	if err != nil {
		return err
	}
	defer reader.Close()
	defer writer.Close()

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
	rH := populateRGBHist(alphaCol)
	// cH, cHkMax := populateCMYKHist(alphaCol)

	res, err := MakeHistogram(rH, uint32(bounds.Max.X*bounds.Max.Y))
	if err != nil {
		return err
	}

	err = jpeg.Encode(writer, res, nil)
	if err != nil {
		return err
	}

	return nil
}

func populateCMYKHist(data []color.Color) *CMYKHist {
	cmykH := new(CMYKHist)
	cmykH.profile = "SWOP"

	for _, col := range data {
		r, g, b, _ := col.RGBA()
		c, m, y, k := RGBtoCmyk(float32(r), float32(g), float32(b))
		cmykH.C[c]++
		cmykH.M[m]++
		cmykH.Y[y]++
		cmykH.K[k]++
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

func NormRGB(r, g, b float32) (rP, gP, bP float32) {
	total := r + g + b
	rP = (r / total)
	gP = (g / total)
	bP = (b / total)

	return rP, gP, bP
}

func RGBtoCmyk(r, g, b float32) (c, m, y, k uint32) {
	rP, gP, bP := NormRGB(r, g, b)
	tk := 1 - max(rP, gP, bP)
	tc := (1 - rP - tk) / (1 - tk)
	tm := (1 - gP - tk) / (1 - tk)
	ty := (1 - bP - tk) / (1 - tk)

	return uint32(tc * 100), uint32(tm * 100), uint32(ty * 100), uint32(tk * 100)
}
