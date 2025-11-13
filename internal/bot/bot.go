package bot

import (
	"fmt"
	"log"
	"steam-currency-parser/internal/parser"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func FormatCurrencyMessage(currencies map[string]float64) string {
	var sb strings.Builder
	sb.WriteString("📊 *Steam Курсы (относительно USD)*\n\n")

	// Список валют с эмодзи
	mapping := map[string]string{
		"USD": "🇺🇸",
		"EUR": "🇪🇺",
		"RUB": "🇷🇺",
		"JPY": "🇯🇵",
		"CNY": "🇨🇳",
		"AED": "🇦🇪",
		"GBP": "🇬🇧",
		"CHF": "🇨🇭",
		"PLN": "🇵🇱",
		"BRL": "🇧🇷",
		"KRW": "🇰🇷",
		"TRY": "🇹🇷",
		"UAH": "🇺🇦",
		"CAD": "🇨🇦",
		"AUD": "🇦🇺",
		"NZD": "🇳🇿",
		"INR": "🇮🇳",
		"CLP": "🇨🇱",
		"PEN": "🇵🇪",
		"COP": "🇨🇴",
		"ZAR": "🇿🇦",
		"HKD": "🇭🇰",
		"TWD": "🇹🇼",
		"SAR": "🇸🇦",
		"ILS": "🇮🇱",
		"KZT": "🇰🇿",
		"KWD": "🇰🇼",
		"QAR": "🇶🇦",
		"CRC": "🇨🇷",
		"UYU": "🇺🇾",
	}

	for currency, rate := range currencies {
		emoji := mapping[currency]
		if emoji == "" {
			emoji = "💱"
		}
		sb.WriteString(fmt.Sprintf("%s %s: %.2f\n", emoji, currency, rate))
	}

	return sb.String()
}

func StartBot(token, appID, hashName string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я бот, который показывает курсы валют Steam.")
				bot.Send(msg)
			case "currency":
				log.Println("Fetching currency rates...")
				rates, err := parser.GetSteamCurrencyRates(appID, hashName)
				if err != nil {
					log.Printf("Error fetching currency rates: %v", err)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось получить курсы валют.")
					bot.Send(msg)
					continue
				}
				text := FormatCurrencyMessage(rates)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				msg.ParseMode = "Markdown"
				bot.Send(msg)
			}
		}
	}
}