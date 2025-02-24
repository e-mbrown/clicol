package proc

import (
	"image"
	"image/color"
	"image/draw"
)

var greyRGB = color.RGBA{37, 37, 37, 0}

// var greyCMYK = color.CMYK{0, 0, 0, 0}

type BChart struct {
	canvas image.Rectangle
	it     int
}

func MakeHistogram(data *RGBHist, tPx uint32) (image.Image, error) {
	c := makeDefaultBChart()
	img := image.NewRGBA(c.canvas)

	draw.Draw(img, img.Bounds(), &image.Uniform{greyRGB}, image.Point{}, draw.Src)

	drawGraph(img)
	populateGraph(img, data, tPx)

	return img, nil
}

func drawGraph(img *image.RGBA) {
	yBar := image.NewRGBA(
		image.Rectangle{Min: image.Point{50, 100}, Max: image.Point{52, 900}},
	)
	draw.Draw(yBar, yBar.Bounds(), image.White, image.Point{}, draw.Src)

	xBar := image.NewRGBA(
		image.Rectangle{Min: image.Point{50, 898}, Max: image.Point{950, 900}},
	)
	draw.Draw(xBar, xBar.Bounds(), image.White, image.Point{}, draw.Src)

	draw.Draw(img, img.Bounds(), yBar, image.Point{}, draw.Src)
	draw.Draw(img, img.Bounds(), xBar, image.Point{}, draw.Src)
}

func populateGraph(img *image.RGBA, h *RGBHist, tPx uint32) {
	var total uint32
	lw := 900 / 260
	sp := 53
	ep := sp + lw
	for _, bounds := range h.R {
		total += bounds
		bar := image.NewRGBA(
			image.Rectangle{Min: image.Point{sp, 80}, Max: image.Point{ep, 900}},
		)
		sp = ep
		ep = ep + lw

		draw.Draw(bar, bar.Bounds(), image.White, image.Point{}, draw.Src)
		draw.Draw(img, img.Bounds(), bar, image.Point{}, draw.Src)
	}

}

func makeDefaultBChart() *BChart {
	return &BChart{
		canvas: image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{1000, 1000},
		},
		it: 100,
	}
}
