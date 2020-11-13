// Play Global
package pg

import (
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten"
)

/*
	GLOBALS
*/
var ScreenWidth = 320
var ScreenHeight = 240

var Screen *ebiten.Image
var GophersImage *ebiten.Image
var Count int

/*
	State Machine Control Update
*/
var updateList []func(bool) bool

// AddUpdateList ... Add control function into map. Reture handle
func AddUpdateList(cf func(bool) bool) int {
	hundle := len(updateList)
	updateList = append(updateList, cf)
	return hundle
}
func PlayUpdate() {
	for i := 0; i < len(updateList); i++ {
		(updateList[i])(false)
	}
}

/*
	Draw List
*/
var DrawBgList []func()
var DrawStageList []func()
var DrawFeList []func()

func AppendDrawBg(f func()) {
	DrawBgList = append(DrawBgList, f)
}
func AppendDrawStage(f func()) {
	DrawStageList = append(DrawStageList, f)
}
func AppendDrawFe(f func()) {
	DrawFeList = append(DrawFeList, f)
}

func DoDraw() {
	for i := 0; i < len(DrawBgList); i++ {
		(DrawBgList[i])()
	}
	for i := 0; i < len(DrawStageList); i++ {
		(DrawStageList[i])()
	}
	for i := 0; i < len(DrawFeList); i++ {
		(DrawFeList[i])()
	}
}
