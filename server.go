package main

import (
	"fmt"
	"github.com/inagacky/weather_linebot/weather/util"
	"log"
	"net/http"
)

func main() {

	envErr := env.Init()
	if envErr != nil {
		log.Panic("Environment Not Found: ", envErr)
	}

	http.HandleFunc("/report", castWeather)
	err := http.ListenAndServe(":" + env.GetEnv().Port, nil)
	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}

func castWeather(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hogehoge")

}