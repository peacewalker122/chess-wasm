package engine

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

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

func getPiece(iswhitefirst bool, row, col int) (*string, PieceType) {
	// Determine piece positions based on whether white goes first
	var whitePawnRow, blackPawnRow int
	var whiteBackRow, blackBackRow int

	if iswhitefirst {
		// Standard chess setup (white on bottom)
		whitePawnRow = 6 // Second row from bottom
		whiteBackRow = 7 // Bottom row
		blackPawnRow = 1 // Second row from top
		blackBackRow = 0 // Top row
	} else {
		// Black goes first setup (black on bottom)
		blackPawnRow = 6 // Second row from bottom
		blackBackRow = 7 // Bottom row
		whitePawnRow = 1 // Second row from top
		whiteBackRow = 0 // Top row
	}

	// Set up pawns
	if row == whitePawnRow {
		val := getIcon(true, Pawn)
		return &val, Pawn
	}
	if row == blackPawnRow {
		val := getIcon(false, Pawn)
		return &val, Pawn
	}

	// Set up back rows (rooks, knights, bishops, etc.)
	switch row {
	case whiteBackRow:
		// White pieces
		switch col {
		case 0, 7:
			val := getIcon(true, Rook)
			return &val, Rook
		case 1, 6:
			val := getIcon(true, Knight)
			return &val, Knight
		case 2, 5:
			val := getIcon(true, Bishop)
			return &val, Bishop
		case 3:
			val := getIcon(true, Queen)
			return &val, Queen
		case 4:
			val := getIcon(true, King)
			return &val, King
		}
	case blackBackRow:
		// Black pieces
		switch col {
		case 0, 7:
			val := getIcon(false, Rook)
			return &val, Rook
		case 1, 6:
			val := getIcon(false, Knight)
			return &val, Knight
		case 2, 5:
			val := getIcon(false, Bishop)
			return &val, Bishop
		case 3:
			val := getIcon(false, Queen)
			return &val, Queen
		case 4:
			val := getIcon(false, King)
			return &val, King
		}
	}

	return nil, ""
}

func CreateBoard(isWhiteFirst *bool) [64]Board {
	if isWhiteFirst == nil {
		defaultValue := true
		isWhiteFirst = &defaultValue
	}

	board := [64]Board{}

	for i := 7; i >= 0; i-- {
		for j := 7; j >= 0; j-- {
			// Chess notation is column (a-h) + row (1-8)
			// i is row (0-7), j is column (0-7)
			position := fmt.Sprintf("%c%d", 'a'+j, 8-i)
			piece, piecetype := getPiece(*isWhiteFirst, i, j)
			board[i*8+j] = Board{
				Row:       i,
				Col:       j,
				IsLight:   (i+j)%2 == 0,
				Position:  position,
				Piece:     piece,
				PieceType: piecetype,
			}
		}
	}

	return board
}

// from a2 to a4 for example
func (b *Game) CreateMove(from string, to string) ([64]Board, error) {
	log.Debug().Str("from", from).Str("to", to).Msg("CreateMove")

	// legalmoves, err := b.GetLegalMove(from)
	// if err != nil {
	// 	return [64]Board{}, err
	// }
	// log.Debug().Strs("legalmoves", legalmoves).Send()

	// legal := false
	// for _, move := range legalmoves {
	// 	log.Debug().Str("move", move).Msg("CheckingMoves")
	// 	if move == to {
	// 		legal = true
	// 	}
	// }

	// if !legal {
	// 	log.Error().Msg("Illegal Moves")
	// 	return [64]Board{}, errors.New("illegal move")
	// }

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
	previousindex := row*8 + col

	b.Moves = append(b.Moves, Move{
		From:    from,
		To:      to,
		IsToEat: b.Board[targetindex].Piece != nil,
	})

	b.Board[targetindex].Piece = b.Board[previousindex].Piece
	b.Board[previousindex].Piece = nil

	return b.Board, nil
}

func (g *Game) GetLegalMove(from string) ([]string, error) {
	row, col := GetIndexFromCoordinate(from)
	index := row*8 + col

	piecetype := g.Board[index].PieceType

	log.Debug().Str("piecetype", string(piecetype)).Int("row", row).Int("col", col).Str("from", from).Send()

	switch piecetype {
	case Pawn:
		// Get piece information
		fromIndex := row*8 + col
		isPieceWhite := false
		if g.Board[fromIndex].Piece != nil {
			isPieceWhite = string((*g.Board[fromIndex].Piece)[0]) == "w"
		}

		positions := []string{}

		// Different movement directions based on color
		if isPieceWhite {
			// White pawn moves up (decreasing row number)
			row, err := strconv.Atoi(string(from[1]))
			if err != nil {
				return nil, err
			}

			// Forward movement (1 square)
			if row > 0 {
				forwardIdx := (row-1)*8 + col
				if g.Board[forwardIdx].Piece == nil {
					positions = append(positions, fmt.Sprintf("%c%d", 'a'+col, row))

					// Two squares from starting position
					if row == 6 && row > 1 {
						doubleForwardIdx := (row-2)*8 + col
						if g.Board[doubleForwardIdx].Piece == nil {
							positions = append(positions, fmt.Sprintf("%c%d", 'a'+col, row-1))
						}
					}
				}
			}

			// Capturing moves
			if row > 0 {
				// Left diagonal capture
				if col > 0 {
					leftCapIdx := (row-1)*8 + (col - 1)
					if g.Board[leftCapIdx].Piece != nil && string((*g.Board[leftCapIdx].Piece)[0]) == "b" {
						positions = append(positions, fmt.Sprintf("%c%d", 'a'+(col-1), row))
					}
				}

				// Right diagonal capture
				if col < 7 {
					rightCapIdx := (row-1)*8 + (col + 1)
					if g.Board[rightCapIdx].Piece != nil && string((*g.Board[rightCapIdx].Piece)[0]) == "b" {
						positions = append(positions, fmt.Sprintf("%c%d", 'a'+(col+1), row))
					}
				}
			}

			// TODO: En passant rule

		} else {
			// Black pawn moves down (increasing row number)

			// Forward movement (1 square)
			if row < 7 {
				forwardIdx := (row+1)*8 + col
				if g.Board[forwardIdx].Piece == nil {
					positions = append(positions, fmt.Sprintf("%c%d", 'a'+col, row+2))

					// Two squares from starting position
					if row == 1 && row < 6 {
						doubleForwardIdx := (row+2)*8 + col
						if g.Board[doubleForwardIdx].Piece == nil {
							positions = append(positions, fmt.Sprintf("%c%d", 'a'+col, row+3))
						}
					}
				}
			}

			// Capturing moves
			if row < 7 {
				// Left diagonal capture
				if col > 0 {
					leftCapIdx := (row+1)*8 + (col - 1)
					if g.Board[leftCapIdx].Piece != nil && string((*g.Board[leftCapIdx].Piece)[0]) == "w" {
						positions = append(positions, fmt.Sprintf("%c%d", 'a'+(col-1), row+2))
					}
				}

				// Right diagonal capture
				if col < 7 {
					rightCapIdx := (row+1)*8 + (col + 1)
					if g.Board[rightCapIdx].Piece != nil && string((*g.Board[rightCapIdx].Piece)[0]) == "w" {
						positions = append(positions, fmt.Sprintf("%c%d", 'a'+(col+1), row+2))
					}
				}
			}

			// TODO: En passant rule
		}

		// Debug output
		log.Debug().Strs("legal_moves", positions).Str("piece", string(*g.Board[fromIndex].Piece)).Send()

		return positions, nil
	}

	return []string{}, nil
}
