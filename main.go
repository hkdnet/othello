package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Cell int

const (
	Empty Cell = iota
	Black
	White
)

type Board [][]Cell

type Game struct {
	board      Board
	frameCount int
	cursor     [2]int

	pressedKey ebiten.Key
}

const size = 6

func NewGame() *Game {
	b := make([][]Cell, size)
	for i := 0; i < size; i += 1 {
		b[i] = make([]Cell, size)
	}

	x := (size / 2) - 1
	b[x][x] = Black
	b[x][x+1] = White
	b[x+1][x] = White
	b[x+1][x+1] = Black

	return &Game{board: b}
}

func toText(g *Game) string {
	s := ""
	for x, r := range g.board {
		for y, c := range r {
			if (g.frameCount/10)%2 == 0 && x == g.cursor[0] && y == g.cursor[1] {
				s += " "
				continue
			}
			switch c {
			case Empty:
				s += "."
			case Black:
				s += "x"
			case White:
				s += "o"
			}
		}
		s += "\n"
	}
	return s
}
func (g *Game) IncrementCount() {
	g.frameCount += 1
	g.frameCount %= 100000
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
	if isNewtral {
		g.pressedKey = -1
	}

	if oldPressed != g.pressedKey {
		switch g.pressedKey {
		case ebiten.KeyLeft:
			if g.cursor[1] > 0 {
				g.cursor[1] -= 1
			}
		case ebiten.KeyRight:
			if g.cursor[1] < size-1 {
				g.cursor[1] += 1
			}
		case ebiten.KeyUp:
			if g.cursor[0] > 0 {
				g.cursor[0] -= 1
			}
		case ebiten.KeyDown:
			if g.cursor[0] < size-1 {
				g.cursor[0] += 1
			}
		}
	}

	g.IncrementCount()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, toText(g))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Othello")
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
