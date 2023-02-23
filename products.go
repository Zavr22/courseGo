package courseGo

type Projector struct {
	Id         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	CategoryId int
	Quantity   int     `json:"quantity" binding:"required"`
	Brightness string  `json:"brightness" binding:"required"`
	Contrast   string  `json:"contrast" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
}

type VideoWall struct {
	Id         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	CategoryId int
	Quantity   int     `json:"quantity" binding:"required"`
	Brightness string  `json:"brightness" binding:"required"`
	Contrast   string  `json:"contrast" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
}

type Monitor struct {
	Id         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	CategoryId int
	Quantity   int     `json:"quantity" binding:"required"`
	Brightness string  `json:"brightness" binding:"required"`
	Contrast   string  `json:"contrast" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
}

type Mount struct {
	Id         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	CategoryId int
	Quantity   int     `json:"quantity" binding:"required"`
	Size       string  `json:"size" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
}

type Category struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"category"`
}

type ProdInventory struct {
	ProdId     int `json:"prod_Id"`
	CategoryId int `json:"category_Id"`
}
