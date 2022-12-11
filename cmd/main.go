package main

import (
	"fmt"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"github.com/realtemirov/projects/tgbot/config"
	"github.com/realtemirov/projects/tgbot/service"
	"github.com/realtemirov/projects/tgbot/storage/postgres"
	"github.com/realtemirov/projects/tgbot/storage/redis"
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
	/*
		r := gin.Default()

		r.GET("/users", func(c *gin.Context) {
			users, err := s.UserService.GetAll()
			if err != nil {
				fmt.Println(err.Error())
			}
			c.JSON(200, users)
		})
		r.Run(":8080")
	*/

	fmt.Println("Bot is running")
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
