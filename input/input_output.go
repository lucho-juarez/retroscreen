package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ParseInput(input string) (a Action, p Params) {
	inputs := strings.Split(input, ` `)
	p = inputs[1:]

	switch inputs[0] {
	case "new":
		a = ActionNew
		return
	case "delete":
		a = ActionDelete
		return
	case "move":
		a = ActionMove
		return
	case "exit":
		a = ActionExit
		return
	default:
		panic("invalid input")
	}

}
