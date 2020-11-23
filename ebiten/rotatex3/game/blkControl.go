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

type BlkData struct {
	g      *Game
	PosX   float64
	PosY   float64
	Killed bool
	Anger  bool

	DiffX float64
	DiffY float64
}

func (d *BlkData) MoveDiff() bool {
	x := d.PosX
	y := d.PosY
	x = x + d.DiffX
	y = y + d.DiffY

	vT := 1
	vB := 2
	vR := 4
	vL := 8
	cond := 0
	if y <= gLimitTop {
		cond += vT
	} else if y >= gLimitBot {
		cond += vB
	}
	if x <= gLimitLeft {
		cond += vL
	} else if x >= gLimitRight {
		cond += vR
	}

	if cond == vT {
		d.DiffY = math.Abs(d.DiffY)
		d.PosY = gLimitTop + 1
	} else if cond == vB {
		d.DiffY = -math.Abs(d.DiffY)
		d.PosY = gLimitBot - 1
	} else if cond == vL {
		d.DiffX = math.Abs(d.DiffX)
		d.PosX = gLimitLeft + 1
	} else if cond == vR {
		d.DiffX = -math.Abs(d.DiffX)
		d.PosX = gLimitRight - 1
	} else if cond == vT+vL {
		d.DiffX = math.Abs(d.DiffX)
		d.DiffY = math.Abs(d.DiffY)
		d.PosX = gLimitLeft + 1
		d.PosY = gLimitTop + 1
	} else if cond == vT+vR {
		d.DiffX = -math.Abs(d.DiffX)
		d.DiffY = math.Abs(d.DiffY)
		d.PosX = gLimitRight - 1
		d.PosY = gLimitTop + 1
	} else if cond == vB+vL {
		d.DiffX = math.Abs(d.DiffX)
		d.DiffY = -math.Abs(d.DiffY)
		d.PosX = gLimitLeft + 1
		d.PosY = gLimitBot - 1
	} else if cond == vB+vR {
		d.DiffX = -math.Abs(d.DiffX)
		d.DiffY = -math.Abs(d.DiffY)
		d.PosX = gLimitRight - 1
		d.PosY = gLimitBot - 1
	} else {
		d.PosX = x
		d.PosY = y
	}

	//fmt.Println(cond)
	return cond > 0
}

func blkControl(d *BlkData) func(bool, *Game) bool {
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
    //             psggConverterLib.dll converted from psgg-file:blkControl.psgg

    funcIdsBACKTODRAWANGER := id
    id++
    funcIdsDrawAnger := id
    id++
    funcIdsDrawAnger1 := id
    id++
    funcIdsDrawBlk := id
    id++
    funcIdsDrawUpdate := id
    id++
    funcIdsEND := id
    id++
    funcIdsLOADBLK := id
    id++
    funcIdsSTART := id
    id++
    funcIdsTICK := id
    id++
    funcIdsUpdate := id
    id++


	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^S_./->#memlist$
    //             psggConverterLib.dll converted from psgg-file:blkControl.psgg

    var angle    float64

	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^E_./$
    //             psggConverterLib.dll converted from psgg-file:blkControl.psgg

    /*
        E_DrawNic
    */
    var drawfunc func()
    var drawfuncA func()


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
    //             psggConverterLib.dll converted from psgg-file:blkControl.psgg

    /*
        S_BACKTO_DRAWANGER
    */
    sBACKTODRAWANGER := func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsTICK)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_DrawAnger
    */
    sDrawAnger := func( bFirst  bool ) {
        if bFirst {
            drawfuncA = func() {
                        g.DrawImage(g.MascotAImage, float64(d.PosX-16),float64(d.PosY-16),angle,1)
            }
            d.DiffX = 1
            d.DiffY = 1
        }
        if !hasNextState() {
            gotoState(funcIdsTICK)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_DrawAnger1
    */
    sDrawAnger1 := func( bFirst  bool ) {
        if bFirst {
            d.MoveDiff()
            g.AddDrawStage(drawfuncA)
        }
        if !hasNextState() {
            gotoState(funcIdsBACKTODRAWANGER)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_DrawBlk
    */
    sDrawBlk := func( bFirst  bool ) {
        if bFirst {
            drawfunc = func() {
                g.DrawImage(g.EbitenImage, float64(d.PosX-16),float64(d.PosY-16),angle,1)
            }
            g.AddDrawStage(drawfunc)
            d.DiffX = 0.7
            d.DiffY = -0.7
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
        d.MoveDiff()
        calcUpdate()
        g.AddDrawStage(drawfunc)
        if ebiten.IsKeyPressed(ebiten.KeyEscape) {
            gotoState( funcIdsEND )
        } else if d.Killed {
            gotoState( funcIdsEND )
        } else if d.Anger {
            gotoState( funcIdsDrawAnger )
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
        S_LOAD_BLK
    */
    sLOADBLK := func( bFirst  bool ) {
        if bFirst {
            img, _, err := image.Decode(bytes.NewReader(g.Ebiten32_png()))
            if err != nil {
                log.Fatal(err)
            }
            g.EbitenImage = ebiten.NewImageFromImage(img)
            img, _, err = image.Decode(bytes.NewReader(g.MascotA32_png()))
            if err != nil {
                log.Fatal(err)
            }
            g.MascotAImage = ebiten.NewImageFromImage(img)
            drawfunc = func() {
            	g.DrawImage(g.EbitenImage,8,8,0,1)
            }
            g.AddDrawStage(drawfunc)
            drawfuncA = func() {
            	g.DrawImage(g.MascotAImage,8,8,0,1)
            }
        }
        if !hasNextState() {
            gotoState(funcIdsDrawBlk)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_START
    */
    sSTART := func( bFirst  bool ) {
        gotoState(funcIdsLOADBLK)
    }
    /*
        S_TICK
    */
    sTICK := func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsDrawAnger1)
        }
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
        //             psggConverterLib.dll converted from psgg-file:blkControl.psgg

        sBACKTODRAWANGER,
        sDrawAnger,
        sDrawAnger1,
        sDrawBlk,
        sDrawUpdate,
        sEND,
        sLOADBLK,
        sSTART,
        sTICK,
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
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			d.Anger = true
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
