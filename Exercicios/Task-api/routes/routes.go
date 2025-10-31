package routes

import (
	"task-api/handlers"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, h *handlers.TaskHandler) {
	api := r.Group("/api")
	{
		t := api.Group("/tasks")
		{
			t.POST("/", h.CreateTask)
			t.GET("/", h.GetTasks)
			t.GET("/:id", h.GetTask)
			t.PUT("/:id", h.UpdateTask)
			t.DELETE("/:id", h.DeleteTask)
		}
	}
}
