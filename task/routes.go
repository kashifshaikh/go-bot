package task

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Routes struct {
	Echo *echo.Echo
	Log  *zap.SugaredLogger
}

func RegisterRoutes(log *zap.SugaredLogger, e *echo.Echo) *Routes {
	r := &Routes{Log: log, Echo: e}

	tasks := r.Echo.Group("/tasks")
	tasks.GET("", r.getAllTasks)
	tasks.GET("/:id", r.getTask)
	tasks.POST("", r.createTask)
	tasks.POST("/:id", r.updateTask)
	tasks.DELETE("/:id", r.deleteTask)

	return r
}

func (r *Routes) createTask(c echo.Context) error {
	return c.String(http.StatusCreated, "createTask success")
}

func (r *Routes) getAllTasks(c echo.Context) error {
	return c.String(http.StatusOK, "tasks")
}

func (r *Routes) getTask(c echo.Context) error {
	return c.String(http.StatusOK, "task")
}

func (r *Routes) updateTask(c echo.Context) error {
	return c.String(http.StatusOK, "updateTask success")
}

func (r *Routes) deleteTask(c echo.Context) error {
	return c.String(http.StatusOK, "deleteTask success")
}
