
# Simple Restful API in Golang with GORM

Prosty projekt API napisany w Go, wykorzystujący ORM (GORM) do komunikacji z bazą danych.

---

## 📁 Struktura katalogów

```
/orm_project
│
├── /api-test        
│   └── add-car.http
│   └── sell-car.http
│
├── /database
│   └── database.go
│
├── /models        
│   └── dealer.go
│   └── marka.go
│   └── model.go
│   └── samochod.go
│   └── sprzedaz.go
│   └── posSilnik.go
│
├── /routes    
│   └── routes.go
│   └── cars.go
│
├── /database     
│   └── database.go
│
├── main.go      
│
└── go.mod    
└── go.sum      
```

---

## Wymagania

- Go 1.22+
- GORM 
- MSSQL

---

## Uruchomienie

```bash
go run main.go
```

---

## Przykładowe endpointy

### DELETE `/car/sell/:vin`

Usuwa auto o podanym numerze VIN (jesli istnieje)


### POST `/car/add`

Dodaje nowe auto do bazy (jesli jeszcze takie nie istnieje).



## TODO
- [ ] Dockerfile, docker-compose


