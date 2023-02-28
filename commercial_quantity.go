package courseGo

type CommQuantity struct {
	Id         int               `json:"id" binding:"required"`
	ProductId  map[string]string `json:"productsId" binding:"required" db:"prodId"`
	ExtraPosId map[string]string `json:"extraPosId" binding:"required"`
	Reciever   string            `json:"reciever" binding:"required"`
}

type UsersCommQuantity struct {
	id         int
	UserId     int
	QuantityId int
}
