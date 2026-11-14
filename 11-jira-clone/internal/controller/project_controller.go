package controller

import "11-jira-clone/internal/service"

type ProjectController struct {
	projectService *service.ProjectService
}

func NewProjectController(projectService *service.ProjectController) *ProjectController {
	return &ProjectController{projectService: projectService}
}

func (c *ProjectController) List(ctx *gin.Context) {
	projects, err := c.projectService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, projects)
}

func (c *ProjectController) Show(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	project, err := c.projectService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	ctx.JSON(http.StatusOK, project)
}