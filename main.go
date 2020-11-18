package main
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

var db *sql.DB
func rollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Query().Get("name") == "" {
			t, err := template.ParseFiles("simple_list.html")
			if err != nil {
				log.Fatal(err)
			}
			cars, err := dbGetCars()
			if err != nil {
				log.Fatal(err)
			}

			t.Execute(w, cars)
		} else {
			t, err := template.ParseFiles("simple_list.html")
			if err != nil {
				log.Fatal(err)
			}
			cars, err := getCarsByMark(r.URL.Query().Get("name"))
			if err != nil {
				log.Fatal(err)
			}

			t.Execute(w, cars)
		}
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
		name := r.Form.Get("name")
		year := r.Form.Get("year")
		price := r.Form.Get("price")
		country := r.Form.Get("country")
		err := dbAddCar(name, year, price, country)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/add", http.StatusSeeOther)
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
	http.HandleFunc("/add", addCarHandler)
	http.HandleFunc("/count", countCarHandler)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}

func countCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("all_count.html")
		if err != nil {
			log.Fatal(err)
		}
		car, err := getAllCount()
		if err != nil {
			log.Fatal(err)
		}
//df
		t.Execute(w, car)
	}
}

