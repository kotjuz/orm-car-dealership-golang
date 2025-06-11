package models

import "time"

type Sprzedaz struct {
	data        time.Time
	cena        float64
	klientId    int
	dealerNazwa string
	samochodVin string
}
