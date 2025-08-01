package todo

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func ServeApi(port int) error {
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
}

func GetAllTasksHandler(c *gin.Context){
    tasks, err := ListTasks(false, true)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, tasks)
}

func GetTaskByIdHandler(c *gin.Context){
    id := c.Param("id")
    task, err := GetTaskById(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, task)
}

func CreateTaskHandler(c *gin.Context){
    var taskInput TaskCreateInput
    if err := c.ShouldBindJSON(&taskInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    task, err := AddTask(taskInput.Title, taskInput.Done)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "id": task.Id,
    })
}

func UpdateTaskHandler(c *gin.Context){
    id := c.Param("id")
    var taskInput TaskUpdateInput
    if err := c.ShouldBindJSON(&taskInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    task, err := UpdateTask(id, taskInput.Title, taskInput.Done)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, task)
}

func DeleteTaskHandler(c *gin.Context){
    id := c.Param("id")
    if err := DeleteTask(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    c.Status(http.StatusNoContent)
}

func CompleteTaskHandler(c *gin.Context){
    id := c.Param("id")
    if err := DoneTask(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    c.Status(http.StatusNoContent)
}
