package services

import (
	"coupon/app/helpers"
	"coupon/app/repositories"
	"fmt"
	"math"
	"strconv"
)

var maps map[string]float64

type PriceList struct {
	Items map[string]float64
}

type ItemService struct {
}

// METHODS

//calculate : this method return the final item list selected that keep condition
func (itemService ItemService) calculate(Map map[string]float64, amount float64) []string {
	prices := helpers.Values(Map)
	ids := helpers.Keys(Map)
	var itemList []string

	key := fmt.Sprintf("%d-%v", len(prices)-1, strconv.FormatFloat(amount, 'E', -1, 64))

	total := maps[key]

	for i := len(prices) - 1; i > 0; i-- {

		key := fmt.Sprintf("%d-%v", i-1, strconv.FormatFloat(amount, 'E', -1, 64))

		if total != maps[key] {
			itemList = append(itemList, ids[i])
			amount -= prices[i]
			total -= prices[i]
		}
	}

	if total != 0 {
		itemList = append(itemList, ids[0])
	}

	return itemList
}

//processResponse : this method return the service response employ the functions calculate and getFinalTotal
func (itemService ItemService) processResponse(items map[string]float64, amount float64) ([]string, float64) {
	maps = make(map[string]float64)
	prices := helpers.Values(items)

	pos := len(prices) - 1
	total := getFinalTotal(pos, amount, prices)

	return ItemService{}.calculate(items, amount), total

}

// getPriceList : this method return the price list
func (itemService ItemService) getPriceList(itemIds []string) (PriceList, float64) {
	priceList := make(map[string]float64)
	for _, itemId := range itemIds {
		item := repositories.GetItem(itemId)
		if _, present := priceList[item.Id]; !present {
			if item.ValidateStructure() {
				priceList[item.Id] = item.Price
			}
		}
	}
	listMapped := PriceList{}
	listMapped.Items = priceList
	return listMapped, minPrice(priceList)
}

//FUNCTIONS
//getFinalTotal : this func calculate the total value of solution
func getFinalTotal(pos int, amount float64, prices []float64) float64 {
	if pos < 0 || amount == 0 {
		return 0
	}
	key := fmt.Sprintf("%d-%v", pos, strconv.FormatFloat(amount, 'E', -1, 64))

	if _, present := maps[key]; present {
		return maps[key]
	}

	if prices[pos] > amount {
		maps[key] = getFinalTotal(pos-1, amount, prices)
	} else {
		include := getFinalTotal(pos-1, amount-prices[pos], prices) + prices[pos]
		exclude := getFinalTotal(pos-1, amount, prices)
		maps[key] = math.Max(include, exclude)
	}

	return maps[key]
}

//minPrice : this func return the min value in array
func minPrice(items map[string]float64) float64 {
	prices := helpers.Values(items)
	min := prices[0]
	for _, price := range prices {
		min = math.Min(price, min)
	}
	return min
}
