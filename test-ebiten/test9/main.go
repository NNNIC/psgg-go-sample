// Original Copyright 2014 Hajime Hoshi
package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"

	"github.com/NNNIC/psgg-go-sample/test-ebiten/test9/sm"
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

/*
	GLOBALS
*/
var Screen *ebiten.Image

/*
	State Machine Control Update
*/
var updateList []func(bool) bool
var updateNum = 0

// AddUpdateList ... Add control function into map. Reture handle
func AddUpdateList(cf func(bool) bool) int {
	hundle := len(updateList)
	updateList = append(updateList, cf)
	return hundle
}

/*
	Draw List
*/
var DrawBgList []func()
var DrawStageList []func()
var DrawFeList []func()

/*
	Init
*/
var bInitDone = false

func (g *Game) Update(screen *ebiten.Image) error {
	if bInitDone == false {
		bInitDone = true
		maincontrol := sm.MainControl()
		AddUpdateList(maincontrol)
	}
	Screen = screen

	for i := 0; i < len(updateList); i++ {
		(updateList[i])(false)
	}
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(DrawBgList); i++ {
		(DrawBgList[i])()
	}
	for i := 0; i < len(DrawStageList); i++ {
		(DrawBgList[i])()
	}
	for i := 0; i < len(DrawFeList); i++ {
		(DrawBgList[i])()
	}
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
