package domain

import (
	"fmt"

	"github.com/luchojuarez/retroscreen/internal/config"
)

type Machine struct {
	board *Board
}

func NewMachine(boardCfg config.Board) Machine {
	b := NewRectanglesBoard(boardCfg)

	return Machine{
		board: &b,
	}
}

func (m *Machine) Move(symbol Symbol, offsetX, offsetY int) error {
	m.board.Move(symbol, offsetX, offsetY)
	return nil
}
func (m *Machine) Add(shape ShapeI) error {
	if shape, _ := m.board.GetShape(shape.GetSymbol()); shape != nil {
		return fmt.Errorf("symbol '%+v' already exist", shape)
	}

	m.board.Add(shape)
	return nil
}
func (m *Machine) Render() {
	m.board.Draw()
}

func (m *Machine) Delete(symbol Symbol) {
	m.board.Delete(symbol)
}

func (m *Machine) Combine(master, slave Symbol) {
	m.board.Combine(master, slave)
}
