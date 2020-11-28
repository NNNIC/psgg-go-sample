package game

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Gamemain ... invoke game
func Gamemain() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Rotate (StateGo Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

// MainControl ... main flow
func mainControl() func(bool, *Game) bool {
	var g *Game

	rand.Seed(time.Now().Unix())

	curfunc := -1
	nextfunc := -1

	nowait := false

	noWait := func() {
		nowait = true
	}

	gotoState := func(st int) {
		nextfunc = st
	}
	hasNextState := func() bool {
		return nextfunc != -1
	}

	callstack := [...]int{0, 0, 0, 0, 0}
	callstacklevel := 0
	gosubState := func(sub, next int) {
		if callstacklevel >= len(callstack)-1 {
			panic("call stack overflow")
			//return
		}
		callstack[callstacklevel] = next
		callstacklevel++
		gotoState(sub)
	}
	returnState := func() {
		if callstacklevel <= 0 {
			panic("call stack underflow")
			//return
		}
		callstacklevel--
		funcid := callstack[callstacklevel]
		gotoState(funcid)
	}
	// #
	// # Define function ID for state
	// #
	id := 0
	//[STATEGO OUTPUT START] indent(4) $/^S_./->#setids$
    //             psggConverterLib.dll converted from psgg-file:mainControl.psgg

    funcIds1stRotation := id
    id++
    funcIds2ndRotation := id
    id++
    funcIds3rdRotation := id
    id++
    funcIdsBACKTOMENU := id
    id++
    funcIdsCHANGEBG := id
    id++
    funcIdsCLEARALL := id
    id++
    funcIdsCLEARALL1 := id
    id++
    funcIdsCLEARALL2 := id
    id++
    funcIdsCLEARTERM := id
    id++
    funcIdsCOUNT := id
    id++
    funcIdsDISPLAYBLK := id
    id++
    funcIdsDISPLAYGOPHER := id
    id++
    funcIdsDISPLAYSCORE := id
    id++
    funcIdsDrawMascot := id
    id++
    funcIdsEND := id
    id++
    funcIdsESC := id
    id++
    funcIdsFADEIN := id
    id++
    funcIdsFADEIN1 := id
    id++
    funcIdsFADEINTEST := id
    id++
    funcIdsFADEOUT := id
    id++
    funcIdsFADEOUT1 := id
    id++
    funcIdsFADEOUT2 := id
    id++
    funcIdsGAMEOVER := id
    id++
    funcIdsGAMEOVER1 := id
    id++
    funcIdsINIT := id
    id++
    funcIdsINIT1 := id
    id++
    funcIdsISRELEASE := id
    id++
    funcIdsLOADEBITEN := id
    id++
    funcIdsLOADGOPHER := id
    id++
    funcIdsLOADNIC := id
    id++
    funcIdsLOP000 := id
    id++
    funcIdsLOP000LoopCheck := id
    id++
    funcIdsLOP000LoopNext := id
    id++
    funcIdsLOP001 := id
    id++
    funcIdsLOP001LoopCheck := id
    id++
    funcIdsLOP001LoopNext := id
    id++
    funcIdsMENU := id
    id++
    funcIdsOPENING := id
    id++
    funcIdsOPENING1 := id
    id++
    funcIdsPAS000 := id
    id++
    funcIdsPAS001 := id
    id++
    funcIdsPAS002 := id
    id++
    funcIdsRET000 := id
    id++
    funcIdsRET001 := id
    id++
    funcIdsRotationOverRay := id
    id++
    funcIdsSBS000 := id
    id++
    funcIdsSBS001 := id
    id++
    funcIdsSETBG := id
    id++
    funcIdsSETBG1 := id
    id++
    funcIdsSETBGBLUE := id
    id++
    funcIdsSETBGGREEN := id
    id++
    funcIdsSETBGRED := id
    id++
    funcIdsSHOWMASCOT := id
    id++
    funcIdsSHOWMASCOT1 := id
    id++
    funcIdsSTART := id
    id++
    funcIdsTERMINAL := id
    id++
    funcIdsTERMINAL1 := id
    id++
    funcIdsTEXT := id
    id++
    funcIdsTIMEOUT := id
    id++
    funcIdsVERSION := id
    id++
    funcIdsWAIT := id
    id++
    funcIdsWAITCLRDRAW := id
    id++
    funcIdsWAIT1 := id
    id++
    funcIdsWAIT10 := id
    id++
    funcIdsWAIT11 := id
    id++
    funcIdsWAIT12 := id
    id++
    funcIdsWAIT13 := id
    id++
    funcIdsWAIT2 := id
    id++
    funcIdsWAIT3 := id
    id++
    funcIdsWAIT4 := id
    id++
    funcIdsWAIT5 := id
    id++
    funcIdsWAIT6 := id
    id++
    funcIdsWAIT7 := id
    id++
    funcIdsWAIT8 := id
    id++
    funcIdsWAIT9 := id
    id++


	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^S_./->#memlist$
    //             psggConverterLib.dll converted from psgg-file:mainControl.psgg

    var cntsCOUNT int
    var goalsFADEIN float64
    var goalsFADEIN1 float64
    var goalsFADEOUT float64
    var goalsFADEOUT1 float64
    var loop1 = 0
    var loop2 = 0
    var xsOPENING, ysOPENING int
    var timesWAIT int64
    var timesWAITCLRDRAW int64
    var timesWAIT1 int64
    var timesWAIT10 int64
    var timesWAIT11 int64
    var timesWAIT12 int64
    var timesWAIT13 int64
    var timesWAIT2 int64
    var timesWAIT3 int64
    var timesWAIT4 int64
    var timesWAIT5 int64
    var timesWAIT6 int64
    var timesWAIT7 int64
    var timesWAIT8 int64
    var timesWAIT9 int64


	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^E_./$
    //             psggConverterLib.dll converted from psgg-file:mainControl.psgg

    /*
        E_0000
    */
    release:=true
    /*
        E_GAME_ETC
    */
    var gdg func()
    var gde func()
    var gdt func()
    var gds func()
    var gdp func()
    var gdv func()


	//[STATEGO OUTPUT END]
	// USER API
	createBlk := func() {
		maxX := 20
		maxY := 12
		w := 24
		for y := 0; y < maxY; y++ {
			for x := 0; x < maxX; x++ {
				posX := g.ScreenWidth/2 + w/2 - (w * maxX / 2) + x*w
				posY := g.ScreenHeight/2 - (w * maxY / 2) + y*w

				d := BlkData{PosX: float64(posX), PosY: float64(posY)}
				g.BlkDataList = append(g.BlkDataList, &d)
				g.AddUpdate(blkControl(&d))
			}
		}
	}
	createGopher := func(id int) {
		dx := 0
		if id > 0 {
			sym := 1
			if id%2 == 0 {
				sym = -1
			}
			dx = (id + 1) / 2 * 38 * sym
		}

		d := GopherData{PosX: float64(g.ScreenWidth/2 + dx), PosY: float64(g.ScreenHeight - 20)}
		g.GopherDataList = append(g.GopherDataList, &d)
		g.AddUpdate(gopherControl(&d, id))
	}
	createScore := func() {
		d := ScoreData{}
		g.ScoreData0 = &d
		g.AddUpdate(scoreControl(&d))
	}
	// #
	// #  State Function
	// #
	var endofFuncList = func(bFirst bool) { // for end of function list
	}
	//[STATEGO OUTPUT START] indent(4) $/^S_./$
    //             psggConverterLib.dll converted from psgg-file:mainControl.psgg

    /*
        S_1stRotation
    */
    s1stRotation := func( bFirst  bool ) {
        if bFirst {
            if g.GophersImage == nil {
                img, _, err := image.Decode(bytes.NewReader(g.Gophers_jpg()))
                if err != nil {
            	log.Fatal(err)
                }
                g.GophersImage = ebiten.NewImageFromImage(img )
            }
            drawfunc := func() {
            	w, h := g.GophersImage.Size()
            	op := &ebiten.DrawImageOptions{}
            	// Move the image's center to the screen's upper-left corner.
            	// This is a preparation for rotating. When geometry matrices are applied,
            	// the origin point is the upper-left corner.
            	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
            	// Rotate the image. As a result, the anchor point of this rotate is
            	// the center of the image.
            	op.GeoM.Rotate(float64(g.Count % 360) * 2 * math.Pi / 360 )
            	// Move the image to the screen's center.
            	op.GeoM.Translate(float64(g.ScreenWidth/2), float64(g.ScreenHeight/2))
            	g.Screen.DrawImage(g.GophersImage, op)
            }
            g.AddDrawStage(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT)
        }
    }
    /*
        S_2ndRotation
    */
    s2ndRotation := func( bFirst  bool ) {
        if bFirst {
            drawfunc := func() {
            	w, h := g.GophersImage.Size()
            	op := &ebiten.DrawImageOptions{}
            	// Move the image's center to the screen's upper-left corner.
            	// This is a preparation for rotating. When geometry matrices are applied,
            	// the origin point is the upper-left corner.
            	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
            	// Rotate the image. As a result, the anchor point of this rotate is
            	// the center of the image.
            	op.GeoM.Rotate(float64((g.Count*3)%360) * 2 * math.Pi / 360)
            	op.GeoM.Scale(0.5, 0.5)
            	// Move the image to the screen's center.
            	op.GeoM.Translate(float64(g.ScreenWidth/2), float64(g.ScreenHeight/2))
            	g.Screen.DrawImage(g.GophersImage, op)
            }
            g.AddDrawStage(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT6)
        }
    }
    /*
        S_3rdRotation
    */
    s3rdRotation := func( bFirst  bool ) {
        if bFirst {
            drawfunc := func() {
            	w, h := g.GophersImage.Size()
            	op := &ebiten.DrawImageOptions{}
            	// Move the image's center to the screen's upper-left corner.
            	// This is a preparation for rotating. When geometry matrices are applied,
            	// the origin point is the upper-left corner.
            	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
            	// Rotate the image. As a result, the anchor point of this rotate is
            	// the center of the image.
            	op.GeoM.Rotate(float64((g.Count*6)%360) * 2 * math.Pi / 360)
            	op.GeoM.Scale(0.2, 0.2)
            	// Move the image to the screen's center.
            	op.GeoM.Translate(float64(g.ScreenWidth/2), float64(g.ScreenHeight/2))
            	g.Screen.DrawImage(g.GophersImage, op)
            }
            g.AddDrawStage(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsBACKTOMENU)
        }
    }
    /*
        S_BACKTO_MENU
    */
    sBACKTOMENU := func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsISRELEASE)
        }
    }
    /*
        S_CHANGEBG
    */
    sCHANGEBG := func( bFirst  bool ) {
        if bFirst {
            g.TermPrint("= BACKGROUND COLOR CHANGE =");
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT7)
        }
    }
    /*
        S_CLEARALL
    */
    sCLEARALL := func( bFirst  bool ) {
        if bFirst {
            g.ClearAll()
        }
        if !hasNextState() {
            gotoState(funcIdsBACKTOMENU)
        }
    }
    /*
        S_CLEARALL1
    */
    sCLEARALL1 := func( bFirst  bool ) {
        if bFirst {
            g.ClearAll()
        }
        if !hasNextState() {
            gotoState(funcIdsSETBG)
        }
    }
    /*
        S_CLEARALL2
    */
    sCLEARALL2 := func( bFirst  bool ) {
        if bFirst {
            g.ClearAll()
        }
        if !hasNextState() {
            gotoState(funcIdsFADEOUT2)
        }
    }
    /*
        S_CLEARTERM
    */
    sCLEARTERM := func( bFirst  bool ) {
        if bFirst {
            g.TermClear()
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT3)
        }
    }
    /*
        S_COUNT
    */
    sCOUNT := func( bFirst  bool ) {
        if bFirst {
            cntsCOUNT = 0
        }
        g.TermPrint("Count : " + strconv.Itoa(cntsCOUNT))
        cntsCOUNT++
        if cntsCOUNT < 100 {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT4)
        }
    }
    /*
        S_DISPLAY_BLK
    */
    sDISPLAYBLK := func( bFirst  bool ) {
        if bFirst {
            createBlk()
        }
        if !hasNextState() {
            gotoState(funcIdsDISPLAYGOPHER)
        }
    }
    /*
        S_DISPLAY_GOPHER
    */
    sDISPLAYGOPHER := func( bFirst  bool ) {
        if bFirst {
            createGopher(0)
            createGopher(1)
            createGopher(2)
            createGopher(3)
            createGopher(4)
        }
        if !hasNextState() {
            gotoState(funcIdsDISPLAYSCORE)
        }
    }
    /*
        S_DISPLAY_SCORE
    */
    sDISPLAYSCORE := func( bFirst  bool ) {
        if bFirst {
            createScore()
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT13)
        }
    }
    /*
        S_DrawMascot
    */
    sDrawMascot := func( bFirst  bool ) {
        if bFirst {
            x := loop1
            y := loop2
            drawfunc := func() {
                g.DrawImage(g.MascotImage, float64(x*64+32),float64(y*64+32),0,1)
            }
            g.AddDrawStage(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsRET001)
        }
    }
    /*
        S_END
    */
    sEND := func ( bFirst  bool ) {
         // end of state machine
    }
    /*
        S_ESC
    */
    sESC := func( bFirst  bool ) {
        if bFirst {
            fmt.Println("S_ESC")
        }
        if !hasNextState() {
            gotoState(funcIdsCLEARALL2)
        }
    }
    /*
        S_FADEIN
    */
    sFADEIN := func( bFirst  bool ) {
        fadestepnum := 60.0
        col := color.RGBA{255,255,255,255}
        if bFirst {
            if g.FadeImage == nil {
                img := ebiten.NewImage(g.ScreenWidth, g.ScreenHeight)
                g.FadeImage = img
            }
            goalsFADEIN = float64(g.Count) + fadestepnum
            drawfunc := func() {
                alpha := (goalsFADEIN - float64(g.Count)) / fadestepnum * 255.0
                g.FadeImage.Fill(color.RGBA{col.R,col.G,col.B,uint8(clamp255(int(alpha)))})
                w, h := g.FadeImage.Size()
                op := &ebiten.DrawImageOptions{}
                op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
                op.GeoM.Translate(float64(g.ScreenWidth/2), float64(g.ScreenHeight/2))
                g.Screen.DrawImage(g.FadeImage, op)
            }
            g.ClrDrawFe()
            g.AddDrawFe(drawfunc)
        }
        if float64(g.Count) < goalsFADEIN {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsTEXT)
        }
    }
    /*
        S_FADEIN1
    */
    sFADEIN1 := func( bFirst  bool ) {
        fadestepnum := 60.0
        col := color.RGBA{255,255,255,255}
        if bFirst {
            if g.FadeImage == nil {
                img := ebiten.NewImage(g.ScreenWidth, g.ScreenHeight)
                g.FadeImage = img
            }
            goalsFADEIN1 = float64(g.Count) + fadestepnum
            drawfunc := func() {
                alpha := (goalsFADEIN1 - float64(g.Count)) / fadestepnum * 255.0
                g.FadeImage.Fill(color.RGBA{col.R,col.G,col.B,uint8(clamp255(int(alpha)))})
                w, h := g.FadeImage.Size()
                op := &ebiten.DrawImageOptions{}
                op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
                op.GeoM.Translate(float64(g.ScreenWidth/2), float64(g.ScreenHeight/2))
                g.Screen.DrawImage(g.FadeImage, op)
            }
            g.ClrDrawFe()
            g.AddDrawFe(drawfunc)
        }
        if float64(g.Count) < goalsFADEIN1 {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsWAITCLRDRAW)
        }
    }
    /*
        S_FADEINTEST
    */
    sFADEINTEST := func( bFirst  bool ) {
        if bFirst {
            g.TermPrint(" = FADE OUT -> IN = ")
        }
        if !hasNextState() {
            gotoState(funcIdsFADEOUT)
        }
    }
    /*
        S_FADEOUT
    */
    sFADEOUT := func( bFirst  bool ) {
        fadestepnum := 60.0
        col := color.RGBA{255,255,255,255}
        if bFirst {
            if g.FadeImage == nil {
                img := ebiten.NewImage(g.ScreenWidth, g.ScreenHeight)
                g.FadeImage = img
            }
            goalsFADEOUT = float64(g.Count) + fadestepnum
            drawfunc := func() {
                alpha := (fadestepnum - (goalsFADEOUT - float64(g.Count))) / fadestepnum * 255.0
                g.FadeImage.Fill(color.RGBA{col.R,col.G,col.B,clamp255(int(alpha))})
                w, h := g.FadeImage.Size()
                op := &ebiten.DrawImageOptions{}
                op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
                op.GeoM.Translate(float64(g.ScreenWidth/2), float64(g.ScreenHeight/2))
                g.Screen.DrawImage(g.FadeImage, op)
            }
            g.ClrDrawFe()
            g.AddDrawFe(drawfunc)
        }
        if float64(g.Count) < goalsFADEOUT {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsFADEIN)
        }
    }
    /*
        S_FADEOUT1
    */
    sFADEOUT1 := func( bFirst  bool ) {
        fadestepnum := 60.0
        col := color.RGBA{255,255,255,255}
        if bFirst {
            if g.FadeImage == nil {
                img := ebiten.NewImage(g.ScreenWidth, g.ScreenHeight)
                g.FadeImage = img
            }
            goalsFADEOUT1 = float64(g.Count) + fadestepnum
            drawfunc := func() {
                alpha := (fadestepnum - (goalsFADEOUT1 - float64(g.Count))) / fadestepnum * 255.0
                g.FadeImage.Fill(color.RGBA{col.R,col.G,col.B,clamp255(int(alpha))})
                w, h := g.FadeImage.Size()
                op := &ebiten.DrawImageOptions{}
                op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
                op.GeoM.Translate(float64(g.ScreenWidth/2), float64(g.ScreenHeight/2))
                g.Screen.DrawImage(g.FadeImage, op)
            }
            g.ClrDrawFe()
            g.AddDrawFe(drawfunc)
        }
        if float64(g.Count) < goalsFADEOUT1 {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsSETBG1)
        }
    }
    /*
        S_FADEOUT2
    */
    sFADEOUT2 := func( bFirst  bool ) {
        gotoState(funcIdsPAS002)
        noWait()
    }
    /*
        S_GAMEOVER
    */
    sGAMEOVER := func( bFirst  bool ) {
        if bFirst {
            fmt.Println("S_GAMEOVER")
        }
        if !hasNextState() {
            gotoState(funcIdsGAMEOVER1)
        }
    }
    /*
        S_GAMEOVER1
    */
    sGAMEOVER1 := func( bFirst  bool ) {
        if bFirst {
            gdp = func() {
                textdrawBigWFrame(g.Screen, 220, 220, "PERFECT", color.White, color.White)
            }
            gds = func() {
                textdraw(g.Screen, 200, 400, "PUSH SPACE TO END", color.White)
            }
        }
        g.AddDrawStage(gdp)
        g.AddDrawStage(gds)
        if !ebiten.IsKeyPressed(ebiten.KeySpace) && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && ebiten.TouchIDs() == nil {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsCLEARALL2)
        }
    }
    /*
        S_INIT
    */
    sINIT := func( bFirst  bool ) {
        if bFirst {
            g.ClearAll()
            fontInit()
        }
        if !hasNextState() {
            gotoState(funcIdsISRELEASE)
        }
    }
    /*
        S_INIT1
    */
    sINIT1 := func( bFirst  bool ) {
        if bFirst {
            g.ClrDrawStageListOnUpdate = true
            g.GameOver = false
            g.ScoreMul = 1
            g.ScoreBase = 1000
        }
        if !hasNextState() {
            gotoState(funcIdsLOADGOPHER)
        }
    }
    /*
        S_IS_RELEASE
    */
    sISRELEASE := func( bFirst  bool ) {
        if release {
            gotoState( funcIdsSHOWMASCOT1 )
        } else {
            gotoState( funcIdsMENU )
        }
    }
    /*
        S_LOAD_EBITEN
    */
    sLOADEBITEN := func( bFirst  bool ) {
        if bFirst {
            img, _, err := image.Decode(bytes.NewReader(g.Ebiten32_png()))
            if err != nil {
                log.Fatal(err)
            }
            g.EbitenImage = ebiten.NewImageFromImage(img)
            gde = func() {
            	g.DrawImage(g.EbitenImage,400,250,0,1)
            	g.DrawImage(g.EbitenImage,400,290,0,1)
            	g.DrawImage(g.EbitenImage,400,330,0,1)
            	g.DrawImage(g.EbitenImage,450,250,0,1)
            	g.DrawImage(g.EbitenImage,450,290,0,1)
            	g.DrawImage(g.EbitenImage,450,330,0,1)
            }
        }
        if !hasNextState() {
            gotoState(funcIdsVERSION)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_LOAD_GOPHER
    */
    sLOADGOPHER := func( bFirst  bool ) {
        if bFirst {
            img, _, err := image.Decode(bytes.NewReader(g.Gophers128_png()))
            if err != nil {
                log.Fatal(err)
            }
            g.GophersImage2 = ebiten.NewImageFromImage(img)
            gdg = func() {
                g.DrawImage(g.GophersImage2, 315, 290, 0,1)
            }
        }
        if !hasNextState() {
            gotoState(funcIdsLOADEBITEN)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_LOAD_NIC
    */
    sLOADNIC := func( bFirst  bool ) {
        if bFirst {
            if g.MascotImage == nil {
                img, _, err := image.Decode(bytes.NewReader(g.Mascot64_png()))
                if err != nil {
            	log.Fatal(err)
                }
                g.MascotImage = ebiten.NewImageFromImage(img )
            }
            drawfunc := func() {
            	g.DrawImage(g.MascotImage,8,8,0,1)
            }
            g.AddDrawStage(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsCLEARALL1)
        }
    }
    /*
        S_LOP000
    */
    sLOP000 := func ( bFirst bool ) {
        loop1 = 0
        gotoState(funcIdsLOP000LoopCheck)
        noWait()
    }
    sLOP000LoopCheck := func ( bFirst bool ) {
        if loop1 < (g.ScreenWidth / 64) {
            gosubState(funcIdsSBS000, funcIdsLOP000LoopNext)
        } else {
            gotoState(funcIdsPAS000)
        }
        noWait()
    }
    sLOP000LoopNext := func(bFirst bool ) {
        loop1++
        gotoState(funcIdsLOP000LoopCheck)
        noWait()
    }
    /*
        S_LOP001
    */
    sLOP001 := func ( bFirst bool ) {
        loop2 = 0
        gotoState(funcIdsLOP001LoopCheck)
        noWait()
    }
    sLOP001LoopCheck := func ( bFirst bool ) {
        if loop2 < (g.ScreenHeight / 64) {
            gosubState(funcIdsSBS001, funcIdsLOP001LoopNext)
        } else {
            gotoState(funcIdsRET000)
        }
        noWait()
    }
    sLOP001LoopNext := func(bFirst bool ) {
        loop2++
        gotoState(funcIdsLOP001LoopCheck)
        noWait()
    }
    /*
        S_MENU
    */
    sMENU := func( bFirst  bool ) {
        if bFirst {
            g.TermPrint("==== TEST =====")
            g.TermPrint("Push 1 ... Termainal Test")
            g.TermPrint("Push 2 ... Background Color Change Test")
            g.TermPrint("Push 3 ... Ebiten Rotaion Overlay Demo")
            g.TermPrint("Push 4 ... StateGo Mascot Demo")
            g.TermPrint("Push G ... Gopher vs Nic")
            g.TermPrint("Push T ... Try a new feature")
            g.TermPrint("")
            g.TermPrint("Push C ... Cear all");
            g.TermPrint("")
            g.TermPrint("PLEASE PUSH KEY!")
        }
        if ebiten.IsKeyPressed(ebiten.Key1) {
            gotoState( funcIdsTERMINAL )
        } else if ebiten.IsKeyPressed(ebiten.Key2) {
            gotoState( funcIdsCHANGEBG )
        } else if ebiten.IsKeyPressed(ebiten.Key3) {
            gotoState( funcIdsRotationOverRay )
        } else if ebiten.IsKeyPressed(ebiten.Key4) {
            gotoState( funcIdsSHOWMASCOT )
        } else if ebiten.IsKeyPressed(ebiten.KeyG) {
            gotoState( funcIdsSHOWMASCOT1 )
        } else if ebiten.IsKeyPressed(ebiten.KeyT) {
            gotoState( funcIdsFADEINTEST )
        } else if ebiten.IsKeyPressed(ebiten.KeyC) {
            gotoState( funcIdsCLEARALL )
        }
    }
    /*
        S_OPENING
    */
    sOPENING := func( bFirst  bool ) {
        if bFirst {
            xsOPENING = 50
            ysOPENING = 100
            gdt = func() {
                x:=xsOPENING
                y:=ysOPENING
                textdrawBig(g.Screen, clamp(x,-1000,207), y,     "GOPHERS", color.White)
                textdrawBig(g.Screen, clamp(x-100,-1000,242), y+50,  " LOVE", color.White)
                textdrawBig(g.Screen, clamp(x-200,-1000,242), y+100, "EBITEN", color.White)
            }
        }
        //if !ebiten.IsKeyPressed(ebiten.KeyT) {
        //    xsOPENING++
        //}
        xsOPENING+=2
        g.AddDrawStage(gdg)
        g.AddDrawStage(gde)
        g.AddDrawStage(gdt)
        //fmt.Println(xsOPENING)
        g.AddDrawStage(gdv)
        if xsOPENING < 450 {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsOPENING1)
        }
    }
    /*
        S_OPENING1
    */
    sOPENING1 := func( bFirst  bool ) {
        if bFirst {
            gds = func() {
                textdraw(g.Screen, 200, 400, "PUSH SPACE TO START", color.White)
            }
        }
        g.AddDrawStage(gdg)
        g.AddDrawStage(gde)
        g.AddDrawStage(gdt)
        g.AddDrawStage(gds)
        g.AddDrawStage(gdv)
        if !ebiten.IsKeyPressed(ebiten.KeySpace) && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && ebiten.TouchIDs()== nil {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsFADEOUT1)
        }
    }
    /*
        S_PAS000
    */
    sPAS000 := func( bFirst  bool ) {
        gotoState(funcIdsBACKTOMENU)
        noWait()
    }
    /*
        S_PAS001
    */
    sPAS001 := func( bFirst  bool ) {
        gotoState(funcIdsBACKTOMENU)
        noWait()
    }
    /*
        S_PAS002
    */
    sPAS002 := func( bFirst  bool ) {
        gotoState(funcIdsBACKTOMENU)
        noWait()
    }
    /*
        S_RET000
    */
    sRET000 := func ( bFirst bool ) {
        returnState()
        noWait()
    }
    /*
        S_RET001
    */
    sRET001 := func ( bFirst bool ) {
        returnState()
        noWait()
    }
    /*
        S_RotationOverRay
    */
    sRotationOverRay := func( bFirst  bool ) {
        if bFirst {
            g.TermPrint("= ROTAION OVERLAY = ")
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT5)
        }
    }
    /*
        S_SBS000
    */
    sSBS000 := func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsLOP001)
        }
    }
    /*
        S_SBS001
    */
    sSBS001 := func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsDrawMascot)
        }
    }
    /*
        S_SETBG
    */
    sSETBG := func( bFirst  bool ) {
        if bFirst {
            drawfunc := func() {
                g.Screen.Fill(color.RGBA{0,56,133,0xff})
            }
            g.AddDrawBg(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsLOP000)
        }
    }
    /*
        S_SETBG1
    */
    sSETBG1 := func( bFirst  bool ) {
        if bFirst {
            drawfunc := func() {
                g.Screen.Fill(color.RGBA{0,56,133,0xff})
            }
            g.AddDrawBg(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsDISPLAYBLK)
        }
    }
    /*
        S_SETBGBLUE
    */
    sSETBGBLUE := func( bFirst  bool ) {
        if bFirst {
            drawfunc := func() {
                g.Screen.Fill(color.RGBA{0x00,0xff,0,0xff})
            }
            g.AddDrawBg(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT9)
        }
    }
    /*
        S_SETBGGREEN
    */
    sSETBGGREEN := func( bFirst  bool ) {
        if bFirst {
            drawfunc := func() {
                g.Screen.Fill(color.RGBA{0x00,0x00,0xff,0xff})
            }
            g.AddDrawBg(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT10)
        }
    }
    /*
        S_SETBGRED
    */
    sSETBGRED := func( bFirst  bool ) {
        if bFirst {
            drawfunc := func() {
                g.Screen.Fill(color.RGBA{0xff,0,0,0xff})
            }
            g.AddDrawBg(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT8)
        }
    }
    /*
        S_SHOWMASCOT
    */
    sSHOWMASCOT := func( bFirst  bool ) {
        if bFirst {
            g.TermPrint(" = STATE GO MASCOT DEMO = ")
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT11)
        }
    }
    /*
        S_SHOWMASCOT1
    */
    sSHOWMASCOT1 := func( bFirst  bool ) {
        if bFirst {
            g.TermClear()
        }
        if !hasNextState() {
            gotoState(funcIdsINIT1)
        }
    }
    /*
        S_START
    */
    sSTART := func( bFirst  bool ) {
        gotoState(funcIdsINIT)
    }
    /*
        S_TERMINAL
    */
    sTERMINAL := func( bFirst  bool ) {
        if bFirst {
            g.TermPrint("= TERMAINAL TEST = ")
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT1)
        }
    }
    /*
        S_TERMINAL1
    */
    sTERMINAL1 := func( bFirst  bool ) {
        if bFirst {
            g.TermPrint("Count 0 to 99")
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT2)
        }
    }
    /*
        S_TEXT
    */
    sTEXT := func( bFirst  bool ) {
        if bFirst {
            g.AddDrawStage(func() {
            	textdraw(g.Screen, 10, 100, "TEST", color.White)
            })
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT12)
        }
    }
    /*
        S_TIMEOUT
    */
    sTIMEOUT := func( bFirst  bool ) {
        if bFirst {
            fmt.Println("S_TIMEOUT")
        }
        if !hasNextState() {
            gotoState(funcIdsCLEARALL2)
        }
    }
    /*
        S_VERSION
    */
    sVERSION := func( bFirst  bool ) {
        if bFirst {
            gdv = func() {
                textdraw(g.Screen, 0, 20, "R11281054", color.White)
            }
        }
        if !hasNextState() {
            gotoState(funcIdsOPENING)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_WAIT
    */
    sWAIT := func( bFirst  bool ) {
        if bFirst {
            timesWAIT = g.TimeNowMs() + 1000
        }
        if timesWAIT > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIds2ndRotation)
        }
    }
    /*
        S_WAIT_CLRDRAW
        600秒待つ
    */
    sWAITCLRDRAW := func( bFirst  bool ) {
        if bFirst {
            timesWAITCLRDRAW = g.TimeNowMs() + 600*1000
        }
        if ebiten.IsKeyPressed(ebiten.KeyEscape) {
            gotoState( funcIdsESC )
        } else if timesWAITCLRDRAW < g.TimeNowMs() {
            gotoState( funcIdsTIMEOUT )
        } else if g.GameOver {
            gotoState( funcIdsGAMEOVER )
        }
    }
    /*
        S_WAIT1
    */
    sWAIT1 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT1 = g.TimeNowMs() + 500
        }
        if timesWAIT1 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsTERMINAL1)
        }
    }
    /*
        S_WAIT10
    */
    sWAIT10 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT10 = g.TimeNowMs() + 500
        }
        if timesWAIT10 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsBACKTOMENU)
        }
    }
    /*
        S_WAIT11
    */
    sWAIT11 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT11 = g.TimeNowMs() + 1500
        }
        if timesWAIT11 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsLOADNIC)
        }
    }
    /*
        S_WAIT12
    */
    sWAIT12 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT12 = g.TimeNowMs() + 30*1000
        }
        x,y := ebiten.CursorPosition()
        g.TermPrint(strconv.Itoa(x) + "," + strconv.Itoa(y))
        if timesWAIT12 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsPAS001)
        }
    }
    /*
        S_WAIT13
    */
    sWAIT13 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT13 = g.TimeNowMs() + 500
        }
        if timesWAIT13 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsFADEIN1)
        }
    }
    /*
        S_WAIT2
    */
    sWAIT2 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT2 = g.TimeNowMs() + 1500
        }
        if timesWAIT2 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsCOUNT)
        }
    }
    /*
        S_WAIT3
    */
    sWAIT3 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT3 = g.TimeNowMs() + 1500
        }
        if timesWAIT3 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsBACKTOMENU)
        }
    }
    /*
        S_WAIT4
    */
    sWAIT4 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT4 = g.TimeNowMs() + 1000
        }
        if timesWAIT4 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsCLEARTERM)
        }
    }
    /*
        S_WAIT5
    */
    sWAIT5 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT5 = g.TimeNowMs() + 1500
        }
        if timesWAIT5 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIds1stRotation)
        }
    }
    /*
        S_WAIT6
    */
    sWAIT6 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT6 = g.TimeNowMs() + 1000
        }
        if timesWAIT6 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIds3rdRotation)
        }
    }
    /*
        S_WAIT7
    */
    sWAIT7 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT7 = g.TimeNowMs() + 500
        }
        if timesWAIT7 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsSETBGRED)
        }
    }
    /*
        S_WAIT8
    */
    sWAIT8 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT8 = g.TimeNowMs() + 500
        }
        if timesWAIT8 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsSETBGBLUE)
        }
    }
    /*
        S_WAIT9
    */
    sWAIT9 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT9 = g.TimeNowMs() + 500
        }
        if timesWAIT9 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsSETBGGREEN)
        }
    }


	//[STATEGO OUTPUT END]

	var funclist = [...]func(bool){

		//[STATEGO OUTPUT START] indent(8) $/^S_./->#funclist$
        //             psggConverterLib.dll converted from psgg-file:mainControl.psgg

        s1stRotation,
        s2ndRotation,
        s3rdRotation,
        sBACKTOMENU,
        sCHANGEBG,
        sCLEARALL,
        sCLEARALL1,
        sCLEARALL2,
        sCLEARTERM,
        sCOUNT,
        sDISPLAYBLK,
        sDISPLAYGOPHER,
        sDISPLAYSCORE,
        sDrawMascot,
        sEND,
        sESC,
        sFADEIN,
        sFADEIN1,
        sFADEINTEST,
        sFADEOUT,
        sFADEOUT1,
        sFADEOUT2,
        sGAMEOVER,
        sGAMEOVER1,
        sINIT,
        sINIT1,
        sISRELEASE,
        sLOADEBITEN,
        sLOADGOPHER,
        sLOADNIC,
        sLOP000,
        sLOP000LoopCheck,
        sLOP000LoopNext,
        sLOP001,
        sLOP001LoopCheck,
        sLOP001LoopNext,
        sMENU,
        sOPENING,
        sOPENING1,
        sPAS000,
        sPAS001,
        sPAS002,
        sRET000,
        sRET001,
        sRotationOverRay,
        sSBS000,
        sSBS001,
        sSETBG,
        sSETBG1,
        sSETBGBLUE,
        sSETBGGREEN,
        sSETBGRED,
        sSHOWMASCOT,
        sSHOWMASCOT1,
        sSTART,
        sTERMINAL,
        sTERMINAL1,
        sTEXT,
        sTIMEOUT,
        sVERSION,
        sWAIT,
        sWAITCLRDRAW,
        sWAIT1,
        sWAIT10,
        sWAIT11,
        sWAIT12,
        sWAIT13,
        sWAIT2,
        sWAIT3,
        sWAIT4,
        sWAIT5,
        sWAIT6,
        sWAIT7,
        sWAIT8,
        sWAIT9,


		//[STATEGO OUTPUT END]
		endofFuncList}

	nextfunc = funcIdsSTART

	update := func() bool {
		if curfunc != funcIdsEND {
			for true {
				var bFirst = false
				if nextfunc != -1 {
					curfunc = nextfunc
					nextfunc = -1
					bFirst = true
				}
				if curfunc != -1 {
					nowait = false
					funclist[curfunc](bFirst)
				}
				if nowait == false {
					break
				}
			}
			return false
		}
		return true
	}

	// for avoiding "declared but not used"
	_ = hasNextState()
	gosubState(funcIdsSTART, funcIdsEND)
	returnState()
	noWait()

	// initialize
	curfunc = -1
	nextfunc = funcIdsSTART

	// return update function     set true -- from start  set flase -- update
	return func(bFirst bool, ig *Game) bool {
		g = ig
		if bFirst {
			curfunc = -1
			nextfunc = funcIdsSTART
		}
		return update()
	}
}

