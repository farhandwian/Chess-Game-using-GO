package chess

import (
	"chess_game/chess/model"
)

type Piece struct {
	color     string
	pieceType model.PieceType
}

func NewPiece(color string, pieceType model.PieceType) *Piece {
	return &Piece{color: color, pieceType: pieceType}
}

func (p Piece) GetPossibleMove(board Board) Box {

}
