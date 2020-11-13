// Original Copyright 2014 Hajime Hoshi

package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
    "github.com/NNNIC/psgg-go-sample/test-ebiten/test6/sm"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	gophersImage *ebiten.Image
)

type Game struct {
	count int
}

var drawfunc func()
var bDone = false

func (g *Game) Update(screen *ebiten.Image) error {
	if !bDone {
		bDone = true
		drawfunc = func() {
			w, h := gophersImage.Size()
			op := &ebiten.DrawImageOptions{}

			// Move the image's center to the screen's upper-left corner.
			// This is a preparation for rotating. When geometry matrices are applied,
			// the origin point is the upper-left corner.
			op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

			// Rotate the image. As a result, the anchor point of this rotate is
			// the center of the image.
			op.GeoM.Rotate(float64(g.count%360) * 2 * math.Pi / 360)

			// Move the image to the screen's center.
			op.GeoM.Translate(screenWidth/2, screenHeight/2)

			screen.DrawImage(gophersImage, op)
		}
		sm.TestControl()

	}

	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawfunc()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	img, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
	if err != nil {
		log.Fatal(err)
	}
	gophersImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Rotate (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
