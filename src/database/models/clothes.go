package models

type Hat struct {
	Price    float64 `gorm:"type:float;null"`
	Currency string  `gorm:"type:string;null"`
}

type Shoes struct {
	Price    float64 `gorm:"type:float;null"`
	Currency string  `gorm:"type:string;null"`
}

type Pant struct {
	Price    float64 `gorm:"type:float;null"`
	Currency string  `gorm:"type:string;null"`
}

type Shirt struct {
	Price    float64 `gorm:"type:float;null"`
	Currency string  `gorm:"type:string;null"`
}
