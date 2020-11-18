package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	_ "os"
)

type CarShop struct{
	Name string
	Year string
	Price string
	Country string
	AllCount string
}
const (
	DbUser     = "postgres"
	DbPassword = "ashumqwe"
	DbName     = "lab"
)
func dbConnect() error {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS car_shop (car_name text, car_year text,car_price text, car_country text)"); err != nil {
	return err
}
return nil
}
func dbAddCar(name, year, price, country string) error {
	sqlstmt := "INSERT INTO car_shop VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(sqlstmt, name, year, price, country)
	if err != nil {
		return err
	}
	return nil
}
func dbGetCars() ([]CarShop, error) {
	var cars []CarShop
	stmt, err := db.Prepare("SELECT car_name, car_year,car_price, car_country  FROM car_shop")
	if err != nil {
		return cars, err
	}
	res, err := stmt.Query()
	if err != nil {
		return cars, err
	}
	var tempCar CarShop
	for res.Next() {
		err = res.Scan(&tempCar.Name, &tempCar.Year, &tempCar.Price, &tempCar.Country)
		if err != nil {
			return cars, err
		}
		cars = append(cars, tempCar)
	}
	return cars, err
}
func getAllCount() (CarShop, error){
	var car CarShop
	stmt, err := db.Prepare("select count(*) num from car_shop")
	if err != nil {
		return car, err
	}
	res, err := stmt.Query()
	if err != nil {
		return car, err
	}
	var tempCar CarShop
	for res.Next() {
		err = res.Scan(&tempCar.AllCount)
		if err != nil {
			return car, err
		}
		car = tempCar
	}
	return car, err
}

func getCarsByMark(mark string) ([]CarShop, error) {
	var cars []CarShop
	stmt, err := db.Prepare("SELECT car_name, car_year,car_price, car_country  FROM car_shop WHERE car_name ='" + mark + "'")
	if err != nil {
		return cars, err
	}
	res, err := stmt.Query()
	if err != nil {
		return cars, err
	}
	var tempCar CarShop
	for res.Next() {
		err = res.Scan(&tempCar.Name, &tempCar.Year, &tempCar.Price, &tempCar.Country)
		if err != nil {
			return cars, err
		}
		cars = append(cars, tempCar)
	}
	return cars, err
}

