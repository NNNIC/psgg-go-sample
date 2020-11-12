﻿package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	testControl()
}

func testControl() {

	rand.Seed(time.Now().Unix())

	var curfunc = -1
	var nextfunc = -1

	var gotoState = func(st int) {
		nextfunc = st
	}
	var hasNextState = func() bool {
		return nextfunc != -1
	}

	var callstack = [...]int{0, 0, 0, 0, 0}
	var callstacklevel = 0
	var gosubState = func(sub, next int) {
		if callstacklevel >= len(callstack)-1 {
			panic("call stack overflow")
			//return
		}
		callstack[callstacklevel] = next
		callstacklevel++
		gotoState(sub)
	}
	var returnState = func() {
		if callstacklevel <= 0 {
			panic("call stack underflow")
			//return
		}
		callstacklevel--
		var funcid = callstack[callstacklevel]
		gotoState(funcid)
	}
	// #
	// # Define function ID for state
	// #
	var id = 0
	//[STATEGO OUTPUT START] indent(4) $/^S_./->#setids$
    //             psggConverterLib.dll converted from psgg-file:testControl.psgg

    var funcIds0001 = id
    id++
    var funcIds0002 = id
    id++
    var funcIds0003 = id
    id++
    var funcIds0004 = id
    id++
    var funcIds0005 = id
    id++
    var funcIds0006 = id
    id++
    var funcIds0007 = id
    id++
    var funcIdsEND = id
    id++
    var funcIdsGSB000 = id
    id++
    var funcIdsHELLOWORLD = id
    id++
    var funcIdsLOP000 = id
    id++
    var funcIdsLOP000LoopCheck = id
    id++
    var funcIdsLOP000LoopNext = id
    id++
    var funcIdsNEWERA = id
    id++
    var funcIdsPAS000 = id
    id++
    var funcIdsPAS001 = id
    id++
    var funcIdsRET000 = id
    id++
    var funcIdsRET001 = id
    id++
    var funcIdsSBS000 = id
    id++
    var funcIdsSBS001 = id
    id++
    var funcIdsSTART = id
    id++


	//[STATEGO OUTPUT END]

	// [STATEGO OUTPUT START] indent(4) $/^S_./->#memlist$
    //             psggConverterLib.dll converted from psgg-file:testControl.psgg

    var counter = 0

	//[STATEGO OUTPUT END]

	// #
	// #  State Function
	// #
	var endofFuncList = func(bFirst bool) { // for end of function list
	}
	//[STATEGO OUTPUT START] indent(4) $/./$
    //             psggConverterLib.dll converted from psgg-file:testControl.psgg

    /*
        S_0001
    */
    var s0001 = func( bFirst  bool ) {
        var n = rand.Intn(5)
        if n==0 {
            gotoState( funcIds0004 )
        } else if n==1 {
            gotoState( funcIds0002 )
        } else if n==2 {
            gotoState( funcIds0003 )
        } else {
            gotoState( funcIds0007 )
        }
    }
    /*
        S_0002
    */
    var s0002 = func( bFirst  bool ) {
        if bFirst {
            fmt.Println("1")
        }
        if !hasNextState() {
            gotoState(funcIdsPAS000)
        }
    }
    /*
        S_0003
    */
    var s0003 = func( bFirst  bool ) {
        if bFirst {
            fmt.Println("2")
        }
        if !hasNextState() {
            gotoState(funcIdsPAS000)
        }
    }
    /*
        S_0004
    */
    var s0004 = func( bFirst  bool ) {
        if bFirst {
            fmt.Println("0")
        }
        if !hasNextState() {
            gotoState(funcIdsPAS000)
        }
    }
    /*
        S_0005
    */
    var s0005 = func( bFirst  bool ) {
        if bFirst {
            fmt.Println("In a subroutine!")
        }
        if !hasNextState() {
            gotoState(funcIdsRET000)
        }
    }
    /*
        S_0006
    */
    var s0006 = func( bFirst  bool ) {
        if bFirst {
            fmt.Println("In a loop subroutine! no.",counter)
        }
        if !hasNextState() {
            gotoState(funcIdsRET001)
        }
    }
    /*
        S_0007
    */
    var s0007 = func( bFirst  bool ) {
        if bFirst {
            fmt.Println("Other")
        }
        if !hasNextState() {
            gotoState(funcIdsPAS000)
        }
    }
    /*
        S_END
    */
    var sEND = func ( bFirst  bool ) {
         // end of state machine
    }
    /*
        S_GSB000
    */
    var sGSB000 = func ( bFirst bool ) {
        gosubState(funcIdsSBS000, funcIdsPAS001)
    }
    /*
        S_HELLOWORLD
    */
    var sHELLOWORLD = func( bFirst  bool ) {
        if bFirst {
            fmt.Println("Hello World.")
        }
        if !hasNextState() {
            gotoState(funcIdsNEWERA)
        }
    }
    /*
        S_LOP000
    */
    var sLOP000 = func ( bFirst bool ) {
        counter = 0
        gotoState(funcIdsLOP000LoopCheck)
    }
    var sLOP000LoopCheck = func ( bFirst bool ) {
        if counter < 10 {
            gosubState(funcIdsSBS001, funcIdsLOP000LoopNext)
        } else {
            gotoState(funcIdsEND)
        }
    }
    var sLOP000LoopNext = func(bFirst bool ) {
        counter++
        gotoState(funcIdsLOP000LoopCheck)
    }
    /*
        S_NEWERA
    */
    var sNEWERA = func( bFirst  bool ) {
        if bFirst {
            fmt.Println("We are in the visual programming era.")
        }
        if !hasNextState() {
            gotoState(funcIds0001)
        }
    }
    /*
        S_PAS000
    */
    var sPAS000 = func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsGSB000)
        }
    }
    /*
        S_PAS001
    */
    var sPAS001 = func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIdsLOP000)
        }
    }
    /*
        S_RET000
    */
    var sRET000 = func ( bFirst bool ) {
        returnState()
    }
    /*
        S_RET001
    */
    var sRET001 = func ( bFirst bool ) {
        returnState()
    }
    /*
        S_SBS000
    */
    var sSBS000 = func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIds0005)
        }
    }
    /*
        S_SBS001
    */
    var sSBS001 = func( bFirst  bool ) {
        if !hasNextState() {
            gotoState(funcIds0006)
        }
    }
    /*
        S_START
    */
    var sSTART = func( bFirst  bool ) {
        gotoState(funcIdsHELLOWORLD)
    }


	//[STATEGO OUTPUT END]

	var funclist = [...]func(bool){

		//[STATEGO OUTPUT START] indent(8) $/^S_./->#funclist$
        //             psggConverterLib.dll converted from psgg-file:testControl.psgg

        s0001,
        s0002,
        s0003,
        s0004,
        s0005,
        s0006,
        s0007,
        sEND,
        sGSB000,
        sHELLOWORLD,
        sLOP000,
        sLOP000LoopCheck,
        sLOP000LoopNext,
        sNEWERA,
        sPAS000,
        sPAS001,
        sRET000,
        sRET001,
        sSBS000,
        sSBS001,
        sSTART,


		//[STATEGO OUTPUT END]
		endofFuncList}

	nextfunc = funcIdsSTART

	for curfunc != funcIdsEND {
		var bFirst = false
		if nextfunc != -1 {
			curfunc = nextfunc
			nextfunc = -1
			bFirst = true
		}
		if curfunc != -1 {
			funclist[curfunc](bFirst)
		}
	}
}

/*  :::: PSGG MACRO ::::
:psgg-macro-start

commentline=// {%0}

#setids=@@@
var funcId[[state>>lc]] = id
id++
<<<?state-typ/^loop$/
var funcId[[state>>lc]]LoopCheck = id
id++
var funcId[[state>>lc]]LoopNext = id
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
