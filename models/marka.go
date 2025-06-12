package models

type Marka struct {
	Nazwa        string `gorm:"primaryKey;size:100"`
	RokZalozenia int
}
