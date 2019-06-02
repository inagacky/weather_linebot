package main

import (
	"github.com/inagacky/weather_linebot/weather/util"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"net/url"
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

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Unsupported HTTP method"))
		return
	}

	proxyURL, _ := url.Parse(env.GetEnv().FixieUrl)
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
	}

	bot, err := linebot.New(env.GetEnv().ChannelSecret, env.GetEnv().AccessToken, linebot.WithHTTPClient(client))

	if err != nil {
		log.Printf("[ERROR] linebot.New Fatal: %s", err)
		return
	}

	received, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			log.Printf("[ERROR] ErrInvalidSignature: %s", err)
			return
		}
		return
	}

	for _, event := range received {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
					return
				}
			}
		}
	}

}