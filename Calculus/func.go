package main

import (
	"fmt"
	"math"
)

const (
	width  = 160
	height = 10
)

func Plot(x, y []float64) {

	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	funPlot := func(xPos, yPos int) {
		if xPos >= 0 && xPos < width && yPos >= 0 && yPos < height {

			grid[height-1-yPos][xPos] = '&'
		}
	}

	minX, maxX := minMax(x)
	minY, maxY := minMax(y)

	for i := 0; i < len(x)-1; i++ {
		x0 := int((x[i] - minX) / (maxX - minX) * (width - 1))
		y0 := int((y[i] - minY) / (maxY - minY) * (height - 1))
		x1 := int((x[i+1] - minX) / (maxX - minX) * (width - 1))
		y1 := int((y[i+1] - minY) / (maxY - minY) * (height - 1))
		BresehamLineAlgortihm(x0, y0, x1, y1, funPlot)
	}

	for _, line := range grid {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func minMax(arr []float64) (min float64, max float64) {
	min, max = arr[0], arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return
}

func gMin(v []int) int {
	min := 0
	for i := range v {
		for j := range v[i] {
			if v[j] < min {
				min = v[j]
			}
		}

	}
	return min
}
func BresehamLineAlgortihm(x0, y0, x1, y1 int, plotPoint func(x, y int)) {
	dx := int(math.Abs(float64(x1 - x0)))
	dy := int(-math.Abs(float64(y1 - y0)))
	sx := 0
	sy := 0
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}

	if y0 < y1 {
		sy = 1.0
	} else {
		sy = -1.0
	}

	e := dx + dy
	for {
		plotPoint(x0, y0)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * e
		if e2 >= dy {
			if x0 == x1 {
				break
			}
			e = e + dy
			x0 = x0 + sx
		}
		if e2 <= dx {
			if y0 == y1 {
				break
			}
			e = e + dx
			y0 = y0 + sy
		}

	}

}

func CosF() ([]float64, []float64) {

	s := 90
	y := make([]float64, s)
	x := make([]float64, s)

	for i := 0; i < len(x); i++ {
		x[i] = float64(i)
	}
	for i := 0; i < len(y); i++ {
		y[i] = math.Cos(float64(i))
	}
	return x, y
}

func DerivativePowerRule(s string) string {

	formula := ""
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && i-1 > -1 {
			if s[i] == '^' {
				coefficient := s[i-1]
				expoent := s[i+1]
				intexpo := int(expoent) - 1
				formula += string(expoent) + string(coefficient) + "^" + string(intexpo)
			}
			validation := s[i] == '+' || s[i] == '-' || s[i] == ' ' || s[i] == '/'
			if validation {
				formula += string(s[i])
			}
		}
	}
	return formula

}

func main() {
	expression := "x^4"
	res := DerivativePowerRule(expression)
	fmt.Println("Derivative:" + res)
}
