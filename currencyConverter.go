package main

import (
	"html/template"
	"net/http"

	"strconv"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var templates *template.Template
var client *redis.Client
var UserData = make(map[string]string)

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", webGethandler).Methods("GET")
	r.HandleFunc("/", webPosthandler).Methods("POST")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func webGethandler(res http.ResponseWriter, req *http.Request) {
	results, err := client.LRange("results", 0, 0).Result()

	if err != nil || results[0] == "-1" {
		result := make(map[int]string)
		result[0] = "A valid amount must be a number greater than zero"
		templates.ExecuteTemplate(res, "index.html", result)
	} else {
		templates.ExecuteTemplate(res, "index.html", results)
	}
}

func webPosthandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	UserData["sourceCurrency"] = req.PostForm.Get("sourceCurrency")
	UserData["destinationCurrency"] = req.PostForm.Get("destinationCurrency")
	UserData["amount"] = req.PostForm.Get("amount")
	data := CurrencyConverter(UserData)
	client.LPush("results", data)
	http.Redirect(res, req, "/", http.StatusFound)
}

func CurrencyConverter(UserData map[string]string) float64 {
	currencyData := make(map[string]float64)

	currencyData["USD"] = 1.0
	currencyData["GHS"] = 6.4
	currencyData["KES"] = 113.38
	currencyData["NGN"] = 412.9

	dest := currencyData[UserData["destinationCurrency"]]
	src := currencyData[UserData["sourceCurrency"]]

	amt, err := strconv.ParseFloat(UserData["amount"], 64)
	if err != nil || amt < 0 {
		return -1
	}

	result := float64(int(amt*(dest/src)*10000) / 10000)
	return result
}
