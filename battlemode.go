package main

import (
	"fmt"
)

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

func (bm *BattleMode) ToText(g *Game) string {
	s := ""
	possibles := bm.PossibleXys()
	for y := -1; y <= size; y++ {
		for x := -1; x <= size; x++ {
			if (g.frameCount/10)%2 == 0 && x == bm.cursor.x && y == bm.cursor.y {
				s += " "
			} else {
				c := bm.GetCell(Vec2{x, y})
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
		isOp := adj == Black && bm.turnPlayer == White || adj == White && bm.turnPlayer == Black
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

func (g *Game) IncrementCount() {
	g.frameCount += 1
	g.frameCount %= 100000
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
		if len(xys) > 0 {
			bm.SetCurrentCell(bm.turnPlayer)
			for _, xy := range xys {
				bm.SetCell(xy, bm.turnPlayer)
			}
			bm.ChangeCurrentPlayer()
		}
	}
}
func (bm *BattleMode) ChangeCurrentPlayer() {
	if bm.turnPlayer == Black {
		bm.turnPlayer = White
	} else {
		bm.turnPlayer = Black
	}
}
