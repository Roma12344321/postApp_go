package handler

import (
	"github.com/gin-gonic/gin"
	"postApp/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/api/auth")
	{
		auth.POST("/registration", h.createPerson)
	}
	return router
}
