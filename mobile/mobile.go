package mobile

import (
	// "strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/mobile"

	"github.com/ichibankunio/podiumpuzzle/mygame"
)

var mainGame *mygame.Game

const (
	BUTTTON_TWEET = iota
	BUTTON_APP
)

const (
	SCREEN_WIDTH  = 540
	SCREEN_HEIGHT = 960
)

func init() {
	// yourgame.Game must implement ebiten.Game interface.
	// For more details, see
	// * https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Game

	mainGame = &mygame.Game{
		State: 0,
		ButtonData: "100,24,84,84;228,24,84,84;356,24,84,84;90,550,150,100;90,600,150,100;300,600,150,100;195,690,150,100",
		// Time1: 0,
		// Time2: 0,
		// Score1: 0,
		// Score2: 0,
		Record: []string{"0.0", "0", "0.0", "0"},
		Result: []string{"", "", "", ""},
	}

	mobile.SetGame(mainGame)
	// mobile.SetGame(&mygame.Game{})
}

type GameManager interface {
	// IsJustTouch() bool
	// GetDataArray(data []int)
	SetRecord() string
	// GetRecord(string)
	SetLocation() int
}

func GetHighScore(gm GameManager) string {
	return mainGame.Record[0] + ";" + mainGame.Record[1] + ";" + mainGame.Record[2] + ";" + mainGame.Record[3]
}

func SetHighScore(gm GameManager) {
	s := gm.SetRecord()
	arr := strings.Split(s, ";")

	mainGame.Record[0] = arr[0]
	mainGame.Record[1] = arr[1]
	mainGame.Record[2] = arr[2]
	mainGame.Record[3] = arr[3]
}

func GetResultText(gm GameManager) string {
	return mainGame.Result[0] + ";" + mainGame.Result[1] + ";" + mainGame.Result[2] + ";" + mainGame.Result[3]
}

func GetState(gm GameManager) int {
	return mainGame.State
	// if am.IsJustTouch() {

	// }
}

func GetButtonData(gm GameManager) string {
	return mainGame.ButtonData
}

func GetScreenWidth(gm GameManager) int {
	return SCREEN_WIDTH
}

func GetScreenHeight(gm GameManager) int {
	return SCREEN_HEIGHT
}

func SetLang(gm GameManager) {
	mainGame.Lang = gm.SetLocation()
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
