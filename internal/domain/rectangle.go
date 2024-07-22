package domain

func (r Rectangle) Draw(input DrawableBoard) DrawableBoard {
	//outside for used forhorizontal select
	for i := r.BaseShape.X; i < r.BaseShape.X+r.Width && i < len(input[0]); i++ {
		for j := r.BaseShape.Y; j < r.BaseShape.Y+r.Height && j < len(input); j++ {
			input[j][i] = rune(r.BaseShape.Symbol)
		}
	}

	if r.Child != nil {
		input = r.Child.Draw(input)
	}

	return input
}

func (r Rectangle) Move(offsetX, offsetY int) (ShapeI, error) {
	r.BaseShape.Move(offsetX, offsetY)

	if r.Child != nil {
		if child, e := r.Child.Move(offsetX, offsetY); e != nil {
			return r, e
		} else {
			c, _ := child.(Rectangle)
			r.Child = &c
		}
	}

	return r, nil
}

func (r Rectangle) GetSymbol() Symbol {
	return r.Symbol
}

func (r Rectangle) Combine(master, slave ShapeI) (ShapeI, error) {

	if slave == nil {
		return master, nil
	}

	// TODO refactor, this coult be better using Shapes, not rectangle.
	slaveParsed := slave.(Rectangle)
	masterParsed := master.(Rectangle)

	//slaves took master's symbol for slave and their childrens
	slaveParsed.BaseShape.Symbol = master.GetSymbol()
	slaveChild := slaveParsed.Child
	for slaveChild != nil {
		slaveChild.BaseShape.Symbol = master.GetSymbol()
		slaveChild = slaveChild.Child
	}

	//merge both children chains
	if masterParsed.Child != nil {
		lastChild := masterParsed.Child
		for lastChild == nil && lastChild.Child == nil {
			lastChild = lastChild.Child
		}
		lastChild.Child = &slaveParsed

	} else {
		masterParsed.Child = &slaveParsed
	}

	return masterParsed, nil
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
