package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var BlackImage, WhiteImage, CursorImage, EmptyCellImage, BackgroundImage *ebiten.Image

var BackgroundDrawImageOptions *ebiten.DrawImageOptions

var BackgruondColor = color.RGBA{G: 102, A: 255}

var UpLeft, Up, UpRight, Left, Right, DownLeft, Down, DownRight Vec2

const cellWidth = 16
const boardTopLeftX = 120.0 - (size/2)*cellWidth
const boardTopLeftY = 60.0 - (size/2)*cellWidth

var v8 [8]Vec2

func init() {
	// color
	BlackImage = ebiten.NewImage(16, 16)
	BlackImage.Fill(color.Transparent)
	ebitenutil.DrawCircle(BlackImage, 8, 8, 7, color.Black)

	WhiteImage = ebiten.NewImage(16, 16)
	WhiteImage.Fill(color.Transparent)
	ebitenutil.DrawCircle(WhiteImage, 8, 8, 7, color.White)

	EmptyCellImage = ebiten.NewImage(16, 16)
	EmptyCellImage.Fill(color.Black)
	innerCursor := ebiten.NewImage(14, 14)
	innerCursor.Fill(BackgruondColor)
	innerGm := ebiten.GeoM{}
	innerGm.Translate(1.0, 1.0)
	EmptyCellImage.DrawImage(innerCursor, &ebiten.DrawImageOptions{GeoM: innerGm})

	CursorImage = ebiten.NewImage(16, 16)
	CursorImage.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 128})

	BackgroundImage = ebiten.NewImage(640, 480)
	BackgroundImage.Fill(BackgruondColor)
	for y := 0; y < size; y++ {
		gm := ebiten.GeoM{}
		gm.Translate(boardTopLeftX, boardTopLeftY+float64(y)*cellWidth)
		for x := 0; x < size; x++ {
			opt := &ebiten.DrawImageOptions{GeoM: gm}
			BackgroundImage.DrawImage(EmptyCellImage, opt)
			gm.Translate(cellWidth, 0)
		}
	}

	BackgroundDrawImageOptions = &ebiten.DrawImageOptions{}

	// vec2
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
type BattleMode struct {
	board  Board
	cursor Vec2

	turnPlayer Cell

	previousTurnSkipped bool
	GameSet             bool
}

func toOpponent(c Cell) Cell {
	if c == Black {
		return White
	} else {
		return Black
	}
}

func newBattleMode() *BattleMode {
	b := make([][]Cell, size)
	for i := 0; i < size; i += 1 {
		b[i] = make([]Cell, size)
	}

	x := (size / 2) - 1
	b[x][x] = White
	b[x][x+1] = Black
	b[x+1][x] = Black
	b[x+1][x+1] = White

	return &BattleMode{board: b, turnPlayer: Black}
}

func (bm *BattleMode) Draw(screen *ebiten.Image, g *Game) {
	screen.DrawImage(BackgroundImage, BackgroundDrawImageOptions)

	for y := 0; y <= size; y++ {
		gm := ebiten.GeoM{}
		gm.Translate(boardTopLeftX, boardTopLeftY+float64(y)*cellWidth)
		for x := 0; x <= size; x++ {
			c := bm.GetCell(Vec2{x, y})

			opt := &ebiten.DrawImageOptions{GeoM: gm}
			if c == Black {
				screen.DrawImage(BlackImage, opt)
			} else if c == White {
				screen.DrawImage(WhiteImage, opt)
			}
			gm.Translate(cellWidth, 0)
		}
	}
	if (g.frameCount/10)%2 == 0 {
		cursorGm := ebiten.GeoM{}
		cursorGm.Translate(boardTopLeftX+float64(bm.cursor.x)*cellWidth, boardTopLeftY+float64(bm.cursor.y)*cellWidth)
		screen.DrawImage(CursorImage, &ebiten.DrawImageOptions{GeoM: cursorGm})
	}
}

func (bm *BattleMode) GetCell(xy Vec2) Cell {
	if xy.x < 0 || xy.x >= size || xy.y < 0 || xy.y >= size {
		return Wall
	}
	return bm.board[xy.y][xy.x]
}
func (bm *BattleMode) SetCell(v Vec2, c Cell) {
	bm.board[v.y][v.x] = c
}

func (bm *BattleMode) SetCurrentCell(c Cell) {
	bm.SetCell(bm.cursor, c)
}

func (bm *BattleMode) GetCurrentCell() Cell {
	return bm.GetCell(bm.cursor)
}

func (bm *BattleMode) PossibleXys() map[int]map[int]bool {
	ret := make(map[int]map[int]bool)
	for x := 0; x < size; x += 1 {
		ret[x] = make(map[int]bool)

		for y := 0; y < size; y += 1 {
			xy := Vec2{x: x, y: y}
			if bm.GetCell(xy) == Empty && len(bm.ReversedXys(xy)) > 0 {
				ret[x][y] = true
			}
		}
	}
	return ret
}

func (bm *BattleMode) ReversedXys(start Vec2) []Vec2 {
	var ret []Vec2

	for _, v := range v8 {
		adjXy := start.Add(v)
		adj := bm.GetCell(adjXy)
		var tmp []Vec2
		isOp := adj == toOpponent(bm.turnPlayer)
		if isOp {
			tmp = append(tmp, adjXy)
			for {
				adjXy = adjXy.Add(v)
				adj = bm.GetCell(adjXy)
				if adj == Wall || adj == Empty {
					break
				}
				if adj == bm.turnPlayer {
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

func (bm *BattleMode) IsValidCell() bool {
	if bm.GetCurrentCell() != Empty {
		return false
	}
	cells := bm.ReversedXys(bm.cursor)

	return len(cells) > 0
}

func (bm *BattleMode) Left(g *Game) {
	if bm.cursor.x > 0 {
		bm.cursor.x -= 1
	}
}

func (bm *BattleMode) Right(g *Game) {
	if bm.cursor.x < size-1 {
		bm.cursor.x += 1
	}
}
func (bm *BattleMode) Up(g *Game) {
	if bm.cursor.y > 0 {
		bm.cursor.y -= 1
	}
}

func (bm *BattleMode) Down(g *Game) {
	if bm.cursor.y < size-1 {
		bm.cursor.y += 1
	}
}

func (bm *BattleMode) Enter(g *Game) {
	if bm.GetCurrentCell() == Empty {
		xys := bm.ReversedXys(bm.cursor)
		if len(xys) == 0 {
			// invalid input, ignore the enter
			return
		}
		bm.SetCurrentCell(bm.turnPlayer)
		for _, xy := range xys {
			bm.SetCell(xy, bm.turnPlayer)
		}
		bm.ChangeCurrentPlayer()
		if bm.HasValidCell() {
			bm.previousTurnSkipped = false
		} else {
			bm.previousTurnSkipped = true
			bm.ChangeCurrentPlayer()
			if !bm.HasValidCell() {
				bm.GameSet = true
			}
		}
	}
}

func (bm *BattleMode) HasValidCell() bool {
	for _, m := range bm.PossibleXys() {
		for _, ok := range m {
			if ok {
				return true
			}
		}
	}
	return false
}
func (bm *BattleMode) ChangeCurrentPlayer() {
	if bm.turnPlayer == Black {
		bm.turnPlayer = White
	} else {
		bm.turnPlayer = Black
	}
}
