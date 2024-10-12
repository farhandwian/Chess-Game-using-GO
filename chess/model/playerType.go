package model

type PlayerType int

const (
	PLAYER1 PlayerType = iota
	PLAYER2
)

// String method to print the PlayerType as a string
func (p PlayerType) String() string {
	return [...]string{"PLAYER1", "PLAYER2"}[p]
}
