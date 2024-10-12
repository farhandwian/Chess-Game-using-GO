package model

type ColorType int

const (
	BLOCK ColorType = iota
	WHITE
)

// String method to print the ColorType as a string
func (c ColorType) String() string {
	return [...]string{"BLOCK", "WHITE"}[c]
}
