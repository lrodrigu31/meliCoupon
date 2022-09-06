package failures

import (
	"coupon/app/helpers"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func intermediateResultTable(prices []float64, defaultMaxQuantities []float64, amount float64) [][]float64 {
	rows := len(prices)
	columns := int(amount) + 1

	var intermediateResult [][]float64
	// Tabla de ceros

	for i := 0; i < rows; i++ {
		var vectorFila []float64
		for j := 0; j < columns; j++ {
			vectorFila = append(vectorFila, 0)
		}
		intermediateResult = append(intermediateResult, vectorFila)
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= int(amount); j++ {
			if i == 1 {
				if j >= int(prices[i]) {
					intermediateResult[i][j] = defaultMaxQuantities[i]
				}
			} else if j < int(prices[i]) {
				intermediateResult[i][j] = intermediateResult[i-1][j]
			} else {
				//indexI := defaultMaxQuantities
				//indexJ := intermediateResult[i-1][j-int(prices[i])]
				//index := float64(indexI) + intermediateResult[i-1][j-int(prices[i])]
				intermediateResult[i][j] = math.Max(intermediateResult[i-1][j], defaultMaxQuantities[i]+intermediateResult[i-1][j-int(prices[i])])
			}
		}
	}

	return intermediateResult
}

func getResponseIndex(amount float64, prices []float64, keys []string, intermediateResult [][]float64, defaultMaxQuantities []float64) ([]string, float64) {
	var indices []int
	var responseIndex []string
	var total float64

	j := amount

	for i := len(prices) - 1; i > 0; i-- {
		if intermediateResult[i][int(j)] != intermediateResult[i-1][int(j)] && intermediateResult[i][int(j)] == intermediateResult[i-1][int(j-prices[i])]+defaultMaxQuantities[i] {
			indices = append(indices, i)
			j -= prices[i]
		}
	}

	//sort.Ints(indices)

	fmt.Println(indices)
	fmt.Println(keys)
	fmt.Println(prices)

	for i := 0; i < len(indices); i++ {
		j := indices[i] - 1
		responseIndex = append(responseIndex, keys[j])
		total = total + prices[j]
	}

	return responseIndex, total
}

func reducer(items map[string]float64, amount float64) ([]float64, float64) {
	prices := helpers.Values(items)
	min := minPrice(items)
	factor := math.Pow10(len(strconv.Itoa(int(min))) - 1)
	str := strconv.Itoa(int(min))
	fmt.Println(str, reflect.TypeOf(str))
	fmt.Println(amount / factor)

	for i := 0; i < len(prices); i++ {
		prices[i] = prices[i] / factor
	}
	return prices, amount / factor
}

func maximizer(items map[string]float64, amount float64) ([]float64, float64) {
	prices := helpers.Values(items)
	min := minPrice(items)
	factor := math.Pow10(len(strconv.Itoa(int(min))) - 1)
	for i := 0; i < len(prices); i++ {
		prices[i] = prices[i] * factor
	}
	return prices, amount * factor
}

func minPrice(items map[string]float64) float64 {
	prices := helpers.Values(items)
	min := prices[0]
	for _, price := range prices {
		min = math.Min(price, min)
	}
	return min
}

/*
func Calculate2(items map[string]float64, amount float64) ([]string, float64) {

	maps = make(map[string]float64)
	listItemsSelected, total := ProcessResponse(items, amount)
	fmt.Println("::::::::::::::")
	fmt.Println(listItemsSelected)
	fmt.Println(total)
	fmt.Println("::::::::::::::")

	return listItemsSelected, total
}*/

/*func ProcessResponse(items map[string]float64, amount float64) ([]string, float64) {
	prices := helpers.Values(items)
	ids := helpers.Keys(items)
	pos := len(prices) - 1
	total := getFinalTotal(pos, amount, prices)

	return getFinalResultItems(prices, ids, amount), total

}*/
