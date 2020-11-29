package game

import (
	_ "image/jpeg"
	_ "image/png"
	"math"
	"time"

	"github.com/NNNIC/psgg-go-sample/ebiten/resources/sgimg"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth         = 640
	screenHeight        = 480
	maxDrawStageListLen = 500
	maxCommonListLen    = 100
)

// Game ...
type Game struct {
	counter     int
	MainControl func(bool, *Game) bool
	UpdateList  []func(bool, *Game) bool // UpdateList ... contains statego controllers

	// drawList
	DrawBgList               []func()
	DrawStageList            []func()
	DrawFeList               []func()
	DrawTermFunc             []func()
	ClrDrawStageListOnUpdate bool

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

	GophersImage  *ebiten.Image
	GophersImage2 *ebiten.Image
	GophersImage3 *ebiten.Image

	MascotImage  *ebiten.Image
	MascotAImage *ebiten.Image
	EbitenImage  *ebiten.Image

	// Gopher vs Nic
	BlkDataList    []*BlkData
	GopherDataList []*GopherData
	ScoreData0     *ScoreData
	GameOver       bool

	ScoreMul        int
	ScoreBase       int
	ScoreClickCount int
}

// AddUpdate ...
func (g *Game) AddUpdate(cf func(bool, *Game) bool) int {
	handle := len(g.UpdateList)
	g.UpdateList = append(g.UpdateList, cf)
	return handle
}

// ClrUpdate ...
func (g *Game) ClrUpdate() {
	if len(g.UpdateList) == 0 {
		g.UpdateList = make([]func(bool, *Game) bool, 0, maxCommonListLen)
	}
	g.UpdateList = g.UpdateList[:0]
}

// PlayUpdate ...
func (g *Game) PlayUpdate() {
	if g.ClrDrawStageListOnUpdate {
		g.ClrDrawStage()
	}

	g.MainControl(false, g) // Call Main

	for i := 0; i < len(g.UpdateList); i++ {
		g.UpdateList[i](false, g) // Call Others
	}
}

// AddDrawBg ... Add Draw Function to DrawBgList
func (g *Game) AddDrawBg(df func()) int {
	handle := len(g.DrawBgList)
	g.DrawBgList = append(g.DrawBgList, df)
	return handle
}

// SetDrawBg ...
func (g *Game) SetDrawBg(handle int, df func()) {
	g.DrawBgList[handle] = df
}

// AddDrawStage ... Add Draw Function to DrawStageList
func (g *Game) AddDrawStage(df func()) int {
	handle := len(g.DrawStageList)
	g.DrawStageList = append(g.DrawStageList, df)
	return handle
}

// SetDrawStage ...
func (g *Game) SetDrawStage(handle int, df func()) {
	g.DrawStageList[handle] = df
}

// AddDrawFe ... Add Draw Function to DrawFeList
func (g *Game) AddDrawFe(df func()) int {
	handle := len(g.DrawFeList)
	g.DrawFeList = append(g.DrawFeList, df)
	return handle
}

// SetDrawFe ...
func (g Game) SetDrawFe(handle int, df func()) {
	g.DrawFeList[handle] = df
}

// SetDbgTerm ...
func (g *Game) SetDbgTerm(df func()) {
	g.DrawTermFunc = g.DrawTermFunc[:0]
	g.DrawTermFunc = append(g.DrawTermFunc, df)
}

// ClrDrawBg ...
func (g *Game) ClrDrawBg() {
	if len(g.DrawBgList) == 0 {
		g.DrawBgList = make([]func(), 0, maxCommonListLen)
	}
	g.DrawBgList = g.DrawBgList[:0]
}

// ClrDrawStage ...
func (g *Game) ClrDrawStage() { // for reducing gabage memory issue.
	if len(g.DrawStageList) == 0 {
		g.DrawStageList = make([]func(), 0, maxDrawStageListLen)
	}
	g.DrawStageList = g.DrawStageList[:0]
}

// ClrDrawFe ...
func (g *Game) ClrDrawFe() {
	if len(g.DrawFeList) == 0 {
		g.DrawFeList = make([]func(), 0, maxCommonListLen)
	}
	g.DrawFeList = g.DrawFeList[:0]
}

// ClrDbgTerm ...
func (g *Game) ClrDbgTerm() {
	if len(g.DrawTermFunc) == 0 {
		g.DrawTermFunc = make([]func(), 0, maxCommonListLen)
	}
	g.DrawTermFunc = g.DrawTermFunc[:0]
}

// DoDraw ...
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

// ClearAll ...
func (g *Game) ClearAll() {
	g.ClrDrawStageListOnUpdate = false
	g.ClrUpdate()
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
	############# MANAGER APIs #############
*/

// Update ...
func (g *Game) Update() error {
	if bInitDone == false {
		bInitDone = true
		maincontrol := mainControl()
		//maincontrol(true, g)
		g.MainControl = maincontrol
	}

	//g.Screen = screen
	g.PlayUpdate()
	g.Count++
	return nil
}

// Draw ...
func (g *Game) Draw(screen *ebiten.Image) {
	g.Screen = screen

	g.DoDraw()
}

// Layout ...
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.ScreenWidth = screenWidth
	g.ScreenHeight = screenHeight
	return screenWidth, screenHeight
}

// TimeNowMs ...
func (g *Game) TimeNowMs() int64 {
	nowUTC := time.Now().UTC()
	return nowUTC.UnixNano() / int64(time.Millisecond)
}

// TermPrint ...
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

// TermClear ...
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
	op.GeoM.Rotate(float64(int(angle)%360) * 2 * math.Pi / 360)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	g.Screen.DrawImage(image, op)
}

/*
	############# USER API's #############
*/

// Gophers_jpg ...
func (g *Game) Gophers_jpg() []byte {
	return images.Gophers_jpg
}

func (g *Game) Gophers128_png() []byte {
	return sgimg.Gopher128_png
}
func (g *Game) Gophers64_png() []byte {
	return sgimg.Gopher64_png
}
func (g *Game) Gophers32_png() []byte {
	return sgimg.Gopher32_png
}

// Mascot64_png ...
func (g *Game) Mascot64_png() []byte {
	return sgimg.Mascot64_png
}

// Mascot32_png ...
func (g *Game) Mascot32_png() []byte {
	return sgimg.Mascot32_png
}

func (g *Game) MascotA32_png() []byte {
	return sgimg.MascotA32_png
}

// Mascot16_png ...
func (g *Game) Mascot16_png() []byte {
	return sgimg.Mascot16_png
}

func (g *Game) Ebiten32_png() []byte {
	return sgimg.Ebiten32_png
}
func (g *Game) EbitenP32_png() []byte {
	return sgimg.EbitenP32_png
}
func (g *Game) EbitenT32_png() []byte {
	return sgimg.EbitenT32_png
}
