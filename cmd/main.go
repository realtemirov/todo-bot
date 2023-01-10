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

	r.GET("/", u.Main)
	r.GET("/ping", u.Ping)
	r.GET("/users", h.GetAllUsers)
	r.GET("/:id", h.SendTextToUser)
	r.GET("/todos/:id", h.GetAllTodos)
	r.GET("/notifications", h.GetAllNotificationTimes)
	r.GET("/deadlines", h.GetAllDeadlineTimes)
	r.GET("/todo/:id", h.GetTodoById)
	r.GET("/user/:id", h.GetUserById)
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
			msg := tg.NewMessage(265943548, update.CallbackQuery.Data)
			bot.Send(msg)

			u.CallbackQuery(h, &update)
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
		t := time.Now().Add(time.Hour * 5)
		if t.Minute() == 0 {

			n, err := u.NotificationTimes(h)
			if err != nil {
				fmt.Println(err.Error())
			}
			d, err := u.DeadlineTimes(h)
			if err != nil {
				fmt.Println(err.Error())
			}

			for _, v := range n {

				if t.Hour() == v.Time.Hour() && t.Minute() == v.Time.Minute() && t.Second() == v.Time.Second() {
					h.SendTodo(v.ID)
				}
			}

			for _, v := range d {

				if int(t.Month()) == int(v.Time.Month()) && t.Day() == v.Time.Day() && t.Hour() == v.Time.Hour() && t.Minute() == v.Time.Minute() && t.Second() == v.Time.Second() {
					h.SendTodo(v.ID)
				}
			}

			time.Sleep(1 * time.Minute)

		} else {
			time.Sleep(59*time.Minute - time.Duration(t.Minute())*time.Minute)
		}
	}
}
