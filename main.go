package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kk-no/ebisample/tictactoe"
)

func main() {
	game, err := tictactoe.New()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(tictactoe.WindowWidth, tictactoe.WindowHight)
	ebiten.SetWindowTitle(tictactoe.Title)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
