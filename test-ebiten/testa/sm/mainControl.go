package sm

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math"
	"math/rand"
	"time"

	"game"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

// Gamemain ... invoke game
func Gamemain() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Rotate (StateGo Ebiten Demo)")
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}

// MainControl ... main flow
func MainControl() func(bool) bool {
	var g *game.Game

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
	funcIdsEND := id
	id++
	funcIdsINIT := id
	id++
	funcIdsSTART := id
	id++

	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^S_./->#memlist$
	//             psggConverterLib.dll converted from psgg-file:mainControl.psgg

	//[STATEGO OUTPUT END]

	// #
	// #  State Function
	// #
	var endofFuncList = func(bFirst bool) { // for end of function list
	}
	//[STATEGO OUTPUT START] indent(4) $/./$
	//             psggConverterLib.dll converted from psgg-file:mainControl.psgg

	/*
	   S_0000
	*/
	s0000 := func(bFirst bool) {
		if bFirst {
			img, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
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
				op.GeoM.Rotate(float64(g.Count%360) * 2 * math.Pi / 360)
				// Move the image to the screen's center.
				op.GeoM.Translate(float64(g.ScreenWidth/2), float64(g.ScreenHeight/2))
				g.Screen.DrawImage(g.gophersImage, op)
			}
			g.AppendDrawStage(drawfunc)
		}
		if !hasNextState() {
			gotoState(funcIdsEND)
		}
	}
	/*
	   S_END
	*/
	sEND := func(bFirst bool) {
		// end of state machine
	}
	/*
	   S_INIT
	*/
	sINIT := func(bFirst bool) {
		if bFirst {
			fmt.Println("!!")
		}
		if !hasNextState() {
			gotoState(funcIds0000)
		}
	}
	/*
	   S_START
	*/
	sSTART := func(bFirst bool) {
		gotoState(funcIdsINIT)
	}

	//[STATEGO OUTPUT END]

	var funclist = [...]func(bool){

		//[STATEGO OUTPUT START] indent(8) $/^S_./->#funclist$
		//             psggConverterLib.dll converted from psgg-file:mainControl.psgg

		s0000,
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

	return func(bFirst bool, ig *game.Game) bool {
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
