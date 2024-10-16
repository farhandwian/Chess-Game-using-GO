package move

type PossibleMovesProviderHorizontal struct {
	PossibleMovesProvider
}

func (p *PossibleMovesProviderHorizontal) PossibleMovesAsPerCurrentType(piece *Piece, board *Board, additionalBlockers []PieceCellOccupyBlocker, player *Player) []*Cell {
	var result []*Cell
	result = append(result, p.findAllNextMoves(piece, board.GetLeftCell, board, additionalBlockers, player)...)
	result = append(result, p.findAllNextMoves(piece, board.GetRightCell, board, additionalBlockers, player)...)
	return result
}
