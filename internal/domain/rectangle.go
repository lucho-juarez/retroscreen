package domain

func (r Rectangle) Draw(input DrawableBoard) DrawableBoard {
	//outside for used forhorizontal select
	for i := r.BaseShape.X; i < r.BaseShape.X+r.Width && i < len(input[0]); i++ {
		for j := r.BaseShape.Y; j < r.BaseShape.Y+r.Height && j < len(input); j++ {
			input[j][i] = rune(r.BaseShape.Symbol)
		}
	}
	return input
}

func (r Rectangle) Move(symbol Symbol, offsetX, offsetY int) (ShapeI, error) {
	//TODo validate symbol
	r.BaseShape.Move(symbol, offsetX, offsetY)

	return r, nil
}

func (r Rectangle) GetSymbol() Symbol {
	return r.Symbol
}
func NewRectangle(symbol Symbol, x, y, height, width int) Rectangle {
	return Rectangle{
		BaseShape: BaseShape{
			Symbol: symbol,
			X:      x,
			Y:      y,
		},
		Height: height,
		Width:  width,
	}
}
