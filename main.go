package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ichibankunio/podiumpuzzle/mygame"
)

func main() {
	ebiten.SetWindowSize(540, 960)
	ebiten.SetWindowTitle("Hello, World!")

	mainGame := &mygame.Game{
		State: 0,
		ButtonData: "100,24,84,84;228,24,84,84;456,24,84,84;90,550,150,100",
		// Time1: 0,
		// Time2: 0,
		// Score1: 0,
		// Score2: 0,
		Record: []string{"0.0", "0", "0.0", "0"},
		Result: []string{"", "", "", ""},
	}
	if err := ebiten.RunGame(mainGame); err != nil {
		log.Fatal(err)
	}
}
