package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type handler interface {
	Left(g *Game)
	Right(g *Game)
	Up(g *Game)
	Down(g *Game)
	Enter(g *Game)

	ToText(g *Game) string
}

func init() {
	UpLeft = Vec2{-1, -1}
	Up = Vec2{-1, 0}
	UpRight = Vec2{-1, +1}
	Left = Vec2{0, -1}
	Right = Vec2{0, +1}
	DownLeft = Vec2{1, -1}
	Down = Vec2{1, 0}
	DownRight = Vec2{1, +1}

	v8 = [8]Vec2{UpLeft, Up, UpRight, Left, Right, DownLeft, Down, DownRight}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Othello")
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
