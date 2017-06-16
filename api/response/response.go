package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"bitbucket.org/mirkorakic/engagedhits/errors"

	"github.com/labstack/echo"
)

//ErrorWithInfo returns error with information
func ErrorWithInfo(c echo.Context, info string, err interface{}) error {
	return ErrorWithStatus(c, err, info)

}

//Error returns the error
func Error(c echo.Context, err interface{}) error {
	return ErrorWithStatus(c, err, http.StatusOK)
}

//ErrorWithStatus handles all errors
func ErrorWithStatus(c echo.Context, err interface{}, info interface{}) error {
	data := map[string]interface{}{
		"success": false,
	}

	switch err.(type) {
	case errors.ErrorCode:
		e := errors.FromCode(err.(errors.ErrorCode))
		data["code"] = e.Code
		data["message"] = e.Message
	case *errors.Error:
		data["code"] = err.(*errors.Error).Code
		data["message"] = err.(*errors.Error).Message

	default:
		var str string
		switch err.(type) {
		case error:
			str = err.(error).Error()
		}
		data["code"] = 0
		data["message"] = str
		if strings.Contains(str, "[_VALIDATION_]") {
			msg := validation(str, c)

			switch msg.(type) {
			case error:
				return msg.(error)
			}

			data["code"] = errors.ErrorInvalid
			data["message"] = msg
		}
	}
	switch info.(type) {
	case string:
		data["message"] = info.(string) + " " + data["message"].(string)
	}

	return c.JSON(http.StatusOK, data)
}

func validation(str string, c echo.Context) interface{} {
	if strings.Contains(str, "[_VALIDATION_]") {
		list := strings.Split(str, "[_VALIDATION_]")
		if list[len(list)-1] == "" {
			list = list[:len(list)-1]
		}

		msg := map[string]interface{}{}

		for _, m := range list {
			var j map[string]interface{}
			err := json.Unmarshal([]byte(m), &j)
			if err != nil {
				return Error(c, err)
			}

			for k, v := range j {
				e := v.(map[string]interface{})

				if _, ok := e["invalid"]; ok {
					e["message"] = k + " does not validate as " + e["invalid"].(string)
				} else {
					if e[k].(string) == "empty" {
						e[k] = "required"
					}

					e["invalid"] = e[k].(string)
					e["message"] = k + " does not validate as " + e[k].(string)
					e["value"] = ""
					delete(e, k)
				}

				msg[strings.ToLower(k)] = e
			}

		}

		return msg
	}

	return nil
}

func Success(c echo.Context, data interface{}) error {
	if data == nil {
		data = map[string]interface{}{}
	}

	switch data.(type) {
	case map[string]interface{}:
		data.(map[string]interface{})["success"] = true
	}

	return c.JSON(http.StatusOK, data)
}

func SuccessCreated(c echo.Context, data interface{}) error {
	if data == nil {
		data = map[string]interface{}{}
	}

	switch data.(type) {
	case map[string]interface{}:
		data.(map[string]interface{})["success"] = true
		data.(map[string]interface{})["created"] = true
	}

	return c.JSON(http.StatusCreated, data)
}
