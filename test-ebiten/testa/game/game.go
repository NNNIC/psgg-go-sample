// Original Copyright 2014 Hajime Hoshi
package game

import (
	_ "image/jpeg"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
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

	GophersImage *ebiten.Image
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
func (g *Game) DoDraw() {
	for i := 0; i < len(g.DrawBgList); i++ {
		(g.DrawBgList[i])()
	}
	for i := 0; i < len(g.DrawStageList); i++ {
		(g.DrawStageList[i])()
	}
	for i := 0; i < len(g.DrawFeList); i++ {
		(g.DrawFeList[i])()
	}

}

/*
	Init
*/
var bInitDone = false

func (g *Game) Update(screen *ebiten.Image) error {
	if bInitDone == false {
		bInitDone = true
		maincontrol := MainControl()
		//maincontrol(true, g)
		g.AddUpdate(maincontrol)
	}

	g.Screen = screen
	g.PlayUpdate()
	g.Count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Screen = screen

	g.DoDraw()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.ScreenWidth = screenWidth
	g.ScreenHeight = screenHeight
	return screenWidth, screenHeight
}

func (g *Game) Gophers_jpg() []byte {
	return images.Gophers_jpg
}

func (g *Game) TimeNowMs() int64 {
	nowUTC := time.Now().UTC()
	return nowUTC.UnixNano() / int64(time.Millisecond)
}
