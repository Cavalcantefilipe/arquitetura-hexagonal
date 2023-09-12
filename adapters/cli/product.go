package cli

import (
	"fmt"

	"github.com/filipe/exagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "get":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s\n", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s\n", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		product, err = service.Enable(product)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s\n", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		product, err = service.Disable(product)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s\n", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	default:
		result = "Invalid action"
	}

	return result, nil
}
