package telegram

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
)

const pass = "11"
const token = "5838671444:AAGjbaYQiPFOoJVChKEf5bGfveKmMqWcHQs"

var users = make([]int64, 10, 100)

func Start(wg sync.WaitGroup, news [][]string) {
	defer wg.Done()
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	handlerFunc(bot, updates, news)
}

func handlerFunc(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel, news [][]string) {
	for update := range updates {
		flag := false
		if update.Message != nil {
			// If we got a message
			log.Printf("[%s] %d %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)
		}
		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("И снова здравствуйте, %s! \n\nДля того чтобы получить доступ к боту введите пароль.", update.Message.From.UserName))
			bot.Send(msg)
		default:
			for _, user := range users {
				if update.Message.From.ID == int64(user) {
					flag = true
					break
				}
			}
			if !flag {
				if update.Message.Text != pass {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неверный пароль.")
					bot.Send(msg)
				} else {
					users = append(users, update.Message.From.ID)
					flag = false
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Доступ разрешен.\n\nВведите дату в формате dd.mm.yyyy"))
					bot.Send(msg)
				}
				break
			}
			for i, n := range news {
				if n[0] == update.Message.Text {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n\n%s", n[0], n[1]))
					bot.Send(msg)
					break
				}
				if i == len(news)-1 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Новостей по данной дате не найдено!")
					bot.Send(msg)
				}
			}
		}

	}
}
