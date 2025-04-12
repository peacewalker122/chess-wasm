package engine

import "fmt"

type Game struct {
	Moves []Move
	Board [64]Board
}

func getPieceShorthand(piece PieceType) string {
	switch piece {
	case Pawn:
		return "P"
	case Knight:
		return "N"
	case Bishop:
		return "B"
	case Rook:
		return "R"
	case Queen:
		return "Q"
	case King:
		return "K"
	default:
		return ""
	}
}

func getIcon(isLight bool, piece PieceType) string {
	if isLight {
		return fmt.Sprintf("w%s.svg", piece)
	}
	return fmt.Sprintf("b%s.svg", piece)
}

func getPiece(row, col int) *string {
	if row == 1 {
		val := getIcon(true, Pawn)
		return &val
	}
	if row == 6 {
		val := getIcon(false, Pawn)
		return &val
	}

	if row == 0 || row == 7 {
		var isLight bool
		if row == 0 {
			isLight = true
		} else {
			isLight = false
		}

		switch col {
		case 0, 7:
			val := getIcon(isLight, Rook)
			return &val
		case 1, 6:
			val := getIcon(isLight, Knight)
			return &val
		case 2, 5:
			val := getIcon(isLight, Bishop)
			return &val
		case 3:
			val := getIcon(isLight, Queen)
			return &val
		case 4:
			val := getIcon(isLight, King)
			return &val
		default:
			return nil
		}
	}
	return nil
}

func CreateBoard(isWhiteFirst *bool) [64]Board {
	if isWhiteFirst == nil {
		defaultValue := true
		isWhiteFirst = &defaultValue
	}

	board := [64]Board{}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			// Chess notation is column (a-h) + row (1-8)
			// i is row (0-7), j is column (0-7)
			position := fmt.Sprintf("%c%d", 'a'+j, 8-i)
			board[i*8+j] = Board{
				Row:      i,
				Col:      j,
				IsLight:  (i+j)%2 == 0,
				Position: position,
				Piece:    getPiece(i, j),
			}
		}
	}

	return board
}

// from a2 to a4 for example
func (b *Game) CreateMove(from string, to string) ([64]Board, error) {
	// TODO: add move validation

	row, col := GetIndexFromCoordinate(from)
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return [64]Board{}, fmt.Errorf("invalid from coordinate: %s", from)
	}
	if b.Board[row*8+col].Piece == nil {
		return [64]Board{}, fmt.Errorf("no piece at from coordinate: %s", from)
	}

	row2, col2 := GetIndexFromCoordinate(to)
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return [64]Board{}, fmt.Errorf("invalid from coordinate: %s", from)
	}

	targetindex := row2*8 + col2
	previousindex := row*8 + row

	b.Moves = append(b.Moves, Move{
		From:    from,
		To:      to,
		IsToEat: b.Board[targetindex].Piece != nil,
	})

	b.Board[targetindex].Piece = b.Board[previousindex].Piece
	b.Board[previousindex].Piece = nil

	return b.Board, nil
}
