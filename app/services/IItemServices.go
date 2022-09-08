package services

//service : is the interface for process the service data
type service interface {
	calculate(items map[string]float64, amount float64) []string
}

// GetPriceList : is the public method that return the price list
func (itemService ItemService) GetPriceList(itemIds []string) (PriceList, float64) {
	return itemService.getPriceList(itemIds)
}

// ProcessResponse : is the public method for process the service data
func (itemService ItemService) ProcessResponse(items map[string]float64, amount float64) ([]string, float64) {
	return itemService.processResponse(items, amount)
}
