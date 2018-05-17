package api

import (
	"github.com/jeyem/godemy/model/site"
	"github.com/labstack/echo"
)

func subscribe(c echo.Context) error {
	form := new(SubscribeForm)

	if err := c.Bind(form); err != nil {
		return c.JSON(400, echo.Map{
			"error":  err.Error(),
			"status": FailedStatus,
		})
	}
	if err := form.validate(); err != nil {
		return c.JSON(400, echo.Map{
			"error":  err.Error(),
			"status": FailedStatus,
		})
	}
	sub := new(site.Subscribe)
	sub.Email = form.Email
	sub.Path = form.Path
	if err := sub.Save(); err != nil {
		return c.JSON(403, echo.Map{
			"error":  err.Error(),
			"status": FailedStatus,
		})
	}
	return c.JSON(200, echo.Map{
		"msg":    "successfully subscribed email",
		"status": SuccessStatus,
	})
}
