package api

import (
    "fmt"
    "github.com/gin-gonic/gin"
	t "todo/internal/task"
)

func Serve(port int) error {
    r := gin.Default()
    RegisterRoutes(r)
    return r.Run(fmt.Sprintf(":%d", port))
}

func RegisterRoutes(r *gin.Engine){
    r.GET("/tasks", t.GetAllTasksHandler)
    r.GET("/tasks/:id", t.GetTaskByIdHandler)
    r.POST("/tasks", t.CreateTaskHandler)
    r.PUT("/tasks/:id", t.UpdateTaskHandler)
    r.DELETE("/tasks/:id", t.DeleteTaskHandler)
    r.POST("/tasks/:id/complete", t.CompleteTaskHandler)
	r.GET("/tasks/stats", t.GetTaskStatsHandler)
}

