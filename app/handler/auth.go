package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
  

func Login (c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if e := c.Validate(u); e != nil {
		log.Print(e)
		return c.JSON(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusOK, "!!!Hello, World!!!")
}
