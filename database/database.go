package database

import (
	"fmt"
	"os"
	"time"

	"example.com/ormproject/models"
	"github.com/lpernett/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()

	if err != nil {
		panic("Could not load env")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&TrustServerCertificate=true",
		user, password, host, database)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = db

	err = db.AutoMigrate(
		&models.Samochod{},
		&models.PosSilnik{},
		&models.TypSilnika{},
		&models.Dealer{},
		&models.Model{},
		&models.Sprzedaz{},
		&models.Marka{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Could not migrate")
	}
}

func SeedData() {

	marki := []models.Marka{
		{Nazwa: "Toyota", RokZalozenia: 1937},
		{Nazwa: "Volvo", RokZalozenia: 1927},
		{Nazwa: "Ford", RokZalozenia: 1903},
	}
	for _, m := range marki {
		DB.FirstOrCreate(&m, models.Marka{Nazwa: m.Nazwa})
	}

	var nilPoprzednik *int = nil

	modele := []models.Model{
		{ID: 1, Nazwa: "Corolla", RokWprowadzenia: 2020, Typ: "osobowy", LiczbaPasazerow: 5, PojemnoscBagaznika: 450, Poprzednik: nilPoprzednik, MarkaNazwa: "Toyota"},
		{ID: 2, Nazwa: "XC60", RokWprowadzenia: 2021, Typ: "ciezarowy", Ladownosc: 2000, Poprzednik: nilPoprzednik, MarkaNazwa: "Volvo"},
		{ID: 3, Nazwa: "Focus", RokWprowadzenia: 2019, Typ: "osobowy", LiczbaPasazerow: 5, PojemnoscBagaznika: 350, Poprzednik: nilPoprzednik, MarkaNazwa: "Ford"},
	}
	for _, m := range modele {
		if m.Typ == "osobowy" || m.Typ == "ciezarowy" {
			DB.FirstOrCreate(&m, models.Model{ID: m.ID})
		} else {
			fmt.Println("Couldn't add to database - wrong type", m)
		}

	}

	silniki := []models.TypSilnika{
		{ID: 1, RodzajPaliwa: "benzyna", OpisParametrow: "1.6L, 120KM"},
		{ID: 2, RodzajPaliwa: "diesel", OpisParametrow: "2.0L, 150KM"},
		{ID: 3, RodzajPaliwa: "hybryda", OpisParametrow: "1.8L, 100KM + silnik elektryczny"},
	}
	for _, s := range silniki {
		DB.FirstOrCreate(&s, models.TypSilnika{ID: s.ID})
	}

	dealerzy := []models.Dealer{
		{Nazwa: "Auto Świat", Adres: "Warszawa, ul. Samochodowa 1"},
		{Nazwa: "Nordic Cars", Adres: "Gdańsk, ul. Morska 5"},
		{Nazwa: "Fast Wheels", Adres: "Kraków, ul. Królewska 10"},
	}
	for _, d := range dealerzy {
		DB.FirstOrCreate(&d, models.Dealer{Nazwa: d.Nazwa})
	}

	samochody := []models.Samochod{
		{VIN: "JTDKB20U687923335", RokProdukcji: 2022, Przebieg: 15000, SkrzyniaBiegow: "automatyczna", TypSil: 1, KrajPochodzenia: "Japonia", ModelID: 1, DealerNazwa: "Auto Świat"},
		{VIN: "YV1LW58T6Y2512365", RokProdukcji: 2023, Przebieg: 5000, SkrzyniaBiegow: "manualna", TypSil: 2, KrajPochodzenia: "Szwecja", ModelID: 2, DealerNazwa: "Nordic Cars"},
		{VIN: "WF0DP3TH5J4125632", RokProdukcji: 2021, Przebieg: 30000, SkrzyniaBiegow: "manualna", TypSil: 3, KrajPochodzenia: "USA", ModelID: 3, DealerNazwa: "Fast Wheels"},
		{VIN: "TEST1111111111111", RokProdukcji: 2021, Przebieg: 30000, SkrzyniaBiegow: "manualna", TypSil: 3, KrajPochodzenia: "USA", ModelID: 3, DealerNazwa: "Nordic Cars"},
	}
	for _, s := range samochody {
		DB.FirstOrCreate(&s, models.Samochod{VIN: s.VIN})
	}

	sprzedaze := []models.Sprzedaz{
		{Data: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), Cena: 95000.00, DealerNazwa: "Auto Świat", SamochodVIN: "JTDKB20U687923335"},
		{Data: time.Date(2024, 2, 20, 0, 0, 0, 0, time.UTC), Cena: 150000.00, DealerNazwa: "Nordic Cars", SamochodVIN: "YV1LW58T6Y2512365"},
		{Data: time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC), Cena: 75000.00, DealerNazwa: "Fast Wheels", SamochodVIN: "WF0DP3TH5J4125632"},
	}
	for _, sp := range sprzedaze {
		DB.FirstOrCreate(&sp, models.Sprzedaz{Data: sp.Data, DealerNazwa: sp.DealerNazwa, SamochodVIN: sp.SamochodVIN})
	}

	posSilniki := []models.PosSilnik{
		{TypSilId: 1, ModelId: 1},
		{TypSilId: 1, ModelId: 2},
		{TypSilId: 1, ModelId: 3},
		{TypSilId: 2, ModelId: 2},
		{TypSilId: 3, ModelId: 3},
	}
	for _, ps := range posSilniki {
		DB.FirstOrCreate(&ps, models.PosSilnik{TypSilId: ps.TypSilId, ModelId: ps.ModelId})
	}
}
