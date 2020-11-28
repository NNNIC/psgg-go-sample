package game

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gLimitTop   = 40
	gLimitBot   = screenHeight - 40
	gLimitLeft  = 40
	gLimitRight = screenWidth - 40
)

type GopherData struct {
	g *Game

	PosX float64
	PosY float64

	DiffX float64
	DiffY float64

	Killed bool
}

// CheckKill ... 1st killed blk  2nd All Killed
func (d *GopherData) CheckKill() (bool, bool) {
	killedBlk := false
	killedCount := 0
	for _, nd := range d.g.BlkDataList {
		if nd.Killed {
			killedCount++
		} else {
			dx := float64(d.PosX - nd.PosX)
			dy := float64(d.PosY - nd.PosY)

			len := math.Sqrt(dx*dx + dy*dy)
			if len <= 16 {
				nd.Killed = true
				killedBlk = true
				break
			}
		}
	}
	if killedBlk {
		return true, false
	}
	if killedCount == len(d.g.BlkDataList) {
		return false, true
	}
	return false, false
}
func (d *GopherData) Reflect() {
	d.DiffX = -d.DiffX
	d.DiffY = -d.DiffY
}
func (d *GopherData) Speedup(add float64) {
	x, y := vectorAdd(d.DiffX, d.DiffY, add, 1.0, 10.0)
	d.DiffX = x
	d.DiffY = y
	//fmt.Println(d.DiffX, ",", d.DiffY)
}
func (d *GopherData) CalcDirClicked() {
	mx, my := ebiten.CursorPosition()
	d.calcDir(mx, my)
}
func (d *GopherData) calcDir(mx, my int) {
	dx := float64(mx) - d.PosX
	dy := float64(my) - d.PosY
	nx, ny := vectorNormalize(dx, dy)
	curlen := vectorLen(d.DiffX, d.DiffY)
	d.DiffX = nx * curlen
	d.DiffY = ny * curlen
}
func (d *GopherData) CalcDirTouched() {
	mx, my := ebiten.TouchPosition(ebiten.TouchIDs()[0])
	d.calcDir(mx, my)
}
func (d *GopherData) MoveDiff() bool {
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

func gopherControl(d *GopherData, id int) func(bool, *Game) bool {
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
	sid := 0
	//[STATEGO OUTPUT START] indent(4) $/^S_./->#setids$
    //             psggConverterLib.dll converted from psgg-file:gopherControl.psgg

    funcIdsBACKTOMOVEDIFF := sid
    sid++
    funcIdsDrawGopher := sid
    sid++
    funcIdsDrawUpdate := sid
    sid++
    funcIdsDrawUpdate1 := sid
    sid++
    funcIdsEND := sid
    sid++
    funcIdsLOADGOPHER := sid
    sid++
    funcIdsMouseClick := sid
    sid++
    funcIdsMouseClick1 := sid
    sid++
    funcIdsMOVEDIFF := sid
    sid++
    funcIdsPERFECT := sid
    sid++
    funcIdsREFLECT := sid
    sid++
    funcIdsSTART := sid
    sid++
    funcIdsTick := sid
    sid++
    funcIdsTick1 := sid
    sid++
    funcIdsTouched := sid
    sid++
    funcIdsWAIT1 := sid
    sid++
    funcIdsWAIT2 := sid
    sid++


	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^S_./->#memlist$
    //             psggConverterLib.dll converted from psgg-file:gopherControl.psgg

    var drawfunc func()
    var timesWAIT1 int64


	//[STATEGO OUTPUT END]
	// [STATEGO OUTPUT START] indent(4) $/^E_./$
    //             psggConverterLib.dll converted from psgg-file:gopherControl.psgg



	//[STATEGO OUTPUT END]

	// #
	// #  State Function
	// #
	var endofFuncList = func(bFirst bool) { // for end of function list
	}
	//[STATEGO OUTPUT START] indent(4) $/^S_./$
    //             psggConverterLib.dll converted from psgg-file:gopherControl.psgg

    /*
        S_BACKTO_MOVEDIFF
    */
    sBACKTOMOVEDIFF := func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsMOVEDIFF)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_DrawGopher
    */
    sDrawGopher := func( bFirst  bool ) {
        if bFirst {
            drawfunc = func() {
                g.DrawImage(g.GophersImage3, float64(d.PosX-16),float64(d.PosY-16),0,1)
            }
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT1)
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
        } else {
            gotoState( funcIdsTick )
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_DrawUpdate1
    */
    sDrawUpdate1 := func( bFirst  bool ) {
        d.MoveDiff()
        g.AddDrawStage(drawfunc)
        if ebiten.IsKeyPressed(ebiten.KeyEscape) {
            gotoState( funcIdsEND )
        } else {
            gotoState( funcIdsTick1 )
        }
        if hasNextState() {
            noWait()
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
            if (g.GophersImage3 == nil) {
                img, _, err := image.Decode(bytes.NewReader(g.Gophers32_png()))
                if err != nil {
                    log.Fatal(err)
                }
                g.GophersImage3 = ebiten.NewImageFromImage(img)
            }
        }
        if !hasNextState() {
            gotoState(funcIdsDrawGopher)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_MouseClick
    */
    sMouseClick := func( bFirst  bool ) {
        bMouse := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
        bTouch := ebiten.TouchIDs() != nil
        b := bMouse || bTouch
        if b && g.ScoreClickCount != g.Count {
        	g.ScoreClickCount = g.Count
        	g.ScoreData0.Add(-1000)
        }
        if !b {
            gotoState( funcIdsDrawUpdate )
        } else if bMouse {
            gotoState( funcIdsMouseClick1 )
        } else {
            gotoState( funcIdsTouched )
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_MouseClick1
    */
    sMouseClick1 := func( bFirst  bool ) {
        if bFirst {
            d.CalcDirClicked()
        }
        if !hasNextState() {
            gotoState(funcIdsDrawUpdate)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_MOVEDIFF
    */
    sMOVEDIFF := func( bFirst  bool ) {
        if bFirst {
            hitwall := d.MoveDiff()
            if hitwall {
                d.Speedup(-0.5)
                g.ScoreMul = 1
            }
        }
        bKillBlk, bKillAll :=d.CheckKill()
        if bKillAll {
            gotoState( funcIdsPERFECT )
        } else if bKillBlk {
            gotoState( funcIdsREFLECT )
        } else {
            gotoState( funcIdsMouseClick )
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_PERFECT
    */
    sPERFECT := func( bFirst  bool ) {
        if bFirst {
            g.GameOver = true
        }
        if !hasNextState() {
            gotoState(funcIdsDrawUpdate1)
        }
    }
    /*
        S_REFLECT
    */
    sREFLECT := func( bFirst  bool ) {
        if bFirst {
            d.Speedup(0.2)
            d.Reflect()
            g.ScoreData0.Add(1000 * g.ScoreMul)
            g.ScoreMul++
            g.ScoreBase += 500
        }
        if !hasNextState() {
            gotoState(funcIdsDrawUpdate)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_START
    */
    sSTART := func( bFirst  bool ) {
        gotoState(funcIdsLOADGOPHER)
    }
    /*
        S_Tick
    */
    sTick := func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsBACKTOMOVEDIFF)
        }
    }
    /*
        S_Tick1
    */
    sTick1 := func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsDrawUpdate1)
        }
    }
    /*
        S_Touched
    */
    sTouched := func( bFirst  bool ) {
        if bFirst {
            d.CalcDirTouched()
        }
        if !hasNextState() {
            gotoState(funcIdsDrawUpdate)
        }
        if hasNextState() {
            noWait()
        }
    }
    /*
        S_WAIT1
    */
    sWAIT1 := func( bFirst  bool ) {
        if bFirst {
            timesWAIT1 = g.TimeNowMs() + 3 * 1000
        }
        g.AddDrawStage(drawfunc)
        dt := func() {
            red := color.RGBA{255, 0, 0, 255}
            textdrawBigWFrame(g.Screen,  90,   240  , "CLICK ON SCREEN", color.White, red)
            textdrawBigWFrame(g.Screen, 175,   310  , "TO MOVE", color.White, red)
        }
        g.AddDrawStage(dt)
        b := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
        if !b {
             return
        }
        if !hasNextState() {
            gotoState(funcIdsWAIT2)
        }
    }
    /*
        S_WAIT2
    */
    sWAIT2 := func( bFirst  bool ) {
        if bFirst {
            d.DiffX = -1.0
            d.DiffY = -1.0
        }
        if !hasNextState() {
            gotoState(funcIdsMOVEDIFF)
        }
        if hasNextState() {
            noWait()
        }
    }


	//[STATEGO OUTPUT END]

	var funclist = [...]func(bool){

		//[STATEGO OUTPUT START] indent(8) $/^S_./->#funclist$
        //             psggConverterLib.dll converted from psgg-file:gopherControl.psgg

        sBACKTOMOVEDIFF,
        sDrawGopher,
        sDrawUpdate,
        sDrawUpdate1,
        sEND,
        sLOADGOPHER,
        sMouseClick,
        sMouseClick1,
        sMOVEDIFF,
        sPERFECT,
        sREFLECT,
        sSTART,
        sTick,
        sTick1,
        sTouched,
        sWAIT1,
        sWAIT2,


		//[STATEGO OUTPUT END]
		endofFuncList}
	/*
																										var dbgfunclist = [...]string{
																											//[STATEGO OUTPUT START] indent(8) $/^S_./->#dbgfunclist$
        //             psggConverterLib.dll converted from psgg-file:gopherControl.psgg

        "S_BACKTO_MOVEDIFF",
        "S_DrawGopher",
        "S_DrawUpdate",
        "S_DrawUpdate1",
        "S_END",
        "S_LOAD_GOPHER",
        "S_MouseClick",
        "S_MouseClick1",
        "S_MOVEDIFF",
        "S_PERFECT",
        "S_REFLECT",
        "S_START",
        "S_Tick",
        "S_Tick1",
        "S_Touched",
        "S_WAIT1",
        "S_WAIT2",


																											//[STATEGO OUTPUT END]
																											"none"}
	*/
	nextfunc = funcIdsSTART

	update := func() bool {
		if curfunc != funcIdsEND {
			for true {
				var bFirst = false
				if nextfunc != -1 {
					curfunc = nextfunc
					nextfunc = -1
					bFirst = true
					//fmt.Println(dbgfunclist[curfunc])
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
funcId[[state>>lc]] := sid
sid++
<<<?state-typ/^loop$/
funcId[[state>>lc]]LoopCheck := sid
sid++
funcId[[state>>lc]]LoopNext := sid
sid++
>>>
@@@

#funclist=@@@
[[state>>lc]],
<<<?state-typ/^loop$/
[[state>>lc]]LoopCheck,
[[state>>lc]]LoopNext,
>>>
@@@

#dbgfunclist=@@@
"[[state]]",
<<<?state-typ/^loop$/
"[[state]]_LoopCheck",
"[[state]]_LoopNext",
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
