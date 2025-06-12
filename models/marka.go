package models

type Marka struct {
	Nazwa        string `gorm:"primaryKey"`
	RokZalozenia int    `gorm:"not null;check:rokZalozenia > 1900"`
}
