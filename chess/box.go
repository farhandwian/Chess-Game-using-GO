package chess

type Box struct {
	x     int
	y     int
	piece *Piece
}

func NewBox(x int, y int, piece *Piece) *Box {
	return &Box{x: x, y: y, piece: piece}
}

func (b *Box) getX() int {
	return b.x
}

func (b *Box) getY() int {
	return b.y
}

func (b *Box) getPiece() *Piece {
	return b.piece
}

func (b *Box) isFree() bool {
	if b.piece == nil {
		return true
	}
	return false
}
