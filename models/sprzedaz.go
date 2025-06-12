package models

import "time"

type Sprzedaz struct {
	ID          int `gorm:"primaryKey"`
	Data        time.Time
	Cena        float64
	DealerNazwa string
	SamochodVIN string
	Dealer      Dealer   `gorm:"foreignKey:DealerNazwa;references:Nazwa"`
	Samochod    Samochod `gorm:"foreignKey:SamochodVIN;references:VIN"`
}
