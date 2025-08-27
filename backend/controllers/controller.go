package controllers

import (
	"net/http"
	"blazestack/models"
	"github.com/gin-gonic/gin"
)

func CreateIncident(c *gin.Context) {
	var incident *models.Incident
	if err := c.ShouldBindJSON(&incident); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.LastID = models.LastID + 1
	incident.ID = models.LastID
	models.DB[models.LastID] = incident

	c.JSON(http.StatusOK, gin.H{"message": "Incident created"})	
}

func ListIncidents(c *gin.Context) {
	c.JSON(http.StatusOK, models.DB)
}