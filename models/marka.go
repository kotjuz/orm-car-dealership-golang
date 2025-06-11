package models

type Marka struct {
	nazwa        string `gorm:"primaryKey"`
	rokZalozenia int    `gorm:"not null;check:year_founded > 1900"`
}
