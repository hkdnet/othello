package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	isLeftClick bool
}

func (g *Game) Update() error {
	g.isLeftClick = isLeftClick()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	message := ""
	if g.isLeftClick {
		message += "left click"
		g.isLeftClick = false
	}
	ebitenutil.DebugPrint(screen, message+"\nHello World")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func isLeftClick() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyEnter)
}
