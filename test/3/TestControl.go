package main

import (
	"fmt"
)

func main() {
	testControl()
}

func testControl() {
	var curfunc = ""
	var nextfunc = ""

	var gotoState = func(st string) {
		nextfunc = st
	}
	var hasNextState = func() bool {
		return nextfunc != ""
	}

	//[STATEGO OUTPUT START] indent(4) $/./$
    //             psggConverterLib.dll converted from psgg-file:TestControl.psgg

    /*
        S_0001
    */
    var s0001 = func( bFirst  bool ) {
        var n = 2
        if n==0 {
            gotoState( "S_0004" )
        } else if n==1 {
            gotoState( "S_0002" )
        } else if n==2 {
            gotoState( "S_0003" )
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
            gotoState("S_END")
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
            gotoState("S_END")
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
            gotoState("S_END")
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
            gotoState("S_NEWERA")
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
            gotoState("S_0001")
        }
    }
    /*
        S_START
    */
    var sSTART = func( bFirst  bool ) {
        gotoState("S_HELLOWORLD")
    }


	//[STATEGO OUTPUT END]

	var funcmap = make(map[string]func(bool))
	//[STATEGO OUTPUT START] indent(4) $/./->#map$
    //             psggConverterLib.dll converted from psgg-file:TestControl.psgg

    funcmap["S_0001"] = s0001
    funcmap["S_0002"] = s0002
    funcmap["S_0003"] = s0003
    funcmap["S_0004"] = s0004
    funcmap["S_END"] = sEND
    funcmap["S_HELLOWORLD"] = sHELLOWORLD
    funcmap["S_NEWERA"] = sNEWERA
    funcmap["S_START"] = sSTART


	//[STATEGO OUTPUT END]

	nextfunc = "S_START"

	for curfunc != "S_END" {
		var bFirst = false
		if nextfunc != "" {
			curfunc = nextfunc
			nextfunc = ""
			bFirst = true
		}
		if curfunc != "" {
			funcmap[curfunc](bFirst)
		}
	}
}

/*  :::: PSGG MACRO ::::
:psgg-macro-start

commentline=// {%0}

#map=@@@
funcmap["[[state]]"] = [[state>>lc]]
@@@

@branch=@@@
<<<?"{%0}"/^brif$/
if [[brcond:{%N}]] {
    gotoState( "{%1}" )
}
>>>
<<<?"{%0}"/^brifc$/
if [[brcond:{%N}]] {
    gotoState( "{%1}" )
>>>
<<<?"{%0}"/^brelseif$/
} else if [[brcond:{%N}]] {
    gotoState( "{%1}" )
}
>>>
<<<?"{%0}"/^brelseifc$/
} else if [[brcond:{%N}]] {
    gotoState( "{%1}" )
>>>
<<<?"{%0}"/^brelse$/
} else {
    gotoState( "{%1}" )
}
>>>
@@@

:psgg-macro-end
*/
