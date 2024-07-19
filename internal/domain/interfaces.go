package domain

var _ ShapeI = Rectangle{}

type BaseShape struct {
	Symbol Symbol
	X      int
	Y      int
}

// symbol is an unique char, used as identifyer and cant be duplicated.
type Symbol rune

type DrawableBoard [][]rune

type ShapeI interface {
	GetSymbol() Symbol
	Draw(input DrawableBoard) DrawableBoard
	Move(symbol Symbol, offsetX, offsetY int) (ShapeI, error)
}

func (s *BaseShape) Move(symbol Symbol, offsetX, offsetY int) error {
	s.X += offsetX
	s.Y += offsetY
	return nil
}

func (s *BaseShape) GetSymbol() Symbol {
	return s.Symbol
}

type Rectangle struct {
	BaseShape
	Height int
	Width  int
}