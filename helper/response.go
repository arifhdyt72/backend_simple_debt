package helper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func FormatTime(timeData string) (time.Time, error) {
	var monthSting string
	year, month, day := time.Now().Date()
	var layoutISO = "2006-01-02 15:04:05"
	if int(month) >= 10 {
		monthSting = strconv.Itoa(int(month))
	} else {
		monthSting = fmt.Sprintf("0%s", strconv.Itoa(int(month)))
	}

	formatTime := fmt.Sprintf("%d-%s-%d %s", year, monthSting, day, timeData)
	toDateFormat, err := time.Parse(layoutISO, formatTime)
	if err != nil {
		return toDateFormat, err
	}

	return toDateFormat, nil
}
