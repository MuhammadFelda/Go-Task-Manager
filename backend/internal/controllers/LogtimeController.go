package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/MuhammadFelda/task-manager/internal/services"
	"github.com/gin-gonic/gin"
)

type LogtimeController struct {
	service *services.LogtimeService
}

func NewLogtimeController(service *services.LogtimeService) *LogtimeController {
	return &LogtimeController{service: service}
}

func (ctrl *LogtimeController) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	logtimes, total, err := ctrl.service.GetAll(
		c.Query("user_id"),
		c.Query("from"),
		c.Query("to"),
		page,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":			logtimes,
		"total":		total,
		"current_page":	page,
		"last_page":	math.Ceil(float64(page) / 10),
	})
}

func (ctrl *LogtimeController) Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	logtime, err := ctrl.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Logtime Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": logtime})
}

func (ctrl *LogtimeController) Create(c *gin.Context) {
	var req services.LogtimeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	logtime, err := ctrl.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": logtime})
}

func (ctrl *LogtimeController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	if err := ctrl.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logtime Deleted"})
}