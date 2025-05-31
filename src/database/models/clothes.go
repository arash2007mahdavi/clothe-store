package models

type Hat struct {
	Id int `gorm:"primarykey"`
	Price    float64 `gorm:"type:float;null"`
	Currency string  `gorm:"type:string;null"`
}

type Shoes struct {
	Id int `gorm:"primarykey"`
	Price    float64 `gorm:"type:float;null"`
	Currency string  `gorm:"type:string;null"`
}

type Pant struct {
	Id int `gorm:"primarykey"`
	Price    float64 `gorm:"type:float;null"`
	Currency string  `gorm:"type:string;null"`
}

type Shirt struct {
	Id int `gorm:"primarykey"`
	Price    float64 `gorm:"type:float;null"`
	Currency string  `gorm:"type:string;null"`
}
