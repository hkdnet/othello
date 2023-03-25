package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Vec2 struct {
	x int
	y int
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{x: v1.x + v2.x, y: v1.y + v2.y}
}

type Cell int

const (
	Empty Cell = iota
	Black
	White
	Wall
)

func (c Cell) String() string {
	switch c {
	case Empty:
		return "."
	case Black:
		return "x"
	case White:
		return "o"
	case Wall:
		return "#"

	}
	return ""
}

type Board [][]Cell

type Game struct {
	handler handler

	pressedKey ebiten.Key
	frameCount int
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

const size = 6

func NewGame() *Game {
	initialMode := newBattleMode()
	return &Game{handler: initialMode}
}
func (g *Game) Update() error {
	oldPressed := g.pressedKey
	isNewtral := true

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.pressedKey = ebiten.KeyLeft
		isNewtral = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.pressedKey = ebiten.KeyRight
		isNewtral = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.pressedKey = ebiten.KeyUp
		isNewtral = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.pressedKey = ebiten.KeyDown
		isNewtral = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.pressedKey = ebiten.KeyEnter
		isNewtral = false
	}

	if isNewtral {
		g.pressedKey = -1
	}

	if oldPressed != g.pressedKey {
		switch g.pressedKey {
		case ebiten.KeyLeft:
			g.handler.Left(g)
		case ebiten.KeyRight:
			g.handler.Right(g)
		case ebiten.KeyUp:
			g.handler.Up(g)
		case ebiten.KeyDown:
			g.handler.Down(g)
		case ebiten.KeyEnter:
			g.handler.Enter(g)
		}
	}

	g.IncrementCount()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.handler.ToText(g))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
