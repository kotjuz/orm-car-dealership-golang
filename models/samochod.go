package models

type Samochod struct {
	VIN             string `gorm:"primaryKey;size:17"`
	RokProdukcji    int
	Przebieg        int
	SkrzyniaBiegow  string `gorm:"size:100"`
	TypSil          int
	KrajPochodzenia string `gorm:"size:100"`
	ModelID         int
	DealerNazwa     string     `gorm:"size:100"`
	TypSilnika      TypSilnika `gorm:"foreignKey:TypSil;references:ID"`
	Model           Model      `gorm:"foreignKey:ModelID;references:ID"`
	Dealer          Dealer     `gorm:"foreignKey:DealerNazwa;references:Nazwa"`
}
