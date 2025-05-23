package dialogue

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DialogueHandler struct {
	service *DialogueService
}

func NewHandler(service *DialogueService) *DialogueHandler {
	return &DialogueHandler{service: service}
}

func (h *DialogueHandler) GetAvailableSpeaker(c *gin.Context) {
	idStr := c.Param("id")

	dialogueID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dialogue ID"})
		return
	}

	speakers, err := h.service.GetAvailableSpeakers(c, dialogueID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No speaker was found"})
		return
	}

	c.JSON(http.StatusOK, speakers)
}

func (h *DialogueHandler) GetDialogue(c *gin.Context) {
	var req DialogueRequest

	idStr := c.Param("id")

	c.ShouldBindJSON(&req)

	dialogueID, err := strconv.Atoi(idStr)

	var speaker string
	if req.Speaker != nil {
		speaker = *req.Speaker
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dialogue ID"})
		return
	}

	dialogue, err := h.service.GetDialogue(c, dialogueID, &speaker)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dialogue)
}
