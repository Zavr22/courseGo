package courseGo

type Projector struct {
	Id            int     `json:"id"`
	Name          string  `json:"name" binding:"required"`
	CategoryId    int     `json:"category-id" binding:"required" db:"category_id"`
	Quantity      int     `json:"quantity" binding:"required"`
	Brightness    string  `json:"brightness" binding:"required"`
	Contrast      string  `json:"contrast" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
	Weight        float64 `json:"weight" binding:"required"`
	FocalDistance float64 `json:"focalDistance" binding:"required"`
}

type VideoWall struct {
	Id         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	CategoryId int
	Quantity   int     `json:"quantity" binding:"required"`
	Brightness string  `json:"brightness" binding:"required"`
	Contrast   string  `json:"contrast" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Weight     float64 `json:"weight" binding:"required"`
}

type Monitor struct {
	Id         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	CategoryId int
	Quantity   int     `json:"quantity" binding:"required"`
	Brightness string  `json:"brightness" binding:"required"`
	Contrast   string  `json:"contrast" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Weight     float64 `json:"weight" binding:"required"`
}

type Mount struct {
	Id         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	CategoryId int
	Quantity   int     `json:"quantity" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	MaxWeight  float64 `json:"max-weight" binding:"required"`
	Roi        float64 `json:"roi" binding:"required"`
}

type Category struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"category"`
}

type ProdInventory struct {
	Name  string  `json:"name"`
	Price float64 `json:"price" binding:"required"`
}

type Params struct {
	CategoryId    int     `json:"category-id" `
	Quantity      int     `json:"quantity" binding:"required"`
	Brightness    string  `json:"brightness" `
	Contrast      string  `json:"contrast" `
	Price         float64 `json:"price" `
	Weight        float64 `json:"weight" binding:"required"`
	ExtraRoi      float64 `json:"extra_roi" `
	FocalDistance float64 `json:"focalDistance" binding:"required"`
}
