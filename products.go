package courseGo

type Projector struct {
	Id         int `json:"id"`
	CategoryId int
	Quantity   int    `json:"quantity" binding:"required"`
	Brightness string `json:"brightness" binding:"required"`
	Contrast   string `json:"contrast" binding:"required"`
	Price      int    `json:"price" binding:"required"`
}

type VideoWall struct {
	Id         int `json:"id"`
	CategoryId int
	Quantity   int    `json:"quantity" binding:"required"`
	Brightness string `json:"brightness" binding:"required"`
	Contrast   string `json:"contrast" binding:"required"`
	Price      int    `json:"price" binding:"required"`
}

type Monitor struct {
	Id         int `json:"id"`
	CategoryId int
	Quantity   int    `json:"quantity" binding:"required"`
	Brightness string `json:"brightness" binding:"required"`
	Contrast   string `json:"contrast" binding:"required"`
	Price      int    `json:"price" binding:"required"`
}

type Mount struct {
	Id         int `json:"id"`
	CategoryId int
	Quantity   int    `json:"quantity" binding:"required"`
	Size       string `json:"size" binding:"required"`
	Price      int    `json:"price" binding:"required"`
}

type Category struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"category"`
}
