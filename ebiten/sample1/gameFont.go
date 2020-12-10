package main

import (
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const sampleText = `The quick brown fox jumps over the lazy dog.`

var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
	jaKanjis        = []rune{}
)

func fontInit() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func textdraw(screen *ebiten.Image, x, y int, s string, col color.Color) {
	text.Draw(screen, s, mplusNormalFont, x, y, col)
}
func textdrawBig(screen *ebiten.Image, x, y int, s string, col color.Color) {
	text.Draw(screen, s, mplusBigFont, x, y, col)
}

// benri
func textdrawBigWFrame(screen *ebiten.Image, x, y int, s string, col, colFrame color.Color) {
	textdrawBig(screen, x+1, y+2, s, colFrame)
	textdrawBig(screen, x+2, y+1, s, colFrame)
	textdrawBig(screen, x-1, y-2, s, colFrame)
	textdrawBig(screen, x-2, y-1, s, colFrame)

	textdrawBig(screen, x+2, y+2, s, colFrame)
	textdrawBig(screen, x+2, y-2, s, colFrame)
	textdrawBig(screen, x-2, y+2, s, colFrame)
	textdrawBig(screen, x-2, y-2, s, colFrame)
	textdrawBig(screen, x, y, s, col)
}
