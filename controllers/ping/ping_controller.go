package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//全てのHadler or Controllerは c *gin.Contextをインターフェースに持つ
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
