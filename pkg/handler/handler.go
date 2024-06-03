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
	api := router.Group("/api", h.personIdentity)
	{
		auth := api.Group("/auth")
		{
			auth.POST("/registration", h.createPerson)
			auth.POST("/login", h.logIn)
		}
	}
	return router
}
