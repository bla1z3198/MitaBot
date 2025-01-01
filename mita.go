package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	weather "github.com/3crabs/go-yandex-weather-api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("7561424171:AAF3-62oAWd876eUdQF56x-dsX5_xq8_jJE")
	if err != nil {
		log.Fatal("error 001")
	}

	bot.Debug = true

	log.Printf("working on username: %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.Text == "/start" {
				msghi := tgbotapi.NewMessage(update.Message.Chat.ID, "привет, я Мита!\nпопробуй нажать 🡺 /fx")
				msghi.ReplyToMessageID = update.Message.MessageID
				bot.Send(msghi)
			} else if update.Message.Text == "/fx" {
				msgfx := tgbotapi.NewMessage(update.Message.Chat.ID, "1) покажу тебе расписание 🡺 /r\n2) подскажу погоду 🡺 /wr\n3) расскажу про транспорт до университета 🡺 /rd\n4) погадаю тебе на сегодняшний день 🡺 /tr")
				msgfx.ReplyToMessageID = update.Message.MessageID
				bot.Send(msgfx)
			} else if update.Message.Text == "/dev" {
				msgdev := tgbotapi.NewMessage(update.Message.Chat.ID, "Mita bot v6 by Mirov Yan, 2024\ngithub.com/bla1z3198/MitaBot")
				msgdev.ReplyToMessageID = update.Message.MessageID
				bot.Send(msgdev)
			} else if update.Message.Text == "/r" {
				msgFAK := tgbotapi.NewMessage(update.Message.Chat.ID, "выбери свой факультет\nФТСБ 🡺 /ftsb\nФЛЭ 🡺 /fla\nФУЭП 🡺 /fup")
				msgFAK.ReplyToMessageID = update.Message.MessageID
				bot.Send(msgFAK)
			} else if update.Message.Text == "/wr" {
				m := wr()
				msgwr := tgbotapi.NewMessage(update.Message.Chat.ID, m)
				msgwr.ReplyToMessageID = update.Message.MessageID
				bot.Send(msgwr)
				msgfxa := tgbotapi.NewMessage(update.Message.Chat.ID, "чем ещё могу помочь?\nсмело жми /fx\nя всегда рядом ;)")
				msgfxa.ReplyToMessageID = update.Message.MessageID
				bot.Send(msgfxa)

			} else if update.Message.Text == "/rd" {
				file, _ := os.Open("rd.png")
				reader := tgbotapi.FileReader{Name: "rd", Reader: file}
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, reader)
				photo.Caption = "автобусы которые тебе понадобятся 🡹🡹🡹"
				bot.Send(photo)
				msgfxa := tgbotapi.NewMessage(update.Message.Chat.ID, "чем ещё могу помочь?\nсмело жми /fx\nя всегда рядом ;)")
				msgfxa.ReplyToMessageID = update.Message.MessageID
				bot.Send(msgfxa)
			} else if update.Message.Text == "/ftsb" {
				msgGROUP := tgbotapi.NewMessage(update.Message.Chat.ID, "выбери свою учебную группу\nАСУВД-24-01 🡺 /asuvd2401\nМПОБАС-24-01 🡺 /mpobas2401\nТЛ-24-01 🡺 /tl2401\nОАБ-24-01 🡺 /oab2401\nОБПрВТ-24-01 🡺 /obprvt2401")
				msgGROUP.ReplyToMessageID = update.Message.MessageID
				bot.Send(msgGROUP)
			} else if update.Message.Text == "/asuvd2401" {
				fkl := "/ftsb"
				grp := update.Message.Text
				var ph = timetable(fkl, grp)
				pht(bot, update, ph)
			} else if update.Message.Text == "/tl2401" {
				fkl := "/ftsb"
				grp := update.Message.Text
				var ph = timetable(fkl, grp)
				pht(bot, update, ph)
			} else if update.Message.Text == "/oab2401" {
				fkl := "/ftsb"
				grp := update.Message.Text
				var ph = timetable(fkl, grp)
				pht(bot, update, ph)
			} else if update.Message.Text == "/mpobas2401" {
				fkl := "/ftsb"
				grp := update.Message.Text
				var ph = timetable(fkl, grp)
				pht(bot, update, ph)
			} else if update.Message.Text == "/obprvt2401" {
				fkl := "/ftsb"
				grp := update.Message.Text
				var ph = timetable(fkl, grp)
				pht(bot, update, ph)
			} else if update.Message.Text == "/fla" {
				msgfla := tgbotapi.NewMessage(update.Message.Chat.ID, "это содержимое пока не доступно =(")
				msgfla.ReplyToMessageID = update.Message.MessageID
				bot.Send(msgfla)
			} else if update.Message.Text == "/fup" {
				msgfla := tgbotapi.NewMessage(update.Message.Chat.ID, "это содержимое пока не доступно =(")
				msgfla.ReplyToMessageID = update.Message.MessageID
				bot.Send(msgfla)
			} else {
				msgun := tgbotapi.NewMessage(update.Message.Chat.ID, "не поняла тебя, прости(")
				msgun.ReplyToMessageID = update.Message.MessageID
				bot.Send(msgun)
			}

		}
	}

}

