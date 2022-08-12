package handlers

import (
	"contentSquare/src/internal/models"
	"contentSquare/src/internal/ports"

	"github.com/gin-gonic/gin"
)

type IngestHandler struct {
	router *gin.RouterGroup
	DbRepo ports.DBRepository
}

func NewIngestHandler(app *gin.RouterGroup, DbRepo ports.DBRepository) *IngestHandler {
	h := &IngestHandler{
		router: app,
		DbRepo: DbRepo,
	}
	h.router.PUT("/ingest", h.ingest)
	return h
}

func (h *IngestHandler) ingest(c *gin.Context) {
	var ingest models.Ingest
	if err := c.ShouldBindJSON(&ingest); err != nil || ingest.Path == "" {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	err := h.DbRepo.IngestFileData(c, ingest.Path)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Issue ingesting file",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Ingestion has been done correctly",
	})
}
