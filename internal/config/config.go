package config

//board
const (
	BOARD_MAX_SHAPES_SIZE = 20
	EMPTY_CHAR            = '-'
)

type Board struct {
	MaxBoardShapes int
	EmptyChar      rune
	Height         int
	Width          int
}

func GetBoardConfig(height, width int) Board {
	return Board{
		EmptyChar:      EMPTY_CHAR,
		Height:         height,
		Width:          width,
		MaxBoardShapes: BOARD_MAX_SHAPES_SIZE,
	}
}
