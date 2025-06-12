package models

type Samochod struct {
	VIN             string `gorm:"primaryKey"`
	RokProdukcji    int
	Przebieg        int
	SkrzyniaBiegow  string
	TypSil          int
	KrajPochodzenia string
	ModelID         int
	DealerNazwa     string
	TypSilnika      TypSilnika `gorm:"foreignKey:TypSil;references:ID"`
	Model           Model      `gorm:"foreignKey:ModelID;references:ID"`
	Dealer          Dealer     `gorm:"foreignKey:DealerNazwa;references:Nazwa"`
}
