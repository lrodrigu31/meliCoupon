package failures

import (
	"fmt"
	"math"
	"sort"
)

// CrearMatriz second intent
func CrearMatriz() {
	prices := []int{
		0, 100, 210, 260, 80, 90,
	}
	defaultMaxQuantities := []int{
		0, 1, 1, 1, 1, 1,
	}
	amount := 500

	rows := len(prices)
	columns := amount + 1

	var intermediateResult [][]int

	for i := 0; i < rows; i++ {
		var vectorFila []int
		for j := 0; j < columns; j++ {
			vectorFila = append(vectorFila, 0)
		}
		intermediateResult = append(intermediateResult, vectorFila)
	}

	//fmt.Println("intermediateResult", intermediateResult)

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= amount; j++ {
			if i == 1 {
				if j >= prices[i] {
					intermediateResult[i][j] = defaultMaxQuantities[i]
				}
			} else if j < prices[i] {
				intermediateResult[i][j] = intermediateResult[i-1][j]
			} else {
				intermediateResult[i][j] = int(math.Max(float64(intermediateResult[i-1][j]), float64(defaultMaxQuantities[i]+intermediateResult[i-1][j-prices[i]])))
			}
		}
	}

	//fmt.Println("intermediateResult", intermediateResult)

	var responseIndex []int

	j := amount
	for i := len(prices) - 1; i > 0; i-- {
		if intermediateResult[i][j] != intermediateResult[i-1][j] && intermediateResult[i][j] == (intermediateResult[i-1][j-prices[i]]+defaultMaxQuantities[i]) {
			responseIndex = append(responseIndex, i)
			j -= prices[i]
		}
	}
	sort.Ints(responseIndex)
	fmt.Println(responseIndex)
}
