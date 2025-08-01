package api

import (
    "fmt"
    "github.com/gin-gonic/gin"
	. "todo/internal/controllers"
)

func Serve(port int) error {
    r := gin.Default()
    RegisterRoutes(r)
    return r.Run(fmt.Sprintf(":%d", port))
}

func RegisterRoutes(r *gin.Engine){
    r.GET("/tasks", GetAllTasksHandler)
    r.GET("/tasks/:id", GetTaskByIdHandler)
    r.POST("/tasks", CreateTaskHandler)
    r.PUT("/tasks/:id", UpdateTaskHandler)
    r.DELETE("/tasks/:id", DeleteTaskHandler)
    r.POST("/tasks/:id/complete", CompleteTaskHandler)
	r.GET("/tasks/stats", GetTaskStatsHandler)
}

