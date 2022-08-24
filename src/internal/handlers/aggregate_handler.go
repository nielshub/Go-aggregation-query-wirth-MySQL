package handlers

import (
	"contentSquare/src/internal/models"
	"contentSquare/src/internal/ports"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AggregateHandler struct {
	router *gin.RouterGroup
	DbRepo ports.DBRepository
}

func NewAggregateHandler(app *gin.RouterGroup, DbRepo ports.DBRepository) *AggregateHandler {
	aggregateAPI := &AggregateHandler{
		router: app,
		DbRepo: DbRepo,
	}
	aggregateAPI.router.GET("/count", aggregateAPI.countEventsWithFilter)
	aggregateAPI.router.GET("/count_distinct_users", aggregateAPI.countDistinctUsersWithFilter)
	return aggregateAPI
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
		fmt.Printf("Error counting events. Error: %s \n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error getting count values"})
		return
	}

	response := models.CountResponse{
		Count: strconv.Itoa(int(countValue)),
	}

	c.JSON(http.StatusOK, response)
}

func (h *AggregateHandler) countDistinctUsersWithFilter(c *gin.Context) {
	filters := models.Filters{
		DateFrom: c.Query("date_from"),
		DateTo:   c.Query("date_to"),
		Event:    c.Query("event"),
	}

	countValue, err := h.DbRepo.CountDistinctUsers(c, filters)
	if err != nil {
		fmt.Printf("Error counting distinct users. Error: %s \n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error counting distinct users"})
		return
	}

	response := models.CountDistinctUsersResponse{
		CountDistinctUsers: strconv.Itoa(int(countValue)),
	}

	c.JSON(http.StatusOK, response)
}
