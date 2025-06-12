package models

type PosSilnik struct {
	TypSilId  int        `gorm:"primaryKey"`
	ModelId   int        `gorm:"primaryKey"`
	TypSilnik TypSilnika `gorm:"foreignKey:TypSilId;references:ID"`
	Model     Model      `gorm:"foreignKey:ModelId;references:ID"`
}
