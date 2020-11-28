package main

import (
	"database/sql"
	"fmt"
)
type Car struct{
	Brand string
	Country string
	ProductionYear int
	Price int
}
const (
	DB_USER = "postgres"
	DB_PASSWORD = "root"
	DB_NAME = "cars"
)
func dbConnect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=0.0.0.0 port=5432 user=%s password=%s dbname=%s sslmode=enable",
	DB_USER, DB_PASSWORD, DB_NAME))
	if err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS car (brand text, country text, production_year integer, price integer )"); err != nil {
	return err
}
return nil
}

func dbAddCar(brand string, country string, production_year int, price int) error {
	sqlstmt := "INSERT INTO car VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(sqlstmt, brand, country, production_year, price)
	if err != nil {
		return err
	}
	return nil
}

func dbGetCars() ([]Car, error) {
	var cars []Car
	stmt, err := db.Prepare("SELECT brand, country, production_year, price FROM car")
	if err != nil {
		return cars, err
	}
	res, err := stmt.Query()
	if err != nil {
		return cars, err
	}
	var tempCar Car
	for res.Next() {
		err = res.Scan(&tempCar.Brand, &tempCar.Country, &tempCar.ProductionYear, &tempCar.Price)
		if err != nil {
			return cars, err
		}
		cars = append(cars, tempCar)
	}
	return cars, err
}
