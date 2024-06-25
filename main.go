package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type FormData struct {
	Name      string
	DobDay    int
	DobMonth  int
	DobYear   int
	AgeYears  int
	AgeMonths int
	AgeDays   int
	Zodiac    string
}

var tmpl = template.Must(template.ParseFiles("form.html"))

func main() {
	http.HandleFunc("/", formHandler)
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		dobDay, _ := strconv.Atoi(r.FormValue("dob_day"))
		dobMonth, _ := strconv.Atoi(r.FormValue("dob_month"))
		dobYear, _ := strconv.Atoi(r.FormValue("dob_year"))

		age := calculateAge(dobDay, dobMonth, dobYear)
		zodiac := getZodiac(dobDay, dobMonth)

		data := FormData{
			Name:      name,
			DobDay:    dobDay,
			DobMonth:  dobMonth,
			DobYear:   dobYear,
			AgeYears:  age["years"],
			AgeMonths: age["months"],
			AgeDays:   age["days"],
			Zodiac:    zodiac,
		}
		tmpl.Execute(w, data)
		return
	}
	tmpl.Execute(w, nil)
}

func calculateAge(day, month, year int) map[string]int {
	birthDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	today := time.Now()
	age := today.Sub(birthDate)

	years := int(age.Hours() / 24 / 365)
	months := int(age.Hours() / 24 / 30)
	days := int(age.Hours() / 24)

	return map[string]int{
		"years":  years,
		"months": months % 12,
		"days":   days % 30,
	}
}

func getZodiac(day, month int) string {
	switch {
	case (month == 12 && day >= 22) || (month == 1 && day <= 20):
		return "Capricorn"
	case (month == 1 && day >= 21) || (month == 2 && day <= 19):
		return "Aquarius"
	case (month == 2 && day >= 20) || (month == 3 && day <= 20):
		return "Pisces"
	case (month == 3 && day >= 21) || (month == 4 && day <= 20):
		return "Aries"
	case (month == 4 && day >= 21) || (month == 5 && day <= 21):
		return "Taurus"
	case (month == 5 && day >= 22) || (month == 6 && day <= 21):
		return "Gemini"
	case (month == 6 && day >= 22) || (month == 7 && day <= 22):
		return "Cancer"
	case (month == 7 && day >= 23) || (month == 8 && day <= 23):
		return "Leo"
	case (month == 8 && day >= 24) || (month == 9 && day <= 23):
		return "Virgo"
	case (month == 9 && day >= 24) || (month == 10 && day <= 23):
		return "Libra"
	case (month == 10 && day >= 24) || (month == 11 && day <= 22):
		return "Scorpio"
	case (month == 11 && day >= 23) || (month == 12 && day <= 21):
		return "Sagittarius"
	default:
		return "Unknown"
	}
}
