package main

import (
	"log"
	"main/mitacom"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	token           = "token"
	numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Какая сегодня погода?")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Какое сегодня расписание?"),
		),
	)
	nk = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("АСУВД 24-01")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ТЛ 24-01")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("МПОБАС 24-01")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ОБПрВТ 24-01")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ОАБ 24-01")),
	)
)

func main() {
	mita, err := tgbotapi.NewBotAPI(token) // Start Mita
	if err != nil {
		panic("error 001 - can't start the bot")
	}

	mita.Debug = true

	log.Printf("Authorized on account %s", mita.Self.UserName) // Started at the any Mita username

	upd := tgbotapi.NewUpdate(0) // Set up offset
	upd.Timeout = 40             // Updates in "n" seconds

	updates := mita.GetUpdatesChan(upd) // Updates - updating chat updates
	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "вот что я могу")

		switch update.Message.Text {
		case "/start":
			mitacom.Start(mita, update)
			msg.ReplyMarkup = numericKeyboard
			mita.Send(msg)
		case "Какая сегодня погода?":
			mitacom.Wr(mita, update)
		case "Какое сегодня расписание?":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выбери свою группу")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			msg.ReplyMarkup = nk
			mita.Send(msg)
		case "АСУВД 24-01", "ТЛ 24-01", "МПОБАС 24-01", "ОБПрВТ 24-01", "ОАБ 24-01":
			grp := update.Message.Text
			mitacom.Rse(grp, mita, update)
		default:
			mitacom.Any(mita, update)
		}
	}
}
