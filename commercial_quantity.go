package courseGo

type CommQuantity struct {
	Id         int    `json:"id" binding:"required"`
	ProductId  int    `json:"productId" binding:"required" db:"prodId"`
	ExtraPosId int    `json:"extraPosId" binding:"required"`
	Reciever   string `json:"reciever" binding:"required"`
}

type UsersCommQuantity struct {
	id         int
	UserId     int
	QuantityId int
}
