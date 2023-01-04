package main

import (
	"fmt"
	"log"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"github.com/realtemirov/projects/tgbot/config"
	"github.com/realtemirov/projects/tgbot/service"
	"github.com/realtemirov/projects/tgbot/storage/postgres"
	"github.com/realtemirov/projects/tgbot/storage/redis"
	"github.com/realtemirov/projects/tgbot/updates"
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
	rds, err := redis.NewRedis(cnf)
	check(err)

	s := service.NewService(db)
	h := u.NewHandler(*s, rds, bot)

	fmt.Println("Bot is running")

	go time_checker(h)
	for update := range updates {

		if update.Message != nil {
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

func time_checker(h *updates.Handler) {
	for {
		if time.Now().Minute() == 0 {
			n, err := u.NotificationTimes(h)
			if err != nil {
				fmt.Println(err.Error())
			}
			for _, v := range n {

				if time.Now().Hour() == v.Notif_date.Hour() && time.Now().Minute() == v.Notif_date.Minute() && time.Now().Second() == v.Notif_date.Second() {
					u.SendTodo(h, v.Todo_ID)
				}
			}
			time.Sleep(1 * time.Second)

		} else {
			time.Sleep(59*time.Minute - time.Duration(time.Now().Minute())*time.Minute)
		}
	}
}
