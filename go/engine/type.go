package engine

type Color int

const (
	White Color = iota
	Black
)

type PieceType string

const (
	Pawn   PieceType = "P"
	Knight           = "N"
	Bishop           = "B"
	Rook             = "R"
	Queen            = "Q"
	King             = "K"
)

type Piece struct {
	Type  PieceType
	Color Color
}

type Move struct {
	From    string
	To      string
	IsToEat bool
}

type Board struct {
	Row      int     `json:"row"`
	Col      int     `json:"col"`
	IsLight  bool    `json:"isLight"`
	Position string  `json:"position"`
	Piece    *string `json:"piece"`
}
