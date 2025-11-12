package matrix

import (
	"fmt"

	"github.com/arsagyr/ra/fraction"
)

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type MatrixInterface interface {
	Determinant()
}

type Matrix[T Number] struct {
	ColSize, StrSize uint
	Cells            [][]T
}

type FractionMatrix struct {
	ColSize, StrSize uint
	Cells            [][]fraction.Fraction
}

func InitMatrix[T Number](m, n uint) Matrix[T] {
	var c [][]T = make([][]T, m, n)
	var i, j uint
	matrix := make([][]T, n)
	for k := range matrix {
		matrix[k] = make([]T, m)
	}
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			matrix[i][j] = 0
		}
	}
	return Matrix[T]{
		ColSize: m,
		StrSize: n,
		Cells:   c,
	}
}

func MakeMatrix[T Number](m, n uint, cells [][]T) Matrix[T] {
	return Matrix[T]{
		ColSize: m,
		StrSize: n,
		Cells:   cells,
	}
}
func MakeMatrixBySlice[T Number](cells [][]T) Matrix[T] {
	return Matrix[T]{
		ColSize: uint(len(cells[0])),
		StrSize: uint(len(cells)),
		Cells:   cells,
	}
}

func (matrix Matrix[T]) Determinant() float64 {
	var d, sign float64 = 0, 1
	var minor Matrix[T]
	var i, j, ii, ij uint

	if matrix.ColSize == matrix.StrSize && matrix.ColSize != 0 && matrix.StrSize != 0 {
		if matrix.ColSize == 1 {
			d = float64(matrix.Cells[0][0])
		} else {
			if matrix.ColSize == 2 {
				d = float64(matrix.Cells[0][0]*matrix.Cells[1][1] - matrix.Cells[1][0]*matrix.Cells[0][1])
			} else {
				if matrix.ColSize == 3 {
					for i = 0; i < matrix.StrSize; i++ {
						for j = 0; j < matrix.ColSize; j++ {
							d = float64(
								matrix.Cells[0][0]*matrix.Cells[1][1]*matrix.Cells[2][2] +
									matrix.Cells[0][2]*matrix.Cells[1][0]*matrix.Cells[2][1] +
									matrix.Cells[0][1]*matrix.Cells[1][2]*matrix.Cells[2][0] -
									matrix.Cells[0][2]*matrix.Cells[1][1]*matrix.Cells[2][0] -
									matrix.Cells[1][0]*matrix.Cells[0][1]*matrix.Cells[2][2] -
									matrix.Cells[0][0]*matrix.Cells[2][1]*matrix.Cells[1][2],
							)
						}
					}
				} else {
					for i = 0; i < matrix.StrSize; i++ {
						for j = 0; j < matrix.ColSize; j++ {

							cells := make([][]T, matrix.StrSize-1)
							for k := range cells {
								cells[k] = make([]T, matrix.ColSize-1)
							}
							fmt.Println("New minor")
							fmt.Println(i, j)
							for ii = 0; ii < matrix.StrSize-1; ii++ {
								for ij = 0; ij < matrix.ColSize-1; ij++ {
									if ii < i || ij < j {
										cells[ii][ij] = matrix.Cells[ii][ij]
									} else {
										cells[ii][ij] = matrix.Cells[ii+1][ij+1]
									}
								}
							}

							minor = MakeMatrix(matrix.ColSize-1, matrix.StrSize-1, cells)
							d = d + sign*float64(matrix.Cells[i][j])*minor.Determinant()

							sign *= -1
						}
					}
				}
			}
		}
	}

	return d
}

func (matrix FractionMatrix) Determinant() fraction.Fraction {
	var d fraction.Fraction = fraction.NewFraction(1, 1)
	var i, j uint
	for i = 0; i < matrix.StrSize; i++ {
		for j = 0; j < matrix.ColSize; j++ {
			d = matrix.Cells[i][j]
		}
	}
	return d
}

// func (m Matrix[T]) Determinant() T {
// 	var d T
// 	switch a := T.(type) {
// 	case uint:
// 	case int:
// 	default:
// 		log.Panic("Error: wrong format")
// 	}
// 	return d
// }

// switch a := ().(type) {
// case uint:
// 	d = float64(a)
// case int:
// 	d = float64(a)
// case float32:
// 	d = float64(a)
// case float64:
// 	d = a
// default:
// 	log.Panic("Error: wrong format")
// }
