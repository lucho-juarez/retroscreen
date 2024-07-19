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

func (b *Board) GetShape(symbol Symbol) *ShapeI {
	for _, shapeI := range b.stack {
		if shapeI == nil {
			return nil
		}
		shape := (*shapeI).(ShapeI)

		if shape.GetSymbol() == symbol {
			return shapeI
		}
	}
	return nil
}

func (b *Board) Move(symbol Symbol, offsetX, offsetY int) error {
	for i, shapeI := range b.stack {
		if shapeI == nil {
			return nil
		}

		shape := (*shapeI).(ShapeI)
		if shape.GetSymbol() == symbol {
			movedShape, _ := shape.Move(symbol, offsetX, offsetY)
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

func print(board [][]rune) {
	for _, row := range board {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Printf("\n")
	}
}
