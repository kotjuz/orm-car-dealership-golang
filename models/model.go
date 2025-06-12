package models

type Model struct {
	ID                 int `gorm:"primaryKey"`
	Nazwa              string
	RokWprowadzenia    int
	Typ                string
	LiczbaPasazerow    int
	PojemnoscBagaznika int
	Ladownosc          int
	Poprzednik         int
	MarkaNazwa         string `gorm:"size:100"`

	Marka          Marka  `gorm:"foreignKey:MarkaNazwa;references:Nazwa"`
	PoprzedniModel *Model `gorm:"foreignKey:Poprzednik;references:ID"`
}
