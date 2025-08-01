package gopackage

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// Custom Validator

type CustomValidatorData struct {
	validator *validator.Validate
}

func CustomValidator() *CustomValidatorData {
	return &CustomValidatorData{validator: validator.New()}
}

func (cvd *CustomValidatorData) Validate(i any) error {
	return cvd.validator.Struct(i)
}

// Custom HTTP Error for Echo Framework

type EchoCustomHTTPErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Data        any    `json:"data"`
	Internal    error  `json:"-"`
}

func EchoCustomHTTPErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)

	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, strings.ToUpper(strings.ReplaceAll(http.StatusText(http.StatusInternalServerError), " ", "_")))
	}

	logrus.WithFields(logrus.Fields{
		"tag": "GoPackage.Validator.EchoCustomHTTPErrorHandler.01",
	}).Error(strings.ToLower(cast.ToString(report.Message)))

	c.Logger().Error(report)

	c.JSON(report.Code, EchoCustomHTTPErrorResponse{
		Internal:    report.Internal,
		Code:        fmt.Sprintf("%04d", report.Code),
		Description: strings.ToUpper(strings.ReplaceAll(cast.ToString(report.Message), " ", "_")),
	})
}
