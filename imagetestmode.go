package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageTestMode struct {
}

func NewImageTestMode() *ImageTestMode {
	return &ImageTestMode{}
}

func (itm *ImageTestMode) Left(g *Game) {
}

func (itm *ImageTestMode) Right(g *Game) {
}
func (itm *ImageTestMode) Up(g *Game) {
}

func (itm *ImageTestMode) Down(g *Game) {
}

func (itm *ImageTestMode) Enter(g *Game) {
}

func (itm *ImageTestMode) Draw(screen *ebiten.Image, g *Game) {
	gm := ebiten.GeoM{}

	screen.DrawImage(BackgroundImage, &ebiten.DrawImageOptions{GeoM: gm})
	screen.DrawImage(BlackImage, &ebiten.DrawImageOptions{GeoM: gm})
	gm.Translate(cellWidth, 0)
	screen.DrawImage(CursorImage, &ebiten.DrawImageOptions{GeoM: gm})
	gm.Translate(cellWidth, 0)
	screen.DrawImage(ValidCellImage, &ebiten.DrawImageOptions{GeoM: gm})
}
