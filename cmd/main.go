package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/luchojuarez/retroscreen/input"
	"github.com/luchojuarez/retroscreen/internal/config"
	"github.com/luchojuarez/retroscreen/internal/domain"
)

func main() {

	// validate input data
	if len(os.Args) != 3 {
		panic(fmt.Sprintf("invalid arguments '%+v', init size must be two integers", os.Args[1:]))
	}

	strHeight, strWidth := os.Args[1], os.Args[2]
	height, err := strconv.Atoi(strHeight)
	if err != nil {
		panic(fmt.Sprintf("invalid height argument '%+v' must be integer", os.Args[1]))
	}
	width, err := strconv.Atoi(strWidth)
	if err != nil {
		panic(fmt.Sprintf("invalid width argument '%+v' must be integer", os.Args[2]))
	}

	boardConfig := config.GetBoardConfig(height, width)
	machine := domain.NewMachine(boardConfig)

	machine.Render()
	action, params, err := input.ParseInput(input.PromptInput(input.Help()))

	for {
		for err != nil {
			action, params, err = input.ParseInput(input.PromptInput(err.Error()))
		}
		if action == input.ActionExit {
			return
		}
		switch action {
		case input.ActionNew:
			symbol := []rune(params[0])[0]
			x, _ := strconv.Atoi(params[1])
			y, _ := strconv.Atoi(params[2])
			height, _ := strconv.Atoi(params[3]) // alto
			width, _ := strconv.Atoi(params[4])  // ancho

			machine.Add(domain.NewRectangle(domain.Symbol(symbol), x, y, height, width))

		case input.ActionDelete:
			symbol := []rune(params[0])[0]
			machine.Delete(domain.Symbol(symbol))

		case input.ActionMove:
			symbol := []rune(params[0])[0]
			x, _ := strconv.Atoi(params[1])
			y, _ := strconv.Atoi(params[2])
			machine.Move(domain.Symbol(symbol), x, y)

		case input.ActionCombine:
			master := []rune(params[0])[0]
			slave := []rune(params[1])[0]

			machine.Combine(domain.Symbol(master), domain.Symbol(slave))

		}

		machine.Render()
		action, params, err = input.ParseInput(input.PromptInput(">"))
	}

}
