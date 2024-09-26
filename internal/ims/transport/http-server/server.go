package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/transport"
	"log"
)

type Http struct {
	router *gin.Engine
	dataUC DataUC
	cfg    *transport.Config
}

func New(
	cfg *transport.Config,
	clientUC DataUC,
) *Http {
	return &Http{
		cfg:    cfg,
		router: gin.Default(),
		dataUC: clientUC,
	}
}

func (s *Http) Run() {
	s.registerHandlers()
	if err := s.router.Run(s.cfg.HTTP.Port); err != nil {
		log.Println("[IMS][TRANSPORT][HTTP-SERVER][RUN] ERROR ", err)
	}
}
