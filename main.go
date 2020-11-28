package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)
var db *sql.DB
func rollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_list.html")
		if err != nil {
			log.Fatal(err)
		}
		cars, err := dbGetCars()
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, cars)
	}
}
func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("count.html")
		if err != nil {
			log.Fatal(err)
		}
		count, err := dbGetCarsCount()
		if err != nil {
			log.Fatal(err)
		}
		var countRes CountRes
		countRes.Count = count
		t.Execute(w, countRes)
	}
}
func addCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_form.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		brand := r.Form.Get("brand")
		country:= r.Form.Get("country")
		production_year, errI := strconv.Atoi(r.Form.Get("production_year"))
		price, errI := strconv.Atoi(r.Form.Get("price"))
		err := dbAddCar(brand, country, production_year, price)
		if err != nil || errI != nil {
			log.Fatal(err)
		}
	}
}
func getByBrandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("brand_ask.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		t, err := template.ParseFiles("brand_res.html")
		r.ParseForm()
		brand := r.Form.Get("brand")
		res, err := dbGetCarsByBrand(brand)
		log.Println(len(res))
		if err != nil || err != nil {
			log.Fatal(err)
		}
		t.Execute(w, res)
	}
}
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println(port)
	}
	return ":" + port
}

func main() {
	err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", rollHandler)
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/add", addCarHandler)
	http.HandleFunc("/brand", getByBrandHandler)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}
