package services

import "coupon/app/models"

//ResponseServicesI : is the interface for services response
type ResponseServicesI interface {
	ResponseItemServices(data models.InputData) (models.OutputData, bool)
}

// ResponseItemServices : is the public method for services response
func (response ResponseServices) ResponseItemServices(data models.InputData) (models.OutputData, bool) {
	return response.responseItemServices(data)
}
