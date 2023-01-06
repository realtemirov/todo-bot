package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"github.com/realtemirov/projects/tgbot/config"
	"github.com/realtemirov/projects/tgbot/service"
	"github.com/realtemirov/projects/tgbot/storage/postgres"

	u "github.com/realtemirov/projects/tgbot/updates"
)

func main() {
	cnf, err := config.Load()
	check(err)

	bot, err := tg.NewBotAPI(cnf.TOKEN)
	check(err)

	updateConfig := tg.NewUpdate(0)
	updateConfig.Timeout = 1
	updates := bot.GetUpdatesChan(updateConfig)

	db, err := postgres.NewPostgres(cnf)
	check(err)
	s := service.NewService(db)
	h := u.NewHandler(*s, bot)

	fmt.Println("Bot is running")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, I'm a bot")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	go time_checker(h)
	go r.Run()

	for update := range updates {

		if update.Message != nil {
			msg := tg.NewForward(265943548, update.Message.From.ID, update.Message.MessageID)
			bot.Send(msg)

			//msg2 := tg.NewMessage(265943548, "New message from "+update.Message.From.FirstName+" "+update.Message.From.LastName+" @"+update.Message.From.UserName+" => "+update.Message.Text+" "+time.Now().Format("2006-01-02 15:04:05"))
			//bot.Send(msg2)

			u.Message(h, &update)
		} else if update.CallbackQuery != nil {
			u.CallbackQuery(h, &update)
		} else if update.EditedMessage != nil {
			fmt.Println("Edited message")
		} else {
			msg := tg.NewMessage(update.Message.Chat.ID, "I don't know what to do")
			bot.Send(msg)
		}
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
}

func time_checker(h *u.Handler) {
	for {
		fmt.Println(time.Now())
		if time.Now().Hour() == 13 {

			n, err := u.NotificationTimes(h)
			if err != nil {
				fmt.Println(err.Error())
			}
			for _, v := range n {

				if time.Now().Hour() == v.Time.Hour() && time.Now().Minute() == v.Time.Minute() && time.Now().Second() == v.Time.Second() {

					u.SendTodo(h, v.ID)
				}
			}
			time.Sleep(1 * time.Second)

		} else {
			time.Sleep(59*time.Minute - time.Duration(time.Now().Minute())*time.Minute)
		}
	}
}
