package courseGo

type Quantity struct {
	Id         int    `json:"id" binding:"required"`
	ProductId  int    `json:"productId" binding:"required" db:"prodId"`
	ExtraPosId int    `json:"extraPosId" binding:"required"`
	Reciever   string `json:"reciever" binding:"required"`
}

type UsersQuantity struct {
	id         int
	UserId     int
	QuantityId int
}
