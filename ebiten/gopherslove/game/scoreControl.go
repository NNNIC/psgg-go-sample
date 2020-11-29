package game

import (
	"image/color"
	"strconv"
)

const (
	pointListMax = 5
)

type ScoreData struct {
	g         *Game
	PointList []int
	Score     int
	Msg1      string
	Msg2      string
}

func (d *ScoreData) Add(point int) {
	if len(d.PointList) == 0 {
		d.PointList = make([]int, pointListMax)
		for i := 0; i < pointListMax; i++ {
			d.PointList[i] = -1
		}
	}
	if d.PointList[pointListMax-1] != -1 {
		d.Score += d.PointList[0]
		for i := 0; i < pointListMax-1; i++ {
			d.PointList[i] = d.PointList[i+1]
		}
		d.PointList[pointListMax-1] = -1
	}
	for i := 0; i < pointListMax; i++ {
		if d.PointList[i] == -1 {
			d.PointList[i] = point
			if i+1 < pointListMax {
				d.PointList[i+1] = -1
			}
			break
		}
	}
}
func (d *ScoreData) ToString() (string, string) {
	s1 := "Score:" + strconv.Itoa(d.Score)
	s2 := ""
	for i := 0; i < pointListMax; i++ {
		v := d.PointList[i]
		if v == -1 {
			break
		}
		s2 += "+" + strconv.Itoa(v)
	}
	return s1, s2
}
func (d *ScoreData) Update() {
	if d.g.Count%10 == 0 {
		if d.PointList[0] != -1 {
			d.Score += d.PointList[0]
		}
		for i := 0; i < pointListMax-1; i++ {
			d.PointList[i] = d.PointList[i+1]
		}
		d.PointList[pointListMax-1] = -1
	}
}

func scoreControl(d *ScoreData) func(bool, *Game) bool {
	var g *Game

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
	//             psggConverterLib.dll converted from psgg-file:scoreControl.psgg

	funcIds0000 := sid
	sid++
	funcIdsEND := sid
	sid++
	funcIdsINIT := sid
	sid++
	funcIdsSTART := sid
	sid++

	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^S_./->#memlist$
	//             psggConverterLib.dll converted from psgg-file:scoreControl.psgg

	//[STATEGO OUTPUT END]
	// [STATEGO OUTPUT START] indent(4) $/^E_./$
	//             psggConverterLib.dll converted from psgg-file:scoreControl.psgg

	/*
	   E_0000
	*/
	df := func() {
		textdraw(g.Screen, 10, 20, d.Msg1, color.White)
		textdraw(g.Screen, 50, 50, d.Msg2, color.White)
	}

	//[STATEGO OUTPUT END]

	// #
	// #  State Function
	// #
	var endofFuncList = func(bFirst bool) { // for end of function list
	}
	//[STATEGO OUTPUT START] indent(4) $/^S_./$
	//             psggConverterLib.dll converted from psgg-file:scoreControl.psgg

	/*
	   S_0000
	*/
	s0000 := func(bFirst bool) {
		d.Update()
		d.Msg1, d.Msg2 = d.ToString()
		g.AddDrawStage(df)
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
			d.Score = 0
			d.Add(0)
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
		//             psggConverterLib.dll converted from psgg-file:scoreControl.psgg

		s0000,
		sEND,
		sINIT,
		sSTART,

		//[STATEGO OUTPUT END]
		endofFuncList}
	/*
																														var dbgfunclist = [...]string{
																															//[STATEGO OUTPUT START] indent(8) $/^S_./->#dbgfunclist$
		        //             psggConverterLib.dll converted from psgg-file:scoreControl.psgg

		        "S_0000",
		        "S_END",
		        "S_INIT",
		        "S_START",


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
