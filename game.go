package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	handler handler

	pressedKey ebiten.Key
	frameCount int
}

const size = 6

func NewGame() *Game {
	initialMode := NewMenuMode()
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
	g.handler.Draw(screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) IncrementCount() {
	g.frameCount += 1
	g.frameCount %= 100000
}

func (g *Game) Quit() {
	os.Exit(0)
}
