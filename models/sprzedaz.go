package models

import "time"

type Sprzedaz struct {
	ID          int `gorm:"primaryKey"`
	Data        time.Time
	Cena        float64
	DealerNazwa string   `gorm:"size:100"`
	SamochodVIN string   `gorm:"size:17"`
	Dealer      Dealer   `gorm:"foreignKey:DealerNazwa;references:Nazwa"`
	Samochod    Samochod `gorm:"foreignKey:SamochodVIN;references:VIN"`
}
