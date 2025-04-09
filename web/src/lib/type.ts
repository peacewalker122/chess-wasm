export interface Board {
	row: number;
	col: number;
	isLight: boolean;
	position: string;
	piece: string | null; // e.g. "wP" for white pawn, "bK" for black king
}

export enum pieceType {
	Pawn = "P",
	Knight = "N",
	Bishop = "B",
	Rook = "R",
	Queen = "Q",
	King = "K",
}
