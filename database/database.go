package database

import (
	"gorm.io/gorm"
)

var (
	Conn *gorm.DB
)

type Movie struct {
	gorm.Model
	Title               string              `json:"title"`
	Year                int                 `json:"year"`
	Actors              []Actor             `json:"actors"`
	ProductionCompanies []ProductionCompany `json:"production_company"`
}

type ProductionCompany struct {
	gorm.Model
	Name    string `json:"name"`
	MovieID int    `json:"movie_id"`
}

type Actor struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	MovieID   int    `json:"movie_id"`
}