/*  :::: PSGG MACRO ::::
:psgg-macro-start

commentline=// {%0}

#setids=@@@
funcId[[state>>lc]] := id
id++
<<<?state-typ/^loop$/
funcId[[state>>lc]]LoopCheck := id
id++
funcId[[state>>lc]]LoopNext := id
id++
>>>
@@@

#funclist=@@@
[[state>>lc]],
<<<?state-typ/^loop$/
[[state>>lc]]LoopCheck,
[[state>>lc]]LoopNext,
>>>
@@@

#memlist=@@@
[[members]]
@@@

@branch=@@@
<<<?"{%0}"/^brif$/
if [[brcond:{%N}]] {
    gotoState( funcId$lc:{%1}$ )
}
>>>
<<<?"{%0}"/^brifc$/
if [[brcond:{%N}]] {
    gotoState( funcId$lc:{%1}$ )
>>>
<<<?"{%0}"/^brelseif$/
} else if [[brcond:{%N}]] {
    gotoState( funcId$lc:{%1}$ )
}
>>>
<<<?"{%0}"/^brelseifc$/
} else if [[brcond:{%N}]] {
    gotoState( funcId$lc:{%1}$ )
>>>
<<<?"{%0}"/^brelse$/
} else {
    gotoState( funcId$lc:{%1}$ )
}
>>>
@@@

:psgg-macro-end

*/
