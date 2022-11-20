package controllers

import (
	"fmt"
	"latihan1/params/request"
	"latihan1/params/views"
	"latihan1/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// GetTodos godoc
// @Summary Get All TODOS
// @Schemes
// @Description get all todo
// @Tags TODOS
// @Accept json
// @Produce json
// @Success 200 {object} views.GetTodosSuccessSwag
// @Failure 404 {object} views.GetTodosFailureSwag
// @Router /todos [get]
func GetAll(c *gin.Context) {
	var todoRepo repositories.TodoRepo
	data, err := todoRepo.GetAllTodos()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	var resp = views.GeneralSuccessPayload{
		Status:  http.StatusOK,
		Message: "Get All Todos Success",
		Payload: data,
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// CreateTodo godoc
// @Summary Create TODO
// @Schemes
// @Description create todo
// @Tags TODOS
// @Accept json
// @Produce json
// @Param        request body      request.CreateTodo  true "CreateToDo"
// @Success 200 {object} views.CreateTodoSuccessSwag
// @Failure 400 {object} views.CreateTodoFailureSwag
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	var req request.CreateTodo

	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	var todoRepo repositories.TodoRepo
	err = todoRepo.InsertTodo(req)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	var resp = views.GeneralSuccessPayload{
		Status:  http.StatusCreated,
		Message: "CREATE TODO SUCCESS",
		Payload: true,
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// GetTodos godoc
// @Summary Get Todo By Id
// @Schemes
// @Description Get Todo By Id
// @Tags TODOS
// @Accept json
// @Produce json
// @Param        id path      int  true "Request Param"
// @Success 200 {object} views.GetTodoByIdSuccessSwag
// @Failure 404 {object} views.GetTodosFailureSwag
// @Router /todos/{id} [get]
func GetByID(c *gin.Context) {
	var id = c.Param("id")
	var todoRepo repositories.TodoRepo
	data, err := todoRepo.GetTodo(id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	var resp = views.GeneralSuccessPayload{
		Status:  http.StatusOK,
		Message: "GET TODO SUCCESS",
		Payload: data,
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// GetTodos godoc
// @Summary Update TODO By Id
// @Schemes
// @Description Update TODO By Id
// @Tags TODOS
// @Accept json
// @Produce json
// @Param        id path      		int  true "Request Param"
// @Param        request body      	request.CreateTodo  true "UpdateTodo"
// @Success 200 {object} views.UpdateTodoByIdSuccessSwag
// @Failure 400 {object} views.UpdateTodoFailureSwag
// @Router /todos/{id} [put]
func UpdateByID(c *gin.Context) {
	var id = c.Param("id")
	var req request.CreateTodo

	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	var todoRepo repositories.TodoRepo
	err = todoRepo.UpdateTodo(req, id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	var resp = views.GeneralSuccessPayload{
		Status:  http.StatusOK,
		Message: "UPDATE TODO SUCCESS",
		Payload: true,
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// DeleteTodo godoc
// @Summary Delete TODO by id
// @Schemes
// @Description Delete TODO by id
// @Tags TODOS
// @Accept json
// @Produce json
// @Param        id path      		int  true "Request Param"
// @Success 200 {object} views.DeleteTodoByIdSuccessSwag
// @Failure 404 {object} views.DeleteTodoFailureSwag
// @Router /todos/{id} [delete]
func DeleteByID(c *gin.Context) {
	var id = c.Param("id")
	var todoRepo repositories.TodoRepo
	err := todoRepo.DeleteTodo(id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	var resp = views.GeneralSuccessPayload{
		Status:  http.StatusOK,
		Message: "DELETE TODO SUCCESS",
		Payload: true,
	}

	c.JSON(resp.Status, resp)
}
