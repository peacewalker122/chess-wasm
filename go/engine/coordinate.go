package engine

func GetIndexFromCoordinate(coordinate string) (row int, col int) {
	potetialcol := coordinate[0]
	potetialrow := coordinate[1]

	col = int(potetialcol - 'a')
	row = int('8' - potetialrow)

	return
}
