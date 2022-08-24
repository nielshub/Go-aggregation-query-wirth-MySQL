package handlers

import (
	"contentSquare/src/internal/models"
	"contentSquare/src/internal/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AggregateHandler struct {
	router *gin.RouterGroup
	DbRepo ports.DBRepository
}

func NewAggregateHandler(app *gin.RouterGroup, DbRepo ports.DBRepository) *AggregateHandler {
	h := &AggregateHandler{
		router: app,
		DbRepo: DbRepo,
	}
	h.router.GET("/count", h.countEventsWithFilter)
	return h
}

func (h *AggregateHandler) countEventsWithFilter(c *gin.Context) {
	filters := models.Filters{
		UserId:   c.Query("user_id"),
		DateFrom: c.Query("date_from"),
		DateTo:   c.Query("date_to"),
		Event:    c.Query("event"),
	}

	countValue, err := h.DbRepo.CountEvents(c, filters)
	if err != nil {
		//log.Logger.Error().Msgf("Error getting users. Error: %s", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error getting users"})
		return
	}

	response := models.CountResponse{
		Count: strconv.Itoa(int(countValue)),
	}

	c.JSON(http.StatusOK, response)
}
