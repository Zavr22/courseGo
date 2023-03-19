package courseGo

type CommQuantity struct {
	Id       int             `json:"id" binding:"required"`
	Product  []ProdInventory `json:"products" binding:"required" `
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
