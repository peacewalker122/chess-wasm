package engine

type Color int

const (
	White Color = iota
	Black
)

type PieceType int

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Type  PieceType
	Color Color
}

type Move struct {
	Row       int
	Col       int
	AbleToEat bool
}

type Board struct {
	Row      int
	Col      int
	isLight  bool
	Position string
	Piece    *Piece
}
