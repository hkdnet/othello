package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MenuMode struct {
	selected int
}

type MenuItem int

const (
	NewGameItem MenuItem = iota
	QuitItem
)

var menuItems []MenuItem

func init() {
	menuItems = []MenuItem{NewGameItem, QuitItem}
}

func (m MenuItem) String() string {
	switch m {
	case NewGameItem:
		return "New Game"
	case QuitItem:
		return "Quit"
	}
	return "N/A"
}

func NewMenuMode() *MenuMode {
	return &MenuMode{}
}

func (mm *MenuMode) Draw(screen *ebiten.Image, g *Game) {
	s := ""
	for idx, item := range menuItems {
		if (g.frameCount/10)%2 == 0 && idx == mm.selected {
			s += ""
		} else {
			s += item.String()
		}
		s += "\n"
	}

	ebitenutil.DebugPrint(screen, s)
}

func (mm *MenuMode) Left(g *Game) {
	// do nothing
}

func (mm *MenuMode) Right(g *Game) {
	// do nothing
}
func (mm *MenuMode) Up(g *Game) {
	mm.selected += len(menuItems) - 1
	mm.selected %= len(menuItems)
}

func (mm *MenuMode) Down(g *Game) {
	mm.selected += 1
	mm.selected %= len(menuItems)
}

func (mm *MenuMode) Enter(g *Game) {
	switch menuItems[mm.selected] {

	case NewGameItem:
		g.handler = newBattleMode()
	case QuitItem:
		g.Quit()
	}
}
