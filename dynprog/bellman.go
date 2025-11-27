package dynprog

import (
	"fmt"
)

func BellmanEquationPrint(res int, table [][]int) {
	var i, j, k int
	var l int

	rows := len(table)
	subtable := make([]int, rows)
	for i = 0; i < rows; i++ {
		subtable[i] = table[i][1]
	}

	maxkn := 0
	maxjn := 0
	step := 0
	finalstep := 0
	for i = 1; i < len(table[0])-1; i++ {
		subsubtable := make([]int, rows)
		maxk := table[0][i] + table[0][i+1]
		maxj := maxk
		step = maxjn
		for l = 0; l < res; l++ {
			for j = 0; j < rows; j++ {
				for k = 0; k < rows; k++ {
					if l == (table[j][0] + table[k][0]) {
						if subtable[j]+table[k][i+1] > maxk {
							maxk = subtable[j] + table[k][i+1]
							maxkn = j
						}
					}
				}
				subsubtable[l] = maxk
				if maxj < maxk {
					maxj = maxk
					maxjn = maxkn
				}
			}

		}
		subtable = subsubtable
		finalstep += (maxjn - step)
		fmt.Printf("x%d = %d\n", i, maxjn-step) //Вывод шагов инвестиций
	}

	fmt.Printf("x%d = %d\n", i, rows-finalstep-1) //Вывод последнего шага

	maxk := subtable[0]
	for i = 1; i < rows; i++ {
		if subtable[i] > maxk {
			maxk = subtable[i]
		}
	}

	fmt.Printf("Максимальная прибыль инвестиций - %d\n", maxk)
}

func BellmanEquation(res int, table [][]int) {
	var i, j, k, l int

	rows := len(table)
	cols := len(table[0])
	fmt.Printf("rows = %d\ncols = %d\n", rows, cols)
	steps := make([]int, cols)
	profits := make([]int, rows)
	profits[0] = table[0][1]

	var finalstep, maxkprice, maxjprice int
	for i = 1; i < cols-1; i++ {
		for l = 0; l < res; l++ {
			maxj := table[0][i] + table[0][i+1]
			maxk := table[0][i] + table[0][i+1]
			for j = 0; j < rows; j++ {
				for k = 0; k < rows; k++ {
					if l == (table[j][0] + table[k][0]) {
						if table[j][i]+table[k][i+1] > maxk {
							maxk = profits[i-1] + table[k][i+1]
							maxkprice = table[k][0]
							// fmt.Printf("profits[%d] + table[%d][%d] = %d + %d = %d\n", i, k, i+1, profits[i-1], table[k][i+1], maxk)
						}
					}
				}

				if maxj < maxk {
					maxj = maxk
					maxjprice = maxkprice
				}
			}
			profits[i] = maxj
			steps[i] = maxjprice
		}
		fmt.Printf("step%d = %d\n", i, steps[i])
	}
	finalstep += maxjprice
	fmt.Printf("x%d = %d\n", i, res-finalstep) //Вывод последнего шага

	maxprofit := profits[0]
	for i = 1; i < rows; i++ {
		fmt.Println(profits[i])
		if profits[i] > maxprofit {
			maxprofit = profits[i]
		}
	}

	fmt.Printf("Максимальная прибыль инвестиций - %d\n", maxprofit)
}

func BellmanEquationVer2(res int, table [][]int) {
	var c, i, j, l int
	var pairnum int

	rows := len(table)
	cols := len(table[0])
	// steps := make([]int, cols)
	profits := make([]int, rows)
	for i = 0; i < rows; i++ {
		profits[i] = table[i][1]
	}
	maxpairnum := 0
	for c = 1; c < cols-1; c++ {
		subprofits := make([]int, res+1)
		for l = 0; l <= res; l++ {
			pairsum := table[0][c] + table[0][c+1]
			maxpairsum := pairsum
			for i = 0; i < len(profits); i++ {
				for j = 0; j < rows; j++ {
					if i+table[j][0] == l {
						pairsum = profits[i] + table[j][c+1]
						pairnum = j
					}

				}
				if maxpairsum < pairsum {
					maxpairsum = pairsum
					maxpairnum = pairnum
				}

			}
			subprofits[l] = maxpairsum
		}

		// steps[c] = l - maxpairnum
		profits = subprofits
	}

	maxprofit := profits[0]
	for i = 1; i <= res; i++ {
		if profits[i] > maxprofit {
			maxprofit = profits[i]
		}
	}

	fmt.Printf("Максимальная прибыль инвестиций - %d\n", maxprofit)
}
