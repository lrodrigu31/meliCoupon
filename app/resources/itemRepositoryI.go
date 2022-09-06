package resources

import "coupon/app/models"

//ItemFromRepository : is the interface from ItemRepository
type ItemFromRepository interface {
	getItem(string) models.Item
}

//GetItem : is the public func for access to private func getItem from ItemRepository
func GetItem(itemMeliId string) models.Item {
	return ItemRepository{}.getItem(itemMeliId)
}
