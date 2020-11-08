package main

import "fmt"

var stateMap map[string]func(bool);

var mCurfunc func(bool)

func main() {
    stateMap = map[string]func(bool) {
        "S_START" :         

    }
	mCurfunc = xTest
	test()
}

func test() {
	fmt.Println("you are in test")
}

func xTest(b bool) {
	if b {
		fmt.Println("you are in X_TEST first")
	}
	return
}
