package game

import (
	_ "image/jpeg"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
	DrawTermFunc  []func()

	// ebiten
	Screen *ebiten.Image

	// debug terminal message
	dbgtermmsg    []string
	dbgtermmsgstr string

	// user
	ScreenWidth  int
	ScreenHeight int
	Count        int
	BgImage      *ebiten.Image

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
func (g *Game) SetDbgTerm(df func()) {
	g.DrawTermFunc = g.DrawTermFunc[:0]
	g.DrawTermFunc = append(g.DrawTermFunc, df)
}

func (g *Game) ClrDrawBg() {
	g.DrawBgList = g.DrawBgList[:0]
}
func (g *Game) ClrDrawStage() {
	g.DrawStageList = g.DrawStageList[:0]
}
func (g *Game) ClrDrawFe() {
	g.DrawFeList = g.DrawFeList[:0]
}
func (g *Game) ClrDbgTerm() {
	g.DrawTermFunc = g.DrawTermFunc[:0]
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
	if len(g.DrawTermFunc) > 0 {
		g.DrawTermFunc[0]()
	}
}
func (g *Game) ClearAll() {
	g.ClrDrawBg()
	g.ClrDrawStage()
	g.ClrDrawFe()
	g.TermClear()
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

func (g *Game) TermPrint(str string) {
	/*
		53 chars * 15 lines
		1char 6x16
	*/
	maxline := screenHeight / 16
	g.dbgtermmsg = append(g.dbgtermmsg, str)
	termlen := len(g.dbgtermmsg)
	if termlen > maxline {
		g.dbgtermmsg = g.dbgtermmsg[1:]
	}
	g.dbgtermmsgstr = ""
	for i := 0; i < len(g.dbgtermmsg); i++ {
		if g.dbgtermmsgstr != "" {
			g.dbgtermmsgstr += "\n"
		}
		g.dbgtermmsgstr += g.dbgtermmsg[i]
	}
	f := func() {
		ebitenutil.DebugPrintAt(g.Screen, g.dbgtermmsgstr, 0, 0)
	}
	g.SetDbgTerm(f)
}
func (g *Game) TermClear() {
	g.dbgtermmsg = g.dbgtermmsg[:0]
	g.ClrDbgTerm()
}
