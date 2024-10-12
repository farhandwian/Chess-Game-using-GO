package model

// PieceType is an enumeration of chess pieces
type PieceType int

const (
	KING PieceType = iota
	QUEEN
	ROOK
	KNIGHT
	BISHOP
	PAWN
)

// String method to print the PieceType as a string
func (p PieceType) String() string {
	return [...]string{"KING", "QUEEN", "ROOK", "KNIGHT", "BISHOP", "PAWN"}[p]
}
