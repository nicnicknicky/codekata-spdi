package grid

import (
	"log"
	"strconv"
)

type gridSize struct {
	rows int
	cols int
}

type coord struct {
	row int
	col int
}

func generateOptGrid(gs gridSize, mineCoords []coord) [][]string {
	if gs.rows == 0 || gs.cols == 0 {
		return [][]string{{""}}
	}

	optGrid := make([][]string, gs.rows)
	for i := range optGrid {
		for j := 0; j < gs.cols; j++ {
			optGrid[i] = append(optGrid[i], ".")
		}
	}

	populateResults(mineCoords, optGrid)

	return optGrid
}

func populateResults(mineCoords []coord, optGrid [][]string) {
	// no mines, done!
	if len(mineCoords) == 0 {
		return
	}
	// iterate through each mine
	for _, mine := range mineCoords {
		// insert mine in cell
		optGrid[mine.row][mine.col] = "*"
		// process by rows
		// increment cells around mine perimeter ( if permitted )
		// | +1 | +1 | +1 | prevRow
		// | +1 |  * | +1 | currentRow
		// | +1 | +1 | +1 | nextRow
		processCellsByRow(mine.row-1, mine.col, optGrid)
		processCellsByRow(mine.row, mine.col, optGrid)
		processCellsByRow(mine.row+1, mine.col, optGrid)
	}
}

func processCellsByRow(row int, mineCoordCol int, optGrid [][]string) {
	// row out of range check
	if row < 0 || row >= len(optGrid) {
		return
	}
	// get number of columns from first row which is guaranteed to exist minimally
	colSize := len(optGrid[0])
	// start from prevColumn, stop at nextColumn
	for colOfRow := mineCoordCol - 1; colOfRow <= mineCoordCol+1; colOfRow++ {
		// column out of range check
		if colOfRow < 0 || colOfRow >= colSize {
			continue
		}
		optGrid[row][colOfRow] = processCell(optGrid[row][colOfRow])
	}
}

func processCell(cellValue string) string {
	switch cellValue {
	case "*":
		// mine
		return cellValue
	case ".":
		// not been processed before
		// mine is 1 cell away from me
		return "1"
	default:
		// mine is 1 cell away from me
		cellInt, err := strconv.Atoi(cellValue)
		if err != nil {
			log.Fatal(err)
		}
		cellInt++
		cellOpt := strconv.Itoa(cellInt)
		return cellOpt
	}
}
