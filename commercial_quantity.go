package courseGo

import "github.com/google/uuid"

type CommQuantity struct {
	Id       uuid.UUID       `json:"id" binding:"required" db:"id"`
	Product  []ProdInventory `json:"products"  `
	Receiver string          `json:"receiver"`
}

type UsersCommQuantity struct {
	id         int
	UserId     int
	QuantityId int
}

type Settings struct {
	Roi int `json:"roi" binding:"required"`
}
