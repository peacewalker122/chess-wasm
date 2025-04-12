package engine

import "fmt"

func GetIndexFromCoordinate(coordinate string) (row int, col int) {
	potetialcol := coordinate[0]
	potetialrow := coordinate[1]

	fmt.Println("potetialcol", string(potetialcol))
	fmt.Println("potetialrow", string(potetialrow))

	col = int(potetialcol - 'a')
	row = int(potetialrow - '1')

	return
}
