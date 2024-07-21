package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	TextNew    = "new:\tsymbol x y height width\t(char, int, int, int, int space separated)\n"
	TextDelete = "delete:\tsymbol\t(char)\n"
	MoveText   = "move:\tsymbol deltaX deltaY\t(char, int, int, space separated)\n"
)

func PromptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ParseInput(input string) (a Action, p Params, err error) {
	inputs := strings.Split(input, ` `)
	p = inputs[1:]

	switch inputs[0] {
	case "new":
		if len(p) != 5 {
			err = fmt.Errorf("some param missing\nhelp\n%s", TextNew)
			return
		}
		for i := 1; i < len(p); i++ {
			_, err = strconv.Atoi(p[i])
			if err != nil {
				return
			}
		}
		a = ActionNew
		return
	case "delete":
		if len(p) != 1 {
			err = fmt.Errorf("some param missing\nhelp\n%s", TextDelete)
			return
		}
		a = ActionDelete
		return
	case "move":
		if len(p) != 3 {
			err = fmt.Errorf("some param missing\nhelp\n%s", MoveText)
			return
		}
		for i := 1; i < len(p); i++ {
			_, err = strconv.Atoi(p[i])
			if err != nil {
				return
			}
		}
		a = ActionMove
		return
	case "exit":
		a = ActionExit
		return
	default:
		err = fmt.Errorf("invalid command '%s'", inputs[0])
	}
	return
}

func Help() string {
	return "command list:\n" +
		TextNew +
		TextDelete +
		MoveText +
		"exit\n>"
}
