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
	"github.com/spf13/cast"

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
	r.GET("/users", func(c *gin.Context) {
		users, err := h.GetAllUsers()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(200, gin.H{
			"users": users,
		})
	})
	r.GET("/:id", func(c *gin.Context) {

		id := c.Param("id")
		query := c.DefaultQuery("text", "Hello, I'm a bot")
		m, res := h.SendMessage(id, query)
		c.JSON(200, gin.H{
			"id":      id,
			"text":    query,
			"message": m,
			"result":  res,
		})
	})
	r.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		t, err := h.GetAllTodos(cast.ToInt64(id))
		if err != nil {
			c.JSON(401, gin.H{
				"err": err,
			})
		}
		c.JSON(200, gin.H{
			"todos": t,
		})
	})
	r.GET("/times", func(c *gin.Context) {
		times, err := h.GetAllNotificationTimes()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(200, gin.H{
			"times": times,
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
			for _, v := range n {

				if t.Hour() == v.Time.Hour() && t.Minute() == v.Time.Minute() && t.Second() == v.Time.Second() {
					h.SendTodo(v.ID)
				}
			}
			time.Sleep(1 * time.Minute)

		} else {
			time.Sleep(59*time.Minute - time.Duration(t.Minute())*time.Minute)
		}
	}
}
