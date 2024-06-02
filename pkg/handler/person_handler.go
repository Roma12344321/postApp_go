package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"postApp"
)

func (h *Handler) createPerson(c *gin.Context) {
	var person postApp.Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	id, err := h.service.CreatePerson(&person)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
