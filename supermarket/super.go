package supermarket

import (
	"errors"
)

var listitems = map[string]interface{}{
	"rice":       490.87,
	"coffee":      75.60,
	"cereals":    350.50,
	"chocolates": 100.90,
	"biscuit": 30.50,
}

func GetProduct(productname string) (interface{}, error) {
	if listitems[productname] != nil {
		return listitems[productname], nil
	} else {
		return nil, errors.New("Item not found")
	}
}

func PostProduct(productname string, value interface{}) error {
	if listitems[productname] != nil {
		return errors.New("Item is already in the list")
	} else {
		if listitems == nil {
			listitems = make(map[string]interface{})
			listitems[productname] = value
		} else {
			listitems[productname] = value
		}
		return nil
	}

}

func PutProduct(productname string, value interface{}) error {

	if listitems[productname] == nil {
		return errors.New("Item is not in the list")
	} else {
		listitems[productname] = value
		return nil
	}
}

func DeleteProduct(productname string) error {
	if listitems[productname] != nil {
		delete(listitems, productname)
		return nil
	} else {
		return errors.New("Item not found")
	}
}
func PrintProduct() map[string]interface{} {
	return listitems
}
