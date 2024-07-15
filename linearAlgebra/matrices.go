package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func DeclareMatrix(columns int, rows int, randomValue bool) [][]int {
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

func DeclareVector(s int, randomValue bool) []int {
	v := make([]int, s)
	if randomValue {
		for i := range v {
			rand := rand.Intn(4) + 1
			v[i] = rand
		}
	}
	return v
}

func MultiplyingMatrixWithVector(m [][]int, v []int) ([]int, error) {
	if len(m) == 0 || len(m[0]) != len(v) {
		return nil, fmt.Errorf("incompatible dimensions between matrix and vector")
	}
	valuesM := make([]int, len(m))
	for i := range m {
		sum := 0
		for j := range m[i] {
			fmt.Println(m[i][j], "*", v[j], "= ", m[i][j]*v[i])
			sum += m[i][j] * v[j]

		}
		valuesM[i] = sum
	}
	fmt.Println(valuesM)
	return valuesM, nil

}

func ScalarMatrixInt(m [][]int, s int) ([][]int, error) {
	for i := range m {
		for j := range m[i] {
			m[i][j] = m[i][j] * s
		}
	}
	return m, nil
}

func ScalarMatrixFloat(m [][]int, s float32) ([][]float32, error) {
	result := make([][]float32, len(m))

	for i := range m {
		result[i] = make([]float32, len(m[i]))
		for j := range m[i] {
			result[i][j] = float32(m[i][j]) * s
		}
	}

	return result, nil
}

func InverseOfMatrix(m [][]int) ([][]float32, error) {
	if len(m) == 2 && len(m[0]) == 2 {
		var scalar float32
		a, b, c, d := m[0][0], m[0][1], m[1][0], m[1][1]

		det := (d*a - b*c)
		if det != 0 {
			scalar = 1 / float32(det)
		}
		nm := m
		nm[0][0], nm[0][1], nm[1][0], nm[1][1] = d, -c, -b, a
		return ScalarMatrixFloat(nm, scalar)
	}
	return make([][]float32, 0), nil
}

func GaussianEliminationInt(m [][]int, b []int) [][]int {
	fmt.Println("Matriz antes da eliminação Gaussiana: ", m)
	for k := 0; k < len(m)-1; k++ {
		for i := k + 1; i < len(m); i++ {
			pivot := m[k][k]
			eliminated := m[i][k]
			ratio := eliminated / pivot // transformar o PIVOT em 1 e os elementos abaixo dele em 0 :D
			for j := k; j < len(m); j++ {
				m[i][j] = m[i][j] - ratio*m[k][j]
			}
			b[i] = b[i] - ratio*b[k] // alterando as constantes para resolução do sistema de equação :D
		}
	}

	fmt.Println("Matriz após eliminação Gaussiana: ", m)
	return m
}

func MultiplyingMatrixWithMatrix(A [][]int, B [][]int) ([][]int, error) {
	if len(A[0]) != len(B) {
		return make([][]int, 0), errors.New("Columns of Matrix doesnt match with Rows of other")
	}
	final_m := make([][]int, 0)
	for i := range final_m {
		final_m[i] = make([]int, len(B[0]))
	}
	for i := 0; i < len(B[0]); i++ { //COLUNAS
		vec := make([]int, len(A))
		for j := range A {
			vec[j] = B[j][i]
		}
		res, err := MultiplyingMatrixWithVector(A, vec)

		if err != nil {
			fmt.Println(err)
		} else {

			final_m = append(final_m, res)
		}

	}

	return final_m, nil
}

func main() {
	matrix := DeclareMatrix(2, 2, true)
	// vector := DeclareVector(2, true)
	A := [][]int{{1, 2}, {3, -1}}
	b := []int{5, 2}

	fmt.Println("Matrix")
	fmt.Println(matrix)
	fmt.Println("Vector")
	fmt.Println("Multiplying by a Vector at the same size ")
	fmt.Println(len(matrix))
	inverse, err := InverseOfMatrix(matrix)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("O inverso da matrix é: ", inverse)
	}

	GaussianEliminationInt(A, b)
	Ma := [][]int{{4, 1}, {5, 2}}
	Mb := [][]int{{0, 1, 3}, {-2, 2, 0}}
	resultMxM, err := MultiplyingMatrixWithMatrix(Ma, Mb)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultMxM)
	}
}
