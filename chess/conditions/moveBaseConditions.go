package conditions

type MoveBaseCondition interface {
	isBaseConditionFullfilled(Piece piece)
}
