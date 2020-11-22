package game

import (
	"bytes"
	"image"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type NicData struct {
	g      *Game
	PosX   float64
	PosY   float64
	Killed bool
}

func nicControl(d *NicData) func(bool, *Game) bool {
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
    //             psggConverterLib.dll converted from psgg-file:nicControl.psgg

    funcIdsDrawNic := id
    id++
    funcIdsDrawUpdate := id
    id++
    funcIdsEND := id
    id++
    funcIdsLOADNIC := id
    id++
    funcIdsSTART := id
    id++
    funcIdsUpdate := id
    id++


	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^S_./->#memlist$
    //             psggConverterLib.dll converted from psgg-file:nicControl.psgg

    var drawfunc func()
    var angle    float64


	//[STATEGO OUTPUT END]

	// USER FUNCTION

	calcUpdate := func() {
		gx := g.GopherData0.PosX
		gy := g.GopherData0.PosY
		dx := gx - d.PosX
		dy := gy - d.PosY
		angle = math.Atan2(dy, dx) * (180.0 / math.Pi)
		angle = angleNormal(angle + 180)
	}

	//
	// #
	// #  State Function
	// #
	var endofFuncList = func(bFirst bool) { // for end of function list
	}
	//[STATEGO OUTPUT START] indent(4) $/^S_./$
    //             psggConverterLib.dll converted from psgg-file:nicControl.psgg

    /*
        S_DrawNic
    */
    sDrawNic := func( bFirst  bool ) {
        if bFirst {
            drawfunc = func() {
                g.DrawImage(g.MascotImage, float64(d.PosX-16),float64(d.PosY-16),angle,1)
            }
            g.AddDrawStage(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsUpdate)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_DrawUpdate
    */
    sDrawUpdate := func( bFirst  bool ) {
        calcUpdate()
        g.AddDrawStage(drawfunc)
        if ebiten.IsKeyPressed(ebiten.KeyEscape) {
            gotoState( funcIdsEND )
        } else if d.Killed {
            gotoState( funcIdsEND )
        } else {
            gotoState( funcIdsUpdate )
        }
    }
    /*
        S_END
    */
    sEND := func ( bFirst  bool ) {
         // end of state machine
    }
    /*
        S_LOAD_NIC
    */
    sLOADNIC := func( bFirst  bool ) {
        if bFirst {
            img, _, err := image.Decode(bytes.NewReader(g.Mascot32_png()))
            if err != nil {
                log.Fatal(err)
            }
            g.MascotImage = ebiten.NewImageFromImage(img)
            drawfunc := func() {
            	g.DrawImage(g.MascotImage,8,8,0,1)
            }
            g.AddDrawStage(drawfunc)
        }
        if !hasNextState() {
            gotoState(funcIdsDrawNic)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_START
    */
    sSTART := func( bFirst  bool ) {
        gotoState(funcIdsLOADNIC)
    }
    /*
        S_Update
    */
    sUpdate := func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsDrawUpdate)
        }
        if hasNextState() {
            noWait()
        }
    }


	//[STATEGO OUTPUT END]

	var funclist = [...]func(bool){

		//[STATEGO OUTPUT START] indent(8) $/^S_./->#funclist$
        //             psggConverterLib.dll converted from psgg-file:nicControl.psgg

        sDrawNic,
        sDrawUpdate,
        sEND,
        sLOADNIC,
        sSTART,
        sUpdate,


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
		d.g = ig
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