func timetable(fkl string, grp string) string {
	fklt := strings.Trim(fkl, "/")
	grpt := strings.Trim(grp, "/")
	var ph string
	day := time.Now()
	d := day.String()
	r := " "
	index := strings.IndexAny(d, r)
	if index != -1 {
		d = d[:index]
	}

	if d == "2025-02-03" || d == "2025-02-17" || d == "2025-03-03" || d == "2025-03-17" || d == "2025-03-31" {
		ph = "source/" + fklt + "/" + grpt + "/" + "mon1" + ".png"
	} else if d == "2025-02-04" || d == "2025-02-18" || d == "2025-03-04" || d == "2025-03-18" {
		ph = "source/" + fklt + "/" + grpt + "/" + "tue1" + ".png"
	} else if d == "2025-01-22" || d == "2025-02-05" || d == "2025-02-19" || d == "2025-03-05" || d == "2025-03-19" {
		ph = "source/" + fklt + "/" + grpt + "/" + "wed1" + ".png"
	} else if d == "2025-01-23" || d == "2025-02-06" || d == "2025-02-20" || d == "2025-03-06" || d == "2025-03-20" {
		ph = "source/" + fklt + "/" + grpt + "/" + "thu1" + ".png"
	} else if d == "2025-01-24" || d == "2025-02-07" || d == "2025-02-21" || d == "2025-03-07" || d == "2025-03-21" {
		ph = "source/" + fklt + "/" + grpt + "/" + "fri1" + ".png"
	} else if d == "2025-01-27" || d == "2025-02-10" || d == "2025-02-24" || d == "2025-03-10" || d == "2025-03-24" {
		ph = "source/" + fklt + "/" + grpt + "/" + "mon2" + ".png"
	} else if d == "2025-01-28" || d == "2025-02-11" || d == "2025-02-25" || d == "2025-03-11" || d == "2025-03-25" {
		ph = "source/" + fklt + "/" + grpt + "/" + "tue2" + ".png"
	} else if d == "2025-01-29" || d == "2025-02-12" || d == "2025-02-26" || d == "2025-03-12" || d == "2025-03-26" {
		ph = "source/" + fklt + "/" + grpt + "/" + "wed2" + ".png"
	} else if d == "2025-01-30" || d == "2025-02-13" || d == "2025-02-27" || d == "2025-03-13" || d == "2025-03-27" {
		ph = "source/" + fklt + "/" + grpt + "/" + "thu2" + ".png"
	} else if d == "2025-01-31" || d == "2025-02-14" || d == "2025-02-28" || d == "2025-03-14" || d == "2025-03-28" {
		ph = "source/" + fklt + "/" + grpt + "/" + "fri2" + ".png"
	} else if d == "2025-02-01" || d == "2025-02-15" || d == "2025-03-01" || d == "2025-03-15" || d == "2025-03-29" {
		ph = "source/" + fklt + "/" + grpt + "/" + "sat2" + ".png"
	} else if d == "2025-01-25" || d == "2025-02-08" || d == "2025-02-22" || d == "2025-03-22" {
		ph = "source/" + fklt + "/" + grpt + "/" + "sat1" + ".png"
	}

	return ph
}

func wr() string {
	yandexWeatherApiKey := "7eb05bd6-d6b0-4b06-961c-29d722d3d31c"
	w, err := weather.GetWeatherWithCache(yandexWeatherApiKey, 59.6, 30.2, time.Hour)
	if err != nil {
		log.Fatal("error 003")
	}
	cond := w.Fact.GetCondition()
	facttemp := w.Fact.Temp
	feelslike := w.Fact.FeelsLike
	facttempstr := strconv.Itoa(facttemp)
	feelslikestr := strconv.Itoa(feelslike)
	if facttemp < 0 {
		m := "держи погоду на сегодня: " + cond + "\n" + "температура: " + facttempstr + " °C" + "\n" + "будь внимателенее под ногами может быть лёд!" + "\n" + "ощущается как: " + feelslikestr + " °C" + "\n" + "по данным сервиса Яндекс Погода"
		return m
	} else {
		m := "держи погоду на сегодня: " + cond + "\n" + "температура: " + facttempstr + " °C" + "\n" + "ощущается как: " + feelslikestr + " °C" + "\n" + "по данным сервиса Яндекс Погода"
		return m
	}
}

func pht(bot *tgbotapi.BotAPI, update tgbotapi.Update, ph string) {
	file, err := os.Open(ph)
	if err != nil {
		msg3 := tgbotapi.NewMessage(update.Message.Chat.ID, "ооо, сегодня же у тебя выходной, отдохни как следует!")
		msg3.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg3)
	}
	reader := tgbotapi.FileReader{Name: ph, Reader: file}
	photo := tgbotapi.NewPhoto(update.Message.Chat.ID, reader)
	photo.Caption = "твоё расписание на сегодня, успехов в учёбе!"
	bot.Send(photo)
	msgfxa := tgbotapi.NewMessage(update.Message.Chat.ID, "чем ещё могу помочь?\nсмело жми /fx\nя всегда рядом ;)")
	msgfxa.ReplyToMessageID = update.Message.MessageID
	bot.Send(msgfxa)
}
