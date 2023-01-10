package updates

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *Handler) GetAllDeadlineTimes(c *gin.Context) {
	deadlines, err := h.srvc.TodoService.GetAllDeadlineTimes()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"deadlines": deadlines,
	})
}

func Main(c *gin.Context) {
	c.String(200, "Hello, I'm a bot")
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *Handler) SendTextToUser(c *gin.Context) {

	id := c.Param("id")
	query := c.DefaultQuery("text", "Hello, I'm a bot")
	m, res := h.SendMessage(id, query)
	c.JSON(200, gin.H{
		"id":      id,
		"text":    query,
		"message": m,
		"result":  res,
	})
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.srvc.UserService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"users": users,
	})
}


func (h *Handler) GetAllNotificationTimes(c *gin.Context) {
	notification, err := h.srvc.TodoService.GetAllNotificationTimes()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"notification": notification,
	})
}

func (h *Handler) GetTodoById(c *gin.Context) {
	id := c.Param("id")
	t, err := h.srvc.TodoService.GetByID(id)
	if err != nil {
		c.JSON(401, gin.H{
			"err": err,
		})
	}
	c.JSON(200, gin.H{
		"todos": t,
	})
}

func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	query := c.DefaultQuery("done", "false")
	done := false
	if query == "true" {
		done = true
	}
	t, err := h.srvc.TodoService.GetAllByUserID(cast.ToInt64(id), done)
	if err != nil {
		c.JSON(401, gin.H{
			"err": err,
		})
	}
	c.JSON(200, gin.H{
		"todos": t,
	})
}
