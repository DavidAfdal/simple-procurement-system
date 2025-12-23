package validator

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func Validate(input interface{}) map[string]string {
	validate := validator.New()
	validate.RegisterValidation("date", dateValidationFunc)
	validate.RegisterValidation("date_with_time", dateWitTimeValidationFunc)

	err := validate.Struct(input)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return map[string]string{"error": "validasi input gagal: " + err.Error()}
		}

		validationErrors := make(map[string]string)
		inputVal := reflect.ValueOf(input)
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := inputVal.Type().FieldByName(err.Field())
			formTag := field.Tag.Get("form")
			jsonTag := field.Tag.Get("json")

			var fieldName string
			if formTag != "" {
				fieldName = formTag
			} else if jsonTag != "" {
				fieldName = jsonTag
			} else {
				fieldName = err.Field() // fallback to struct field name if both tags are missing
			}

			validationErrors[fieldName] = getErrorMessage(err)
		}
		return validationErrors
	}
	return nil
}

func dateValidationFunc(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	layout := "2006-01-02"
	_, err := time.Parse(layout, date)
	return err == nil
}
func dateWitTimeValidationFunc(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	layout := "2006-01-02 15:04:05"
	_, err := time.Parse(layout, date)
	return err == nil
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s wajib diisi", err.Field())

	case "email":
		return fmt.Sprintf("%s harus berupa alamat email yang valid", err.Field())

	case "min":
		return fmt.Sprintf("%s minimal %s karakter", err.Field(), err.Param())

	case "date":
		return fmt.Sprintf("%s harus berupa tanggal dengan format YYYY-MM-DD", err.Field())

	case "date_with_time":
		return fmt.Sprintf("%s harus berupa tanggal dan waktu dengan format YYYY-MM-DD HH:MM:SS", err.Field())

	case "oneof":
		fields := strings.Split(err.Param(), " ")
		return fmt.Sprintf("harus mengisi salah satu dari field berikut: %s", strings.Join(fields, ", "))

	default:
		return fmt.Sprintf("%s tidak valid", err.Field())
	}
}
