package main

import (
	"fmt"
)

func main() {
	run()
}

var mCurfunc = ""
var mNextfunc = ""
var mNoWait = false
var mFuncMap map[string]func(bool)

func initialize() {
	mCurfunc := ""
	mNextfunc := ""
	mFuncMap := make(map[string]func(bool))
	mFuncMap["S_START"] = sSTART
	mFuncMap["S_END"] = sEND
}

func update() {
	for true {
		var bFirst = false
		mNoWait := false
		if mNextfunc == "" {
			mCurstate := mNextfunc
			mNextfunc := ""
			bFirst := true
		}
		if mCurfunc != "" {
			mFuncMap[mCurfunc](bFirst)
		}
		if mNoWait == false {
			break
		}
	}
}

func gotoState(st string) {
	mNextfunc = st
}
func checkState(st string) bool {
	return st == mCurfunc
}
func hasNextState() bool {
	return mNextfunc != ""
}
func noWait() {
	mNoWait := true
}
func GoSubState(_, _ string) {
	// if m_callstack_level >= MAX_CALL_STACK-1 {
	//     print("CALL STACK OVER FLOW")
	// }
	// else {
	//     m_callstack[m_callstack_level] = next;
	//     m_callstack_level += 1
	//     Goto(sub)
	// }
}
func returnState() {
	// var st: String? = nil;
	// if m_callstack_level <= 0 {
	//     print("CALL STACK UNDER FLOW")
	// }
	// else {
	//     m_callstack_level -= 1
	//     st = m_callstack[m_callstack_level]
	//     Goto(st)
	// }
}

// state
//[STATEGO OUTPUT START] indent(0) $/./$
//             psggConverterLib.dll converted from psgg-file:TestControl.psgg

/*
    S_END
*/
func sEND( bFirst  bool ) {
     // end of state machine
}
/*
    S_HELLOWORLD
*/
func sHELLOWORLD( bFirst  bool ) {
    if bFirst {
        fmt.Println("Hello World.")
    }
    if !hasNextState() {
        gotoState("S_NEWERA");
    }
}
/*
    S_NEWERA
*/
func sNEWERA( bFirst  bool ) {
    if bFirst {
        fmt.Println("We are in the visual programming era.")
    }
    if !hasNextState() {
        gotoState("S_END");
    }
}
/*
    S_START
*/
func sSTART( bFirst  bool ) {
    gotoState("S_HELLOWORLD")
    noWait()
}


//[STATEGO OUTPUT END]

//Run
func run() {
	fmt.Println("RUN!")
	initialize()
	gotoState("S_START")
	for checkState("S_END") == false {
		update()
	}
	fmt.Println("END OF RUN")
}

//let sm = TestControl()
//sm.Run()
