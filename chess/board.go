package chess

type Board struct {
	squares [8][8]Box
}

func initializePiece(rowStart int) {

}

func (b *Board) getLeftCell(box Box) *Box {
	return b.getBoxAtLocation(box.getX(), box.getY()-1)
}

func (b *Board) getRightCell(box Box) *Box {
	return b.getBoxAtLocation(box.getX(), box.getY()+1)
}

func (b *Board) getUpCell(box Box) *Box {
	return b.getBoxAtLocation(box.getX()+1, box.getY())
}

func (b *Board) getBottomCell(box Box) *Box {
	return b.getBoxAtLocation(box.getX()-1, box.getY())
}

func (b *Board) getBoxAtLocation(x, y int) *Box {
	if x < 0 || x >= 8 || y < 0 || y >= 8 {
		return nil
	}
	return &b.squares[x][y]
}
