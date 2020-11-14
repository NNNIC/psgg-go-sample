// Original Copyright 2014 Hajime Hoshi
package main

import (
	"log"

	"github.com/NNNIC/psgg-go-sample/test-ebiten/testa/game"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Rotate (StateGo Ebiten Demo)")
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}
