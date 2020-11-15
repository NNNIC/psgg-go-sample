package game

import (
	_ "image/jpeg"
	_ "image/png"
	"time"

	"github.com/NNNIC/psgg-go-sample/ebiten/resources/sgimg"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
)

const (
	screenWidth  = 640
	screenHeight = 480
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

	// fade image
	FadeImage *ebiten.Image

	// user
	ScreenWidth  int
	ScreenHeight int
	Count        int
	BgImage      *ebiten.Image

	GophersImage *ebiten.Image
	MascotImage  *ebiten.Image
}

func (g *Game) AddUpdate(cf func(bool, *Game) bool) int {
	handle := len(g.UpdateList)
	g.UpdateList = append(g.UpdateList, cf)
	return handle
}
func (g *Game) PlayUpdate() {
	for i := 0; i < len(g.UpdateList); i++ {
		g.UpdateList[i](false, g)
	}
}

// AddDrawBg ... Add Draw Function to DrawBgList
func (g *Game) AddDrawBg(df func()) int {
	handle := len(g.DrawBgList)
	g.DrawBgList = append(g.DrawBgList, df)
	return handle
}
func (g *Game) SetDrawBg(handle int, df func()) {
	g.DrawBgList[handle] = df
}

// AddDrawStage ... Add Draw Function to DrawStageList
func (g *Game) AddDrawStage(df func()) int {
	handle := len(g.DrawStageList)
	g.DrawStageList = append(g.DrawStageList, df)
	return handle
}
func (g *Game) SetDrawStage(handle int, df func()) {
	g.DrawStageList[handle] = df
}

// AddDrawFe ... Add Draw Function to DrawFeList
func (g *Game) AddDrawFe(df func()) int {
	handle := len(g.DrawFeList)
	g.DrawFeList = append(g.DrawFeList, df)
	return handle
}
func (g Game) SetDrawFe(handle int, df func()) {
	g.DrawFeList[handle] = df
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

/*
	MANAGER APIs
*/
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

// DrawImage ... simple draw image
// * screen's origin is upper left
func (g *Game) DrawImage(image *ebiten.Image, x, y, angle, scale float64) {
	w, h := image.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	g.Screen.DrawImage(image, op)
}

/*
	USER API's
*/
func (g *Game) Gophers_jpg() []byte {
	return images.Gophers_jpg
}
func (g *Game) Mascot64_png() []byte {
	return sgimg.Mascot64_png
}
func (g *Game) Mascot32_png() []byte {
	return sgimg.Mascot32_png
}
func (g *Game) Mascot16_png() []byte {
	return sgimg.Mascot16_png
}
func (g *Game) Clamp255(i int) int {
	if i < 0 {
		return 0
	}
	if i > 255 {
		return 255
	}
	return i
}
