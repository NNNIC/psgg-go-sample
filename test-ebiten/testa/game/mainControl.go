﻿package game

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math"
	"math/rand"
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

    funcIds0000 := id
    id++
    funcIds0001 := id
    id++
    funcIds0002 := id
    id++
    funcIds0003 := id
    id++
    funcIds0004 := id
    id++
    funcIds0005 := id
    id++
    funcIdsEND := id
    id++
    funcIdsINIT := id
    id++
    funcIdsSTART := id
    id++


	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^S_./->#memlist$
    //             psggConverterLib.dll converted from psgg-file:mainControl.psgg

    var times0003 int64
    var times0004 int64
    var times0005 int64


	//[STATEGO OUTPUT END]

	// #
	// #  State Function
	// #
	var endofFuncList = func(bFirst bool) { // for end of function list
	}
	//[STATEGO OUTPUT START] indent(4) $/^S_./$
    //             psggConverterLib.dll converted from psgg-file:mainControl.psgg

    /*
        S_0000
    */
    s0000 := func( bFirst  bool ) {
        if bFirst {
            img, _, err := image.Decode(bytes.NewReader(g.Gophers_jpg()))
            if err != nil {
            	log.Fatal(err)
            }
            g.GophersImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
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
            gotoState(funcIds0003)
        }
    }
    /*
        S_0001
    */
    s0001 := func( bFirst  bool ) {
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
            gotoState(funcIds0004)
        }
    }
    /*
        S_0002
    */
    s0002 := func( bFirst  bool ) {
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
            gotoState(funcIdsEND)
        }
    }
    /*
        S_0003
    */
    s0003 := func( bFirst  bool ) {
        if bFirst {
            times0003 = g.TimeNowMs() + 1000
        }
        if times0003 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIds0001)
        }
    }
    /*
        S_0004
    */
    s0004 := func( bFirst  bool ) {
        if bFirst {
            times0004 = g.TimeNowMs() + 1000
        }
        if times0004 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIds0002)
        }
    }
    /*
        S_0005
    */
    s0005 := func( bFirst  bool ) {
        if bFirst {
            times0005 = g.TimeNowMs() + 10000
        }
        if times0005 > g.TimeNowMs() {
             return
        }
        if !hasNextState() {
            gotoState(funcIds0000)
        }
    }
    /*
        S_END
    */
    sEND := func ( bFirst  bool ) {
         // end of state machine
    }
    /*
        S_INIT
    */
    sINIT := func( bFirst  bool ) {
        if bFirst {
            fmt.Println("START")
        }
        if !hasNextState() {
            gotoState(funcIds0005)
        }
    }
    /*
        S_START
    */
    sSTART := func( bFirst  bool ) {
        gotoState(funcIdsINIT)
    }


	//[STATEGO OUTPUT END]

	var funclist = [...]func(bool){

		//[STATEGO OUTPUT START] indent(8) $/^S_./->#funclist$
        //             psggConverterLib.dll converted from psgg-file:mainControl.psgg

        s0000,
        s0001,
        s0002,
        s0003,
        s0004,
        s0005,
        sEND,
        sINIT,
        sSTART,


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
