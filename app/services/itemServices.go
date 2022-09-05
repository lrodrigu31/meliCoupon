package services

import (
	"coupon/app/helpers"
	"coupon/app/models"
	"coupon/app/repository"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

var maps map[string]float64

func Calculate(items map[string]float64, amount float64) ([]string, float64) {

	maps = make(map[string]float64)
	//var prices []float64
	//prices := helpers.Values(items)
	//defaultMaxQuantities := prices
	//keys := helpers.Keys(items)
	listItemsSelected, total := processResponse(items, amount)
	fmt.Println("::::::::::::::")
	fmt.Println(listItemsSelected)
	fmt.Println(total)
	fmt.Println("::::::::::::::")
	//reducer(items, amount)
	//intermediateResult := intermediateResultTable(prices, defaultMaxQuantities, amount)

	//fmt.Println(intermediateResult)

	//list, total := getResponseIndex(amount, prices, keys, intermediateResult, defaultMaxQuantities)

	return listItemsSelected, total
}

func GetPriceList(itemIds []string) (models.PriceList, float64) {
	priceList := make(map[string]float64)
	for _, itemId := range itemIds {
		item := repository.GetItem(itemId)
		if _, present := priceList[item.Id]; !present {
			if item.ValidateStructure() {
				priceList[item.Id] = item.Price
			}
		}
	}
	listMapped := models.PriceList{}
	listMapped.Items = priceList
	return listMapped, MinPrice(priceList)
}

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

func MinPrice(items map[string]float64) float64 {
	prices := helpers.Values(items)
	min := prices[0]
	for _, price := range prices {
		min = math.Min(price, min)
	}
	return min
}

func reducer(items map[string]float64, amount float64) ([]float64, float64) {
	prices := helpers.Values(items)
	min := MinPrice(items)
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
	min := MinPrice(items)
	factor := math.Pow10(len(strconv.Itoa(int(min))) - 1)
	for i := 0; i < len(prices); i++ {
		prices[i] = prices[i] * factor
	}
	return prices, amount * factor
}

func processResponse(items map[string]float64, amount float64) ([]string, float64) {
	prices := helpers.Values(items)
	ids := helpers.Keys(items)
	pos := len(prices) - 1
	total := getTotal(pos, amount, prices)

	return getSelectedItems(prices, ids, amount), total

}

func getTotal(pos int, amount float64, prices []float64) float64 {
	if pos < 0 || amount == 0 {
		return 0
	}
	key := fmt.Sprintf("%d-%v", pos, strconv.FormatFloat(amount, 'E', -1, 64))

	if _, present := maps[key]; present {
		return maps[key]
	}

	if prices[pos] > amount {
		maps[key] = getTotal(pos-1, amount, prices)
	} else {
		include := getTotal(pos-1, amount-prices[pos], prices) + prices[pos]
		exclude := getTotal(pos-1, amount, prices)
		maps[key] = math.Max(include, exclude)
	}

	return maps[key]
}

func getSelectedItems(prices []float64, items []string, amount float64) []string {
	var itemList []string
	key := fmt.Sprintf("%d-%v", len(prices)-1, strconv.FormatFloat(amount, 'E', -1, 64))

	total := maps[key]

	for i := len(prices) - 1; i > 0; i-- {

		key := fmt.Sprintf("%d-%v", i-1, strconv.FormatFloat(amount, 'E', -1, 64))

		if total != maps[key] {
			itemList = append(itemList, items[i])
			amount -= prices[i]
			total -= prices[i]
		}
	}

	if total != 0 {
		itemList = append(itemList, items[0])
	}

	return itemList
}
