package models

type TypSilnika struct {
	ID             int `gorm:"primaryKey"`
	RodzajPaliwa   string
	OpisParametrow string
}
