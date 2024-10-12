package move

import (
	"chess/conditions"
	"chess/model"
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
func (p *PossibleMovesProvider) PossibleMoves(piece *model.Piece, board *model.Board, additionalBlockers []conditions.PieceCellOccupyBlocker, player *model.Player) []*model.Cell {
	if p.baseCondition.IsBaseConditionFulfilled(piece) {
		return p.PossibleMovesAsPerCurrentType(piece, board, additionalBlockers, player)
	}
	return nil
}

// PossibleMovesAsPerCurrentType harus diimplementasikan oleh tipe move tertentu
func (p *PossibleMovesProvider) PossibleMovesAsPerCurrentType(piece *Piece, board *Board, additionalBlockers []PieceCellOccupyBlocker, player *Player) []*Cell {
	return nil
}

// Helper method untuk mendapatkan semua langkah yang mungkin
func (p *PossibleMovesProvider) FindAllNextMoves(piece *Piece, nextCellProvider NextCellProvider, board *Board, cellOccupyBlockers []PieceCellOccupyBlocker, player *Player) []*Cell {
	result := []*Cell{}
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
func (p *PossibleMovesProvider) CheckIfCellCanBeOccupied(piece *Piece, cell *Cell, board *Board, additionalBlockers []PieceCellOccupyBlocker, player *Player) bool {
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
