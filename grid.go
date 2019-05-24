package main

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

func generateGrid(gs gridSize, mineCoords []coord) [][]string {
	if gs.rows == 0 || gs.cols == 0 {
		return [][]string{{""}}
	}

	results := make([][]string, gs.rows)
	for i := range results {
		for j := 0; j < gs.cols; j++ {
			results[i] = append(results[i], ".")
		}
	}

	populateResults(mineCoords, results)

	return results
}

func populateResults(mineCoords []coord, results [][]string) {
	if len(mineCoords) == 0 {
		return
	}

	for _, mine := range mineCoords {
		// Populate mine
		results[mine.row][mine.col] = "*"
		// Process
		populateResultsRow(mine.row-1, mine.col, results)
		populateResultsRow(mine.row, mine.col, results)
		populateResultsRow(mine.row+1, mine.col, results)
	}
}

func populateResultsRow(row int, mineCoordCol int, results [][]string) {
	// out of range check
	if row < 0 || row >= len(results) {
		return
	}
	colSize := len(results[0])
	for colOfRow := mineCoordCol - 1; colOfRow <= mineCoordCol+1; colOfRow++ {
		if colOfRow < 0 || colOfRow >= colSize {
			continue
		}
		results[row][colOfRow] = processCell(results[row][colOfRow])
	}
}

func processCell(cellValue string) string {
	switch cellValue {
	case "*":
		return cellValue
	case ".":
		return "1"
	default:
		// convert to int, add one, convert to string
		cellInt, err := strconv.Atoi(cellValue)
		if err != nil {
			log.Fatal(err)
		}
		cellInt++
		cellOpt := strconv.Itoa(cellInt)
		return cellOpt
	}
}
