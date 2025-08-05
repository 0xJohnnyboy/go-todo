package task

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

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
        "id": task.ID,
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

func GetTaskStatsHandler(c *gin.Context){
	stats, err := GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
