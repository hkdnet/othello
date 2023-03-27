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
	screen.DrawImage(BackgroundImage, BackgroundDrawImageOptions)
	screen.DrawImage(BlackImage, BackgroundDrawImageOptions)
	screen.DrawImage(CursorImage, BackgroundDrawImageOptions)
}
