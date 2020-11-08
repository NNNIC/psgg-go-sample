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
	var S_START = func(bFirst bool) {
		if bFirst {
			fmt.Println("S_START")
			nextfunc = "S_WORK"
		}
	}
	var sEnd = func(bFirst bool) {
		if bFirst {
			fmt.Println("S_END")
		}
	}
	var sWork = func(bFirst bool) {
		nextfunc = "S_END"
	}
	var funcmap = make(map[string]func(bool))
	funcmap["S_START"] = S_START
	funcmap["S_END"] = sEnd
	funcmap["S_WORK"] = sWork

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

// state
//[STATEGO OUTPUT START] indent(0) $/./$
//             psggConverterLib.dll converted from psgg-file:TestControl.psgg

//[STATEGO OUTPUT END]

//Run
func run() {
	// fmt.Println("RUN!")
	// initialize()
	// gotoState("S_START")
	// for checkState("S_END") == false {
	// 	update()
	// }
	// fmt.Println("END OF RUN")
}

//let sm = TestControl()
//sm.Run()
