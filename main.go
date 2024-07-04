package main

import (
	"errors"

	"github.com/Zhima-Mochi/easy-task-api/app"
	"github.com/Zhima-Mochi/easy-task-api/app/dto"
	doc "github.com/Zhima-Mochi/easy-task-api/docs"
	"github.com/Zhima-Mochi/easy-task-api/domain/repo"
	"github.com/Zhima-Mochi/easy-task-api/domain/service"
	"github.com/Zhima-Mochi/easy-task-api/infra/persistence"
	"github.com/Zhima-Mochi/easy-task-api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type handler struct {
	taskAppService app.TaskAppService
}

func NewHandler(taskAppService app.TaskAppService) *handler {
	return &handler{taskAppService: taskAppService}
}

// @Summary Get all tasks
// @Description Get all tasks
// @Accept json
// @Produce json
// @Success 200 {array} dto.TaskResponse
// @Failure 500 {object} Error
// @Router /tasks [get]
func (h *handler) GetAllTask(c *gin.Context) {
	tasks, err := h.taskAppService.GetAllTask(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tasks)

}

// @Summary Get task by ID
// @Description Get task by ID
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} dto.TaskResponse
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /tasks/{id} [get]
func (h *handler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.taskAppService.GetTaskByID(c, id)
	if err != nil {
		if errors.Is(err, service.ErrorTaskNotFound) {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, task)
}

// @Summary Create a task
// @Description Create a task
// @Accept json
// @Produce json
// @Param request body dto.TaskCreateRequest true "Task Create Request"
// @Success 201 {object} dto.TaskCreateResponse
// @Failure 400 {object} Error
// @Failure 409 {object} Error
// @Failure 500 {object} Error
// @Router /tasks [post]
func (h *handler) CreateTask(c *gin.Context) {
	var req dto.TaskCreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task, err := h.taskAppService.CreateTask(c, &req)
	if err != nil {
		if errors.Is(err, service.ErrorTaskAlreadyExists) {
			c.JSON(409, gin.H{"error": err.Error()})
			return
		}

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, task)
}

// @Summary Update a task
// @Description Update a task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param request body dto.TaskUpdateRequest true "Task Update Request"
// @Success 204
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /tasks/{id} [put]
func (h *handler) UpdateTask(c *gin.Context) {
	var req dto.TaskUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.taskAppService.UpdateTask(c, &req)
	if err != nil {
		if errors.Is(err, service.ErrorTaskNotFound) {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}

// @Summary Delete a task
// @Description Delete a task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 204
// @Failure 500 {object} Error
// @Router /tasks/{id} [delete]
func (h *handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := h.taskAppService.DeleteTask(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			logrus.Fatalf("panic: %v", r)
		}
	}()

	persistence := persistence.NewPersistence()
	taskRepo := repo.NewTaskRepository(persistence)
	taskService := service.NewTaskService(taskRepo)
	taskAppService := app.NewTaskAppService(taskService)
	handler := NewHandler(taskAppService)

	router := gin.Default()
	router.Use(middleware.TraceMiddleware())

	router.GET("/tasks", handler.GetAllTask)
	router.GET("/tasks/:id", handler.GetTaskByID)
	router.POST("/tasks", handler.CreateTask)
	router.PUT("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.DeleteTask)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// swagger doc
	doc.SwaggerInfo.Title = "Easy Task API"
	doc.SwaggerInfo.Description = "This is a simple task API."
	doc.SwaggerInfo.Version = "1.0"

	router.Run(":8080")
}
