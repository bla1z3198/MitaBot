package mitacom

import (
	"log"
	"os"
	"strconv"
	"time"

	weather "github.com/3crabs/go-yandex-weather-api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Wr(mita *tgbotapi.BotAPI, update tgbotapi.Update) {
	yandexWeatherApiKey := "Yandex-API-Key"
	wr, err := weather.GetWeatherWithCache(yandexWeatherApiKey, 59.6, 30.2, time.Hour)
	if err != nil {
		log.Fatalln("error 003")
	}
	cond := wr.Fact.GetCondition()
	facttemp := wr.Fact.Temp
	feelslike := wr.Fact.FeelsLike
	facttempstr := strconv.Itoa(facttemp)
	feelslikestr := strconv.Itoa(feelslike)
	m := "держи погоду на сегодня: " + cond + "\n" + "температура: " + facttempstr + " °C" + "\n" + "ощущается как: " + feelslikestr + " °C" + "\n" + "по данным сервиса Яндекс Погода"
	msgwr := tgbotapi.NewMessage(update.Message.Chat.ID, m)
	msgwr.ReplyToMessageID = update.Message.MessageID
	mita.Send(msgwr)
}

func Rse(grp string, mita *tgbotapi.BotAPI, update tgbotapi.Update) {
	d1, d2 := Dd()
	var d22 string
	var dfull string
	if d2 {
		d22 = "1"
		dfull = d1 + d22
	} else {
		d22 = "0"
		dfull = d1 + d22
	}
	fContent, err := os.ReadFile("images/" + grp + "/" + dfull + ".txt")
	if err != nil {
		panic(err)
	}

	r := (string(fContent))
	msgg := tgbotapi.NewMessage(update.Message.Chat.ID, "Расписание на сегодня:\n"+r)
	mita.Send(msgg)
}

func Any(mita *tgbotapi.BotAPI, update tgbotapi.Update, dev int) {
	if dev == 0 {
		msgwr := tgbotapi.NewMessage(update.Message.Chat.ID, "прости, не поняла тебя (")
		msgwr.ReplyToMessageID = update.Message.MessageID
		mita.Send(msgwr)
	} else if dev == 1 {
		msgwr := tgbotapi.NewMessage(update.Message.Chat.ID, "прости, не поняла тебя (")
		msgwr.ReplyToMessageID = update.Message.MessageID
		mita.Send(msgwr)
	}

}

func Dd() (string, bool) {
	now := time.Now()

	dayOfWeek := now.Weekday()
	var dayAbbr string

	switch dayOfWeek {
	case time.Monday:
		dayAbbr = "mon"
	case time.Tuesday:
		dayAbbr = "tue"
	case time.Wednesday:
		dayAbbr = "wed"
	case time.Thursday:
		dayAbbr = "thu"
	case time.Friday:
		dayAbbr = "fri"
	case time.Saturday:
		dayAbbr = "sat"
	case time.Sunday:
		dayAbbr = "sun"
	}

	_, week := now.ISOWeek()
	isEvenWeek := week%2 == 0

	return dayAbbr, isEvenWeek
}

func Start(mita *tgbotapi.BotAPI, update tgbotapi.Update) {
	file, err := os.Open("images/hi.jpg")
	if err != nil {
		log.Fatalln("error 005")
	}
	reader := tgbotapi.FileReader{Name: "hi.jpg", Reader: file}
	photo := tgbotapi.NewPhoto(update.Message.Chat.ID, reader)
	photo.Caption = "Привет! меня зовут Мита, я немного постараюсь помочь тебе в учёбе ;)"
	mita.Send(photo)
}
