package main

import (
	"fmt"
	"math/rand"
)

func declareMatrix(columns int, rows int, randomValue bool) [][]int {
	matrix := make([][]int, columns)
	for i := range matrix {
		matrix[i] = make([]int, rows)
	}

	if randomValue {
		for i := range matrix {

			for j := range matrix[i] {

				randNum := rand.Intn(5)
				matrix[i][j] = randNum
			}
		}
	}
	return matrix
}

func declareVector(s int, randomValue bool) []int {
	v := make([]int, s)
	if randomValue {
		for i := range v {
			rand := rand.Intn(4) + 1
			v[i] = rand
		}
	}
	return v
}

// scalar matrix
func multiplyingMatrixWithVector(m [][]int, v []int) ([]int, error) {
	sums := make([]int, len(v))
	for i := range m {
		for j := range m[i] {
			sums[i] = sums[i] + (m[i][j] * v[i])
		}
	}

	return sums, nil
}

func scalarMatrixInt(m [][]int, s int) ([][]int, error) {
	for i := range m {
		for j := range m[i] {
			m[i][j] = m[i][j] * s
		}
	}
	return m, nil
}

func scalarMatrixFloat(m [][]int, s float32) ([][]float32, error) {
	result := make([][]float32, len(m))

	for i := range m {
		result[i] = make([]float32, len(m[i]))
		for j := range m[i] {
			result[i][j] = float32(m[i][j]) * s
		}
	}

	return result, nil
}

func inverseOfMatrix(m [][]int) ([][]float32, error) {
	fmt.Println("Invertendo a matrix: ", m)
	if len(m) == 2 && len(m[0]) == 2 {
		var scalar float32
		a, b, c, d := m[0][0], m[0][1], m[1][0], m[1][1]

		det := (d*a - b*c)
		if det != 0 {
			scalar = 1 / float32(det)
		}
		nm := m
		nm[0][0], nm[0][1], nm[1][0], nm[1][1] = d, -c, -b, a
		fmt.Println("Escalar: ", scalar)
		fmt.Println("Matrix: ", nm)
		return scalarMatrixFloat(nm, scalar)
	}
	return make([][]float32, 0), nil
}
func main() {
	matrix := declareMatrix(2, 2, true)
	vector := declareVector(2, true)

	fmt.Println("Matrix")
	fmt.Println(matrix)
	fmt.Println("Vector")
	fmt.Println(vector)
	fmt.Println("Multiplying by a Vector at the same size ")
	fmt.Println(len(matrix))
	inverse, err := inverseOfMatrix(matrix)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("O inverso da matrix é: ", inverse)
	}
}
