package http_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// docsEntrypointPing пинг сервиса ims
// @Summary Отправляет пинг на сервис ims для проверки ответа
// @Description Отправляет пинг на сервис ims для проверки ответа
// @Tags dataUC
// @Accept json
// @Produce json
// @Success 200 "Сообщение об успешной отправке на расчет"
// @Failure 400 "Ошибка при отправке данных"
// @Router /ping [get]
func (s *Http) docsEntrypointPing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "pong",
		"server time": time.Now().String(),
	})
}
