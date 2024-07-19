package input

type Action int
type Params []string

const (
	ActionNew Action = iota
	ActionDelete
	ActionMove
	ActionExit
)
