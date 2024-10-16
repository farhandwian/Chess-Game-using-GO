package move

import (
	"chess_game/chess"
	"chess_game/chess/conditions"
)

type PossibleMovesProvider struct {
	maxSteps             int
	baseCondition        conditions.MoveBaseCondition
	moveFurtherCondition conditions.PieceMoveFurtherCondition
	baseBlocker          conditions.PieceCellOccupyBlocker
}

// NewPossibleMovesProvider is the constructor for PossibleMovesProvider
func NewPossibleMovesProvider(maxSteps int, baseCondition conditions.MoveBaseCondition,
	moveFurtherCondition conditions.PieceMoveFurtherCondition, baseBlocker conditions.PieceCellOccupyBlocker) *PossibleMovesProvider {
	return &PossibleMovesProvider{
		maxSteps:             maxSteps,
		baseCondition:        baseCondition,
		moveFurtherCondition: moveFurtherCondition,
		baseBlocker:          baseBlocker,
	}
}

// PossibleMoves returns all possible cells for a given piece and board
func (p *PossibleMovesProvider) PossibleMoves(piece *chess.Piece, board *chess.Board, additionalBlockers []conditions.PieceCellOccupyBlocker, player *chess.Player) []*chess.Box {
	if p.baseCondition.IsBaseConditionFulfilled(piece) {
		return p.PossibleMovesAsPerCurrentType(piece, board, additionalBlockers, player)
	}
	return nil
}

// // PossibleMovesAsPerCurrentType harus diimplementasikan oleh tipe move tertentu
// func (p *PossibleMovesProvider) PossibleMovesAsPerCurrentType(piece *chess.Piece, board *chess.Board, additionalBlockers []PieceCellOccupyBlocker, player *chess.Player) []*chess.Box {
// 	return nil
// }

// Helper method untuk mendapatkan semua langkah yang mungkin
func (p *PossibleMovesProvider) FindAllNextMoves(piece *chess.Piece, nextCellProvider NextCellProvider, board *chess.Board, cellOccupyBlockers []PieceCellOccupyBlocker, player *chess.Player) []*chess.Box {
	result := []*chess.Box{}
	nextCell := nextCellProvider(piece.CurrentCell)
	numSteps := 1

	for nextCell != nil && numSteps <= p.MaxSteps {
		if p.CheckIfCellCanBeOccupied(piece, nextCell, board, cellOccupyBlockers, player) {
			result = append(result, nextCell)
		}
		if !p.MoveFurtherCondition.CanPieceMoveFurtherFromCell(piece, nextCell, board) {
			break
		}
		nextCell = nextCellProvider(nextCell)
		numSteps++
	}

	return result
}

// Mengecek apakah sebuah sel bisa ditempati oleh bidak
func (p *PossibleMovesProvider) CheckIfCellCanBeOccupied(piece *chess.Piece, cell *chess.Box, board *chess.Board, additionalBlockers []PieceCellOccupyBlocker, player *chess.Player) bool {
	if p.BaseBlocker != nil && p.BaseBlocker.IsCellNonOccupiableForPiece(cell, piece, board, player) {
		return false
	}

	for _, blocker := range additionalBlockers {
		if blocker.IsCellNonOccupiableForPiece(cell, piece, board, player) {
			return false
		}
	}

	return true
}
