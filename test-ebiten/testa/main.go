// Original Copyright 2014 Hajime Hoshi
package main

import (
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/NNNIC/psgg-go-sample/test-ebiten/test9/pg"
	"github.com/NNNIC/psgg-go-sample/test-ebiten/test9/sm"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	UpdateList []func(bool, *Game) bool // UpdateList ... contains statego controllers

	// drawList
	DrawBgList    []func()
	DrawStageList []func()
	DrawFeList    []func()

	// ebiten
	Screen *ebiten.Image

	// user
	ScreenWidth  int
	ScreenHeight int
	Count        int
}

func (g *Game) AddUpdate(cf func(bool, *Game) bool) int {
	hundle := len(g.UpdateList)
	g.UpdateList = append(g.UpdateList, cf)
	return hundle
}
func (g *Game) PlayUpdate() {
	for i := 0; i < len(g.UpdateList); i++ {
		g.UpdateList[i](false, g)
	}
}
func (g *Game) AddDrawBg(df func()) {
	g.DrawBgList = append(g.DrawBgList, df)
}
func (g *Game) AddDrawStage(df func()) {
	g.DrawStageList = append(g.DrawStageList, df)
}
func (g *Game) AddDrawFe(df func()) {
	g.DrawFeList = append(g.DrawFeList, df)
}

/*
	Init
*/
var bInitDone = false

func (g *Game) Update(screen *ebiten.Image) error {
	g.Screen = screen
	if bInitDone == false {
		bInitDone = true
		maincontrol := sm.MainControl()
		maincontrol(true)
		g.AddUpdate(maincontrol)
	}
	g.PlayUpdate()
	g.Count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Screen = screen

	pg.DoDraw()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.ScreenWidth = screenWidth
	g.ScreenHeight = screenHeight
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Rotate (StateGo Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
