package api

import (
	"net/http"
	"strconv"
	"github.com/BREACH1247/todo-api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func (api *TodoAPI) AddTodoHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        var todo models.Todo
        if err := c.ShouldBindJSON(&todo); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        
        if err := api.AddTodo(todo); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        
        c.JSON(http.StatusCreated, gin.H{"message": "Todo item added successfully"})
    }
}


func (api *TodoAPI) GetTodoHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        todoID, err := gocql.ParseUUID(id)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
            return
        }

        todo, err := api.GetTodo(todoID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, todo)
    }
}

func (api *TodoAPI) UpdateTodoHandler() gin.HandlerFunc{
    return func(c *gin.Context) {
        id := c.Param("id")
        todoID, err := gocql.ParseUUID(id)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
            return
        }

        var todo models.Todo
        if err := c.ShouldBindJSON(&todo); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        todo.ID = todoID
        if err := api.UpdateTodo(todo); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Todo item updated successfully"})
    }
}

func (api *TodoAPI) DeleteTodoHandler() gin.HandlerFunc{
    return func(c *gin.Context) {
        id := c.Param("id")
        todoID, err := gocql.ParseUUID(id)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
            return
        }

        if err := api.DeleteTodo(todoID); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo item"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Todo item deleted successfully"})
    }
}

func (api *TodoAPI) GetTodoListHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        pageStr := c.Query("page")
        pageSizeStr := c.Query("pageSize")
        status := c.Query("status")

        page, err := strconv.Atoi(pageStr)
        if err != nil {
            page = 1
        }
        pageSize, err := strconv.Atoi(pageSizeStr)
        if err != nil {
            pageSize = 10
        }

        offset := (page - 1) * pageSize

        todos, err := api.GetTodoList(page, pageSize, offset, status)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, todos)
    }
}



