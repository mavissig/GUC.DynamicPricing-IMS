package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/domain"
	"github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/transport/common"
	"net/http"
)

type DataUC interface {
	DataSet(data *domain.Data) error
	DataGetByKey(key uuid.UUID) (*domain.Data, error)
}

// dataEntrypointSet отправляет данные в сервис для расчета
// @Summary Отправить данные на расчет
// @Description Отправляет данные в сервис для рассвета
// @Tags dataUC
// @Accept json
// @Producer json
// @Param   dataUC body domain.Client true "входные данные для расчета"
// @Success 200 "Сообщение об успешной отправке на расчет"
// @Failure 400 "Ошибка при отправке данных"
// @Router /dataUC [post]
func (s *Http) dataEntrypointSet(c *gin.Context) {
	var jsonData domain.Data
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("[IMS][TRANSPORT][HTTP-SERVER][DATA] ERROR %v", err),
		})
		return
	}

	err := s.dataUC.DataSet(&jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("[IMS][TRANSPORT][HTTP-SERVER][DATA] ERROR %v", err),
		})
	}
}

func (s *Http) dataEntrypointGetByKey(c *gin.Context) {
	key := c.Query("id")

	id, err := uuid.Parse(key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	res, err := s.dataUC.DataGetByKey(id)
	if err != nil {
		c.JSON(common.ParseErrToHttpStatus(err.Error()), gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, res)
}
