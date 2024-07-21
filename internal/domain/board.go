package domain

import (
	"fmt"
	"strings"

	"github.com/luchojuarez/retroscreen/internal/config"
)

type Board struct {
	shapesCount int
	height      int
	width       int
	emprtyChar  rune
	stack       []*ShapeI
}

func NewRectanglesBoard(cfg config.Board) Board {
	return Board{
		height:      cfg.Height, // alto
		width:       cfg.Width,  // ancho
		emprtyChar:  cfg.EmptyChar,
		shapesCount: 0,
		stack:       make([]*ShapeI, config.BOARD_MAX_SHAPES_SIZE),
	}
}

func (b *Board) Draw() {
	sliceBoard := b.getEmptyBoard()

	for i := 0; i < b.shapesCount; i++ {
		shape := (*b.stack[i]).(Rectangle)
		sliceBoard = shape.Draw(sliceBoard)
	}

	print(sliceBoard)
}

func (b *Board) GetShape(symbol Symbol) (*ShapeI, int) {
	for i, shapeI := range b.stack {
		if shapeI == nil {
			return nil, -1
		}
		shape := (*shapeI).(ShapeI)

		if shape.GetSymbol() == symbol {
			return shapeI, i
		}
	}
	return nil, -1
}

func (b *Board) Move(symbol Symbol, offsetX, offsetY int) error {
	for i, shapeI := range b.stack {
		if shapeI == nil {
			return nil
		}

		shape := (*shapeI).(ShapeI)
		if shape.GetSymbol() == symbol {
			movedShape, _ := shape.Move(offsetX, offsetY)
			b.stack[i] = &movedShape
		}
	}

	// TODO validate for unsuccessully move
	return nil
}

func (b *Board) getEmptyBoard() DrawableBoard {
	var board DrawableBoard
	for i := 0; i < b.width; i++ {
		board = append(board, []rune(strings.Repeat(fmt.Sprintf("%c", b.emprtyChar), b.height)))
	}

	return board
}

func (b *Board) Delete(symbol Symbol) error {
	for i, shapeI := range b.stack {
		if shapeI == nil {
			return nil
		}
		shape := (*shapeI).(ShapeI)
		if shape.GetSymbol() == symbol {
			b.shapesCount--
			b.stack = append(b.stack[:i], b.stack[i+1:]...)
		}
	}

	return nil
}

func (b *Board) Add(shape ShapeI) error {
	// TODO validate max size
	b.stack[b.shapesCount] = &shape
	b.shapesCount++

	return nil
}

func (b *Board) Combine(master, slave Symbol) error {
	var masterShape ShapeI
	var slaveShape ShapeI
	var masterIndex int
	if m, i := b.GetShape(master); m == nil {
		return fmt.Errorf("master not found '%c'", master)
	} else {
		masterIndex = i
		masterShape = *m
	}

	if s, _ := b.GetShape(slave); s == nil {
		return fmt.Errorf("slave not found '%c'", slave)
	} else {
		slaveShape = *s
	}

	fmt.Printf("combine '%+v' '%+v'", masterShape, slaveShape)

	masterShape, err := masterShape.Combine(masterShape, slaveShape)
	if err != nil {
		return err
	}

	b.Delete(slave)

	b.stack[masterIndex] = &masterShape
	return nil
}

func print(board [][]rune) {
	for _, row := range board {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Printf("\n")
	}
}
