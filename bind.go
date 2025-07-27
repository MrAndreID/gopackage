package gopackage

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Echo Bind Request

type request interface {
	Validate() any
}

func EchoBindRequest(c echo.Context, i request) error {
	var tag string = "GoPackage.Bind.EchoBindRequest."

	if i == nil {
		return nil
	}

	binder := new(echo.DefaultBinder)

	if err := binder.BindBody(c, i); err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "01",
			"error": err.Error(),
		}).Error("failed to default bind body")

		return echo.NewHTTPError(http.StatusBadRequest, strings.ToUpper(strings.ReplaceAll(http.StatusText(http.StatusBadRequest), " ", "_")))
	}

	if err := binder.BindQueryParams(c, i); err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "02",
			"error": err.Error(),
		}).Error("failed to default bind query param")

		return echo.NewHTTPError(http.StatusBadRequest, strings.ToUpper(strings.ReplaceAll(http.StatusText(http.StatusBadRequest), " ", "_")))
	}

	if err := binder.BindPathParams(c, i); err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "03",
			"error": err.Error(),
		}).Error("failed to default bind path param")

		return echo.NewHTTPError(http.StatusBadRequest, strings.ToUpper(strings.ReplaceAll(http.StatusText(http.StatusBadRequest), " ", "_")))
	}

	if err := binder.BindHeaders(c, i); err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "04",
			"error": err.Error(),
		}).Error("failed to default bind header")

		return echo.NewHTTPError(http.StatusBadRequest, strings.ToUpper(strings.ReplaceAll(http.StatusText(http.StatusBadRequest), " ", "_")))
	}

	e := i.Validate()

	if e == nil {
		return nil
	}

	return echo.NewHTTPError(http.StatusBadRequest, e)
}
