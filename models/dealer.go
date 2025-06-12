package models

type Dealer struct {
	Nazwa string `gorm:"primaryKey"`
	Adres string
}
