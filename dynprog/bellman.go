package dynprog

import (
	"fmt"
)

func BellmanEquation(res int, table [][]int) (int, []int) {
	var c, i, j, l int

	rows := len(table)
	cols := len(table[0])
	ways := make([][]int, rows)
	profits := make([]int, rows)
	for i = 0; i < rows; i++ {
		profits[i] = table[i][1]
	}

	for i = 0; i < rows; i++ {
		ways[i] = make([]int, 1)
		ways[i][0] = i
	}
	for c = 1; c < cols-1; c++ {
		subprofits := make([]int, res+1)
		subways := make([][]int, res+1)
		for l = 0; l <= res; l++ {
			pairsum := table[0][c] + table[0][c+1]
			maxpairsum := pairsum
			maxpoints := []int{c - 1, l}
			for i = 0; i < len(profits); i++ {
				points := []int{}
				for j = 0; j < rows; j++ {
					if i+table[j][0] == l {
						pairsum = profits[i] + table[j][c+1]
						points = append(ways[i], j)
					}
				}
				if maxpairsum < pairsum {
					maxpairsum = pairsum
					maxpoints = points
				}
			}
			subprofits[l] = maxpairsum
			subways[l] = maxpoints
		}

		profits = subprofits
		ways = subways
	}

	maxprofitnum := 0
	for i = 1; i <= res; i++ {
		if profits[i] > profits[maxprofitnum] {
			maxprofitnum = i
		}
	}
	return profits[maxprofitnum], ways[maxprofitnum]

}

func BellmanEquationPrint(res int, table [][]int) {
	var c, i, j, l int

	rows := len(table)
	cols := len(table[0])
	ways := make([][]int, rows)
	profits := make([]int, rows)
	for i = 0; i < rows; i++ {
		profits[i] = table[i][1]
	}

	for i = 0; i < rows; i++ {
		ways[i] = make([]int, 1)
		ways[i][0] = i
	}
	for c = 1; c < cols-1; c++ {
		subprofits := make([]int, res+1)
		subways := make([][]int, res+1)
		for l = 0; l <= res; l++ {
			pairsum := table[0][c] + table[0][c+1]
			maxpairsum := pairsum
			maxpoints := []int{c - 1, l}
			for i = 0; i < len(profits); i++ {
				points := []int{}
				for j = 0; j < rows; j++ {
					if i+table[j][0] == l {
						pairsum = profits[i] + table[j][c+1]
						points = append(ways[i], j)
					}
				}
				if maxpairsum < pairsum {
					maxpairsum = pairsum
					maxpoints = points
				}
			}
			subprofits[l] = maxpairsum
			subways[l] = maxpoints
		}

		profits = subprofits
		ways = subways
	}

	maxprofitnum := 0
	for i = 1; i <= res; i++ {
		if profits[i] > profits[maxprofitnum] {
			maxprofitnum = i
		}
	}
	for step, invest := range ways[maxprofitnum] {
		fmt.Printf("Предприятию %d вложить %d \n", step+1, invest)
	}

	fmt.Printf("Максимальная прибыль инвестиций - %d\n", profits[maxprofitnum])
}
