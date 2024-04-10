package utils

import (
	"reflect"
)

func CheckRequired(vars map[string]interface{}, textInit, textEnd string) string {
	var message = ""
	for key, value := range vars {
		if isEmpty(value) {
			message += textInit + " " + key + " " + textEnd + "\n\n"
		}
	}

	return message
}

func isEmpty(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr:
		return v.IsNil()
	default:
		return true
	}
}
