package controllers

import (
	"net/http"
	"strconv"

	"github.com/MuhammadFelda/task-manager/internal/services"
	"github.com/gin-gonic/gin"
)

type ProjectOwnerController struct {
	service *services.ProjectOwnerService
}

func NewProjectOwnerController() *ProjectOwnerController {
	return &ProjectOwnerController{
		service: services.NewProjectOwnerService(),
	}
}

func (ctrl *ProjectOwnerController) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	projectOwners, err := ctrl.service.GetAll(page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projectOwners})
}

func (ctrl *ProjectOwnerController) Create(c *gin.Context) {
	var req services.ProjectOwnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	projectOwner, err := ctrl.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projectOwner})
}

func (ctrl *ProjectOwnerController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	var req services.ProjectOwnerRequestUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	projectOwner, err := ctrl.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projectOwner})
}

func (ctrl *ProjectOwnerController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	if err := ctrl.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project Owner Deleted"})
}