package module

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handlers/module.go
type ModuleHandler struct {
	moduleService *ModuleService
}

func NewModuleHandler(service *ModuleService) *ModuleHandler {
	return &ModuleHandler{moduleService: service}
}

func (h *ModuleHandler) GetAllModules(c *gin.Context) {
	modules, err := h.moduleService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No module was found"})
		return
	}

	c.JSON(http.StatusOK, modules)
}

func (h *ModuleHandler) GetVisibleModules(c *gin.Context) {
	modules, err := h.moduleService.GetVisible(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No visible module was found"})
		return
	}

	c.JSON(http.StatusOK, modules)
}

func (h *ModuleHandler) FindByID(c *gin.Context) {
	idStr := c.Param("id")

	moduleID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid module ID"})
		return
	}

	module, err := h.moduleService.GetByID(c.Request.Context(), moduleID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No module was found"})
		return
	}

	c.JSON(http.StatusOK, module)
}

// Create handles POST /modules requests to create a new module.
func (h *ModuleHandler) Create(c *gin.Context) {
	var req ModuleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	module := Module{
		Title:       req.Title,
		Description: req.Description,
		Visible:     req.Visible,
		CoverURL:    req.CoverURL,
	}

	createdModule, err := h.moduleService.Create(c.Request.Context(), &module)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create the new module"})
		return
	}

	c.JSON(http.StatusCreated, createdModule)
}
