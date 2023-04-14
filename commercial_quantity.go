package courseGo

type CommQuantity struct {
	Id      int    `json:"id" binding:"required" db:"id"`
	Product string `json:"products" db:"products" `
	Status  string `json:"status" db:"status"`
}

type UsersCommQuantity struct {
	id         int
	UserId     int
	QuantityId int
}
