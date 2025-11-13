package dynprog

import "fmt"

func BellmanEquationPrint(table [][]int) {
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
		for l = 0; l < rows; l++ {
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
