package controllers

import (
	"net/http"
	"strconv"

	"github.com/MuhammadFelda/task-manager/internal/services"
	"github.com/gin-gonic/gin"
)

type SkillController struct {
	service *services.SkillService
}

func NewSkillController() *SkillController {
	return &SkillController{
		service: services.NewSkillService(),
	}
}

func (ctrl *SkillController) Index(c *gin.Context) {
	skills, err := ctrl.service.GetAll(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": skills})
}

func (ctrl *SkillController) Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	skill, err := ctrl.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Skill Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": skill})
}

func (ctrl *SkillController) Create(c *gin.Context) {
	var req services.SkillCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	skill, err := ctrl.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": skill})
}

func (ctrl *SkillController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	if err := ctrl.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skill Deleted"})
}