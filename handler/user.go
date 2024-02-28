package handler

import (
	"github.com/labstack/echo/v4"
	"lessonGo/view/user"
)

type UserHandler struct {
}

func (hh *UserHandler) HandlerUserShow(c echo.Context) error {
	return render(c, user.Show())
}
