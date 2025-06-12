package models

type Dealer struct {
	Nazwa string `gorm:"primaryKey;size:100"`
	Adres string `gorm:"size:100"`
}
