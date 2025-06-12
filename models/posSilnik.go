package models

type PosSilnik struct {
	TypSilId int `gorm:"primaryKey"` // część klucza głównego
	ModelId  int `gorm:"primaryKey"` // część klucza głównego

	TypSilnik TypSilnika `gorm:"foreignKey:TypSilId;references:ID"`
	Model     Model      `gorm:"foreignKey:ModelId;references:ID"`
}
