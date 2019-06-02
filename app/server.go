package main

import (
	"fmt"
	"github.com/inagacky/weather_linebot/app/util"
	"log"
	"net/http"
)

func main() {

	envErr := env.Init()
	if envErr != nil {
		log.Panic("Environment Not Found: ", envErr)
	}

	http.HandleFunc("/cast", castWeather)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}

func castWeather(w http.ResponseWriter, r *http.Request) {

	env := env.GetEnv()
	fmt.Fprintf(w, env.AccessToken)

}