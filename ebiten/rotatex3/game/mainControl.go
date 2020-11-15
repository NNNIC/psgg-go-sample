package game

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Gamemain ... invoke game
func Gamemain() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Rotate (StateGo Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

// MainControl ... main flow
func MainControl() func(bool, *Game) bool {
	var g *Game

	rand.Seed(time.Now().Unix())

	curfunc := -1
	nextfunc := -1

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
    funcIdsCLEARTERM := id
    id++
    funcIdsCOUNT := id
    id++
    funcIdsEND := id
    id++
    funcIdsLOADSHOW := id
    id++
    funcIdsMENU := id
    id++
    funcIdsRotationOverRay := id
    id++
    funcIdsSETBGBLUE := id
    id++
    funcIdsSETBGGREEN := id
    id++
    funcIdsSETBGRED := id
    id++
    funcIdsSHOWMASCOT := id
    id++
    funcIdsSTART := id
    id++
    funcIdsTERMINAL := id
    id++
    funcIdsTERMINAL1 := id
    id++
    funcIdsWAIT := id
    id++
    funcIdsWAIT1 := id
    id++
    funcIdsWAIT10 := id
    id++
    funcIdsWAIT11 := id
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
    var timesWAIT int64
    var timesWAIT1 int64
    var timesWAIT10 int64
    var timesWAIT11 int64
    var timesWAIT2 int64
    var timesWAIT3 int64
    var timesWAIT4 int64
    var timesWAIT5 int64
    var timesWAIT6 int64
    var timesWAIT7 int64
    var timesWAIT8 int64
    var timesWAIT9 int64


	//[STATEGO OUTPUT END]

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
                g.GophersImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
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
            gotoState(funcIdsMENU)
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
        S_END
    */
    sEND := func ( bFirst  bool ) {
         // end of state machine
    }
    /*
        S_LOAD_SHOW
    */
    sLOADSHOW := func( bFirst  bool ) {
        if bFirst {
            if g.MascotImage == nil {
                img, _, err := image.Decode(bytes.NewReader(g.Mascot_png()))
                if err != nil {
            	log.Fatal(err)
                }
                g.MascotImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
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
            gotoState(funcIdsBACKTOMENU)
        }
    }
    /*
        S_MENU
    */
    sMENU := func( bFirst  bool ) {
        if bFirst {
            g.TermPrint("==== TEST =====");
            g.TermPrint("Push 1 ... Termainal Test");
            g.TermPrint("Push 2 ... Background Color Change Test");
            g.TermPrint("Push 3 ... Ebiten Rotaion Overlay Demo");
            g.TermPrint("Push 4 ... StateGo Mascot Demo");
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
        } else if ebiten.IsKeyPressed(ebiten.KeyC) {
            gotoState( funcIdsCLEARALL )
        }
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
        S_START
    */
    sSTART := func( bFirst  bool ) {
        gotoState(funcIdsMENU)
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
            gotoState(funcIdsLOADSHOW)
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
        sCLEARTERM,
        sCOUNT,
        sEND,
        sLOADSHOW,
        sMENU,
        sRotationOverRay,
        sSETBGBLUE,
        sSETBGGREEN,
        sSETBGRED,
        sSHOWMASCOT,
        sSTART,
        sTERMINAL,
        sTERMINAL1,
        sWAIT,
        sWAIT1,
        sWAIT10,
        sWAIT11,
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
			var bFirst = false
			if nextfunc != -1 {
				curfunc = nextfunc
				nextfunc = -1
				bFirst = true
			}
			if curfunc != -1 {
				funclist[curfunc](bFirst)
			}
			return false
		}
		return true
	}

	// for avoiding "declared but not used"
	_ = hasNextState()
	gosubState(funcIdsSTART, funcIdsEND)
	returnState()

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
    gotoState( funcId{%1} )
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
