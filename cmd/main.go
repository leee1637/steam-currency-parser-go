package main

import (
	"log"
	"steam-currency-parser/internal/bot"
	"steam-currency-parser/internal/config"
	"steam-currency-parser/internal/parser"
	"time"
)

func main() {
	cfg := config.Load()

	go func() {
		bot.StartBot(cfg.TelegramToken, cfg.AppID, cfg.MarketHashName)
	}()

	// Планирование обновления курсов в 1:00 по Московскому времени (UTC+3)
	for {
		now := time.Now()
		moscow := time.FixedZone("Moscow", 3*3600)
		moscowTime := now.In(moscow)

		next := time.Date(moscowTime.Year(), moscowTime.Month(), moscowTime.Day(), 1, 0, 0, 0, moscow)
		if moscowTime.After(next) {
			next = next.Add(24 * time.Hour)
		}

		log.Printf("Next update at: %v", next)

		time.Sleep(time.Until(next))

		log.Println("Fetching currency rates for scheduled update...")
		rates, err := parser.GetSteamCurrencyRates(cfg.AppID, cfg.MarketHashName)
		if err != nil {
			log.Printf("Error fetching currency rates: %v", err)
			continue
		}

		text := bot.FormatCurrencyMessage(rates)

		log.Println("Scheduled update:\n", text)
	}
}