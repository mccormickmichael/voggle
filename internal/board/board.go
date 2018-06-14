package board

import (
	"math/rand"
	"strings"
)

const (
	BoardSize = 5
)

type Board struct { // TODO: this could be type [][]*Cell
	cells [][]*Cell // TODO: this could be a flat array
}

type Cell struct {
	Row int
	Col int
	Value string
	Neighbors []*Cell
}

func RandomBoard(source rand.Source) *Board {
	r := rand.New(source)

	board := Board{}
	board.cells = make([][]*Cell, BoardSize)
	bcells := make([]*Cell, BoardSize * BoardSize)
	for i := 0; i < BoardSize; i++ {
		board.cells[i], bcells = bcells[:BoardSize], bcells[BoardSize:]
	}

	cs := make([]Cube, len(cubes), len(cubes))
	copy(cs, cubes)

	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			c := &Cell{Row: row, Col: col}
			cubeIndex := r.Intn(len(cs))
			cube := cs[cubeIndex]

			// remove cube from cs
			copy(cs[cubeIndex:], cs[cubeIndex+1:])
			cs = cs[:len(cs)-1]

			c.Value = cube[r.Intn(len(cube))]

			board.cells[row][col] = c
		}
	}
	board.resolveNeighbors()

	return &board
}

func (b *Board) resolveNeighbors() {
	max := BoardSize - 1
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			var nb []*Cell
			cell := b.cells[row][col]
			switch {
				// corners
			case row == 0 && col == 0:
				nb = make([]*Cell, 3)
				nb[0] = b.cells[row+1][col]
				nb[1] = b.cells[row+1][col+1]
				nb[2] = b.cells[row][col+1]
			case row == 0 && col == max:
				nb = make([]*Cell, 3)
				nb[0] = b.cells[row][col-1]
				nb[1] = b.cells[row+1][col-1]
				nb[2] = b.cells[row+1][col]
			case row == max && col == 0:
				nb = make([]*Cell, 3)
				nb[0] = b.cells[row-1][col]
				nb[1] = b.cells[row][col+1]
				nb[2] = b.cells[row-1][col+1]
			case row == max && col == max:
				nb = make([]*Cell, 3)
				nb[0] = b.cells[row-1][col]
				nb[1] = b.cells[row-1][col-1]
				nb[2] = b.cells[row][col-1]
				// edges
			case row == 0: // top row sans corners
				nb = make([]*Cell, 5)
				nb[0] = b.cells[row][col-1]
				nb[1] = b.cells[row+1][col-1]
				nb[2] = b.cells[row+1][col]
				nb[3] = b.cells[row+1][col+1]
				nb[4] = b.cells[row][col+1]
			case row == max: // bottom row sans corners
				nb = make([]*Cell, 5)
				nb[0] = b.cells[row-1][col]
				nb[1] = b.cells[row-1][col-1]
				nb[2] = b.cells[row][col-1]
				nb[3] = b.cells[row][col+1]
				nb[4] = b.cells[row-1][col+1]
			case col == 0: // left col sans corners
				nb = make([]*Cell, 5)
				nb[0] = b.cells[row-1][col]
				nb[1] = b.cells[row+1][col]
				nb[2] = b.cells[row+1][col+1]
				nb[3] = b.cells[row][col+1]
				nb[4] = b.cells[row-1][col+1]
			case col == max: // right col sans corners
				nb = make([]*Cell, 5)
				nb[0] = b.cells[row-1][col]
				nb[1] = b.cells[row-1][col-1]
				nb[2] = b.cells[row][col-1]
				nb[3] = b.cells[row+1][col-1]
				nb[4] = b.cells[row+1][col]
			default:
				nb = make([]*Cell, 8)
				nb[0] = b.cells[row-1][col]
				nb[1] = b.cells[row-1][col-1]
				nb[2] = b.cells[row][col-1]
				nb[3] = b.cells[row+1][col-1]
				nb[4] = b.cells[row+1][col]
				nb[5] = b.cells[row+1][col+1]
				nb[6] = b.cells[row][col+1]
				nb[7] = b.cells[row-1][col+1]
			}
			cell.Neighbors = nb
		}
	}
}

func (b *Board) At(row, col int) *Cell {
	return b.cells[row][col]
}

func (b *Board) String() string {
	var sb strings.Builder
	for _, r := range b.cells {
		for _, c := range r {
			sb.WriteString(c.Value)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
