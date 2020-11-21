package game

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type GopherData struct {
	PosX int
	PosY int
}

func gopherControl(d *GopherData) func(bool, *Game) bool {
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
    //             psggConverterLib.dll converted from psgg-file:gopherControl.psgg

    funcIdsDrawGopher := id
    id++
    funcIdsDrawUpdate := id
    id++
    funcIdsEND := id
    id++
    funcIdsLOADGOPHER := id
    id++
    funcIdsSTART := id
    id++


	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^S_./->#memlist$
    //             psggConverterLib.dll converted from psgg-file:gopherControl.psgg

    var drawfunc func()

	//[STATEGO OUTPUT END]

	// #
	// #  State Function
	// #
	var endofFuncList = func(bFirst bool) { // for end of function list
	}
	//[STATEGO OUTPUT START] indent(4) $/^S_./$
    //             psggConverterLib.dll converted from psgg-file:gopherControl.psgg

    /*
        S_DrawGopher
    */
    sDrawGopher := func( bFirst  bool ) {
        if bFirst {
            drawfunc = func() {
                g.DrawImage(g.GophersImage2, float64(d.PosX-16),float64(d.PosY-16),0,1)
            }
        }
        if !hasNextState() {
            gotoState(funcIdsDrawUpdate)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_DrawUpdate
    */
    sDrawUpdate := func( bFirst  bool ) {
        g.AddDrawStage(drawfunc)
        if ebiten.IsKeyPressed(ebiten.KeyEscape) {
            gotoState( funcIdsEND )
        }
    }
    /*
        S_END
    */
    sEND := func ( bFirst  bool ) {
         // end of state machine
    }
    /*
        S_LOAD_GOPHER
    */
    sLOADGOPHER := func( bFirst  bool ) {
        if bFirst {
            fmt.Println("S_LOAD_GOPHER")
            img, _, err := image.Decode(bytes.NewReader(g.Gophers32_png()))
            if err != nil {
                log.Fatal(err)
            }
            g.GophersImage2, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
        }
        if !hasNextState() {
            gotoState(funcIdsDrawGopher)
        }
    }
    /*
        S_START
    */
    sSTART := func( bFirst  bool ) {
        gotoState(funcIdsLOADGOPHER)
    }


	//[STATEGO OUTPUT END]

	var funclist = [...]func(bool){

		//[STATEGO OUTPUT START] indent(8) $/^S_./->#funclist$
        //             psggConverterLib.dll converted from psgg-file:gopherControl.psgg

        sDrawGopher,
        sDrawUpdate,
        sEND,
        sLOADGOPHER,
        sSTART,


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
