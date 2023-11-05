package middlewares

import (
	"log"
	result "myapp/internal/utils"

	"github.com/gin-gonic/gin"
)

func PanicHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("发生了 panic:", err)

			str, ok := err.(string)
			if ok {
				c.JSON(500, result.Err.WithMsg(str))
			} else {
				c.JSON(500, result.Err.ToString())
			}
			c.Abort()
		}
	}()
	c.Next()
}
