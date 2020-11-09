package main

import (
	"fmt"
)

func main() {
	testControl()
}

func testControl() {
	var curfunc = -1
	var nextfunc = -1

	var gotoState = func(st int) {
		nextfunc = st
	}
	var hasNextState = func() bool {
		return nextfunc != -1
	}

	// #
	// # Define function ID for state
	// #
	var id = 0
	//[STATEGO OUTPUT START] indent(4) $/^S_./->#setids$
    //             psggConverterLib.dll converted from psgg-file:TestControl.psgg

    var funcIds0001 = id
    id++
    var funcIds0002 = id
    id++
    var funcIds0003 = id
    id++
    var funcIds0004 = id
    id++
    var funcIdsEND = id
    id++
    var funcIdsHELLOWORLD = id
    id++
    var funcIdsNEWERA = id
    id++
    var funcIdsSTART = id
    id++


	//[STATEGO OUTPUT END]

	// #
	// #  State Function
	// #
	var endofFuncList = func(bFirst bool) { // for end of function list
	}
	//[STATEGO OUTPUT START] indent(4) $/./$
    //             psggConverterLib.dll converted from psgg-file:TestControl.psgg

    /*
        S_0001
    */
    var s0001 = func( bFirst  bool ) {
        var n = 2
        if n==0 {
            gotoState( funcIds0004 )
        } else if n==1 {
            gotoState( funcIds0002 )
        } else if n==2 {
            gotoState( funcIds0003 )
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
            gotoState(funcIdsEND)
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
            gotoState(funcIdsEND)
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
            gotoState(funcIdsEND)
        }
    }
    /*
        S_END
    */
    var sEND = func ( bFirst  bool ) {
         // end of state machine
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
        S_START
    */
    var sSTART = func( bFirst  bool ) {
        gotoState(funcIdsHELLOWORLD)
    }


	//[STATEGO OUTPUT END]

	var funclist = [...]func(bool){

		//[STATEGO OUTPUT START] indent(8) $/^S_./->#funclist$
        //             psggConverterLib.dll converted from psgg-file:TestControl.psgg

        s0001,
        s0002,
        s0003,
        s0004,
        sEND,
        sHELLOWORLD,
        sNEWERA,
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
@@@

#funclist=@@@
[[state>>lc]],
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
