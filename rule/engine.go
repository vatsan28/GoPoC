package rule

import (
	"GoPoc/models"
	"log"
)

func IsProductValid(product models.Order) bool {
	log.Println("In rule engine check.", product)
	return product.IsProductValid == true
}