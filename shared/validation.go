package shared

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func StringContainWhiteSpace(field validator.FieldLevel) bool {
	return strings.Contains(strings.TrimSpace(field.Field().String()), " ")
}

func JsonTagNameFunc(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}
