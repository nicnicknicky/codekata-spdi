package main

import (
	"reflect"
	"testing"
)

func Test_generateGrid(t *testing.T) {
	type args struct {
		gridSize   gridSize
		mineCoords []coord
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Empty grid size",
			args: args{
				gridSize:   gridSize{rows: 0, cols: 0},
				mineCoords: nil,
			},
			want: [][]string{{""}},
		},
		{
			name: "Empty mine list",
			args: args{
				gridSize:   gridSize{rows: 4, cols: 3},
				mineCoords: []coord{},
			},
			want: [][]string{
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
		},
		{
			name: "Grid size rows 2 cols 3 with one mine",
			args: args{
				gridSize: gridSize{rows: 2, cols: 3},
				mineCoords: []coord{
					coord{row: 0, col: 1},
				},
			},
			want: [][]string{
				{"1", "*", "1"},
				{"1", "1", "1"},
			},
		},
		{
			name: "Grid size rows 3 cols 5 with three mines",
			args: args{
				gridSize: gridSize{rows: 3, cols: 5},
				mineCoords: []coord{
					coord{row: 0, col: 0},
					coord{row: 0, col: 1},
					coord{row: 2, col: 1},
				},
			},
			want: [][]string{
				{"*", "*", "1", ".", "."},
				{"3", "3", "2", ".", "."},
				{"1", "*", "1", ".", "."},
			},
		},
		{
			name: "Grid size rows 6 cols 12 with one mine",
			args: args{
				gridSize: gridSize{rows: 6, cols: 12},
				mineCoords: []coord{
					coord{row: 3, col: 4},
				},
			},
			want: [][]string{
				{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", "1", "1", "1", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", "1", "*", "1", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", "1", "1", "1", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			},
		},
		{
			name: "Grid size rows 6 cols 12 with three mines",
			args: args{
				gridSize: gridSize{rows: 6, cols: 12},
				mineCoords: []coord{
					coord{row: 2, col: 6},
					coord{row: 3, col: 4},
					coord{row: 4, col: 6},
				},
			},
			want: [][]string{
				{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", "1", "1", "1", ".", ".", ".", "."},
				{".", ".", ".", "1", "1", "2", "*", "1", ".", ".", ".", "."},
				{".", ".", ".", "1", "*", "3", "2", "2", ".", ".", ".", "."},
				{".", ".", ".", "1", "1", "2", "*", "1", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", "1", "1", "1", ".", ".", ".", "."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateGrid(tt.args.gridSize, tt.args.mineCoords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
