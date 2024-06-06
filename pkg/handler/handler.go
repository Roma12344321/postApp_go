package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	auth := router.Group("/auth")
	{
		auth.POST("/registration", h.createPerson)
		auth.POST("/login", h.logIn)
	}
	api := router.Group("/api", h.personIdentity)
	{
		api.GET("/test", func(c *gin.Context) {
			id, err := getPersonId(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
			c.JSON(http.StatusOK, gin.H{"id": id})
		})
	}
	return router
}
