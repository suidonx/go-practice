package controller

import "11-jira-clone/internal/service"

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController(taskService *service.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (c *TaskController) List(ctx *gin.Context) {
	tasks, err := c.taskService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, projects)
}

func (c *TaskController) Show(ctx *gin.Context) {
	idStr := ctx.Param("id")
	projectIDStr := ctx.Param("project_id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	projectID, err := strconv.ParseUint(projectIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	task, err := c.taskService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}