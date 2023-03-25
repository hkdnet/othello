package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var UpLeft, Up, UpRight, Left, Right, DownLeft, Down, DownRight Vec2

var v8 [8]Vec2

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

type Vec2 struct {
	x int
	y int
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{x: v1.x + v2.x, y: v1.y + v2.y}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Othello")
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
