package engine

import "fmt"

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
		return fmt.Sprintf("w%s", getPieceShorthand(piece))
	}
	return fmt.Sprintf("b%s", getPieceShorthand(piece))
}

func getPiece(row, col int) *Piece {
	if row == 1 {
		return &Piece{
			Type:  Pawn,
			Color: White,
		}
	}
	if row == 6 {
		return &Piece{
			Type:  Pawn,
			Color: White,
		}
	}

	if row == 0 || row == 7 {
		var color Color
		if row == 0 {
			color = White
		} else {
			color = Black
		}

		switch col {
		case 0, 7:
			return &Piece{
				Type:  Rook,
				Color: color,
			}
		case 1, 6:
			return &Piece{
				Type:  Knight,
				Color: color,
			}
		case 2, 5:
			return &Piece{
				Type:  Bishop,
				Color: color,
			}
		case 3:
			return &Piece{
				Type:  Queen,
				Color: color,
			}
		case 4:
			return &Piece{
				Type:  King,
				Color: color,
			}
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
				isLight:  (i+j)%2 == 0,
				Position: position,
				Piece:    getPiece(i, j),
			}
		}
	}

	return board
}
