package gopackage

import (
	"encoding/json"

	jsonIterator "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

// Custom JSON

type CustomJSONData struct{}

func CustomJSON() *CustomJSONData {
	return &CustomJSONData{}
}

func (cjd *CustomJSONData) Serialize(c echo.Context, i any, indent string) error {
	enc := json.NewEncoder(c.Response())

	if indent != "" {
		enc.SetIndent("", indent)
	}

	return enc.Encode(i)
}

func (cjd *CustomJSONData) Deserialize(c echo.Context, i any) error {
	var (
		ConfigCompatibleWithStandardLibrary = jsonIterator.Config{CaseSensitive: true}.Froze()
		customJSON                          = ConfigCompatibleWithStandardLibrary
	)

	return customJSON.NewDecoder(c.Request().Body).Decode(i)
}
