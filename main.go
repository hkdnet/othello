package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

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
	board      Board
	frameCount int
	cursor     Vec2

	turnPlayer Cell

	pressedKey ebiten.Key
}

var UpLeft, Up, UpRight, Left, Right, DownLeft, Down, DownRight Vec2

var v8 [8]Vec2

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
	b := make([][]Cell, size)
	for i := 0; i < size; i += 1 {
		b[i] = make([]Cell, size)
	}

	x := (size / 2) - 1
	b[x][x] = White
	b[x][x+1] = Black
	b[x+1][x] = Black
	b[x+1][x+1] = White

	return &Game{board: b, turnPlayer: Black}
}

func toText(g *Game) string {
	s := ""
	possibles := g.PossibleXys()
	for y := -1; y <= size; y++ {
		for x := -1; x <= size; x++ {
			if (g.frameCount/10)%2 == 0 && x == g.cursor.x && y == g.cursor.y {
				s += " "
			} else {
				c := g.GetCell(Vec2{x, y})
				if c == Empty {
					isColored := false
					m, xOk := possibles[x]
					if xOk {
						b, yOk := m[y]
						if yOk && b {
							isColored = true
						}
					}
					if isColored {
						s += ":"
					} else {
						s += fmt.Sprint(c)
					}
				} else {
					s += fmt.Sprint(c)
				}
			}
		}
		s += "\n"
	}
	return s
}

type Vec2 struct {
	x int
	y int
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{x: v1.x + v2.x, y: v1.y + v2.y}
}

func (g *Game) GetCell(xy Vec2) Cell {
	if xy.x < 0 || xy.x >= size || xy.y < 0 || xy.y >= size {
		return Wall
	}
	return g.board[xy.y][xy.x]
}
func (g *Game) SetCell(v Vec2, c Cell) {
	g.board[v.y][v.x] = c
}

func (g *Game) SetCurrentCell(c Cell) {
	g.SetCell(g.cursor, c)
}

func (g *Game) GetCurrentCell() Cell {
	return g.GetCell(g.cursor)
}

func (g *Game) PossibleXys() map[int]map[int]bool {
	ret := make(map[int]map[int]bool)
	for x := 0; x < size; x += 1 {
		ret[x] = make(map[int]bool)

		for y := 0; y < size; y += 1 {
			xy := Vec2{x: x, y: y}
			if g.GetCell(xy) == Empty && len(g.ReversedXys(xy)) > 0 {
				ret[x][y] = true
			}
		}
	}
	return ret
}

func (g *Game) ReversedXys(start Vec2) []Vec2 {
	var ret []Vec2

	for _, v := range v8 {
		adjXy := start.Add(v)
		adj := g.GetCell(adjXy)
		var tmp []Vec2
		isOp := adj == Black && g.turnPlayer == White || adj == White && g.turnPlayer == Black
		if isOp {
			tmp = append(tmp, adjXy)
			for {
				adjXy = adjXy.Add(v)
				adj = g.GetCell(adjXy)
				if adj == Wall || adj == Empty {
					break
				}
				if adj == g.turnPlayer {
					ret = append(ret, tmp...)
					break
				} else {
					tmp = append(tmp, adjXy)
				}
			}
		}
	}
	return ret
}

func (g *Game) IsValidCell() bool {
	if g.GetCurrentCell() != Empty {
		return false
	}
	cells := g.ReversedXys(g.cursor)

	return len(cells) > 0
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
			if g.cursor.x > 0 {
				g.cursor.x -= 1
			}
		case ebiten.KeyRight:
			if g.cursor.x < size-1 {
				g.cursor.x += 1
			}
		case ebiten.KeyUp:
			if g.cursor.y > 0 {
				g.cursor.y -= 1
			}
		case ebiten.KeyDown:
			if g.cursor.y < size-1 {
				g.cursor.y += 1
			}
		case ebiten.KeyEnter:
			if g.GetCurrentCell() == Empty {
				xys := g.ReversedXys(g.cursor)
				if len(xys) > 0 {
					g.SetCurrentCell(g.turnPlayer)
					for _, xy := range xys {
						g.SetCell(xy, g.turnPlayer)
					}
					g.ChangeCurrentPlayer()
				}
			}
		}
	}

	g.IncrementCount()
	return nil
}

func (g *Game) ChangeCurrentPlayer() {
	if g.turnPlayer == Black {
		g.turnPlayer = White
	} else {
		g.turnPlayer = Black
	}
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
