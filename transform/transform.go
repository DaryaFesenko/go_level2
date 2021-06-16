package transform

import (
	"errors"
	"reflect"
)

type TypeDontMatchError struct {
	message string
	info    string
}

func New(info string) error {
	return &TypeDontMatchError{
		message: "ERROR: value types do not match. ",
		info:    info,
	}
}

func (t *TypeDontMatchError) Error() string {
	return t.message + t.info
}

func MapToStruct(in interface{}, values map[string]interface{}) error {
	val := reflect.ValueOf(in)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return errors.New("ERROR: `in` is not a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		nameField := field.Name
		typeField := field.Type

		if value, ok := values[nameField]; ok {
			if value == nil {
				return errors.New(nameField + " is nil")
			}

			switch typeField.Kind() {
			case reflect.Int:
				v, okVal := value.(int)
				if !okVal {
					return New("field " + nameField + " - " + typeField.String() + ". map value - " + reflect.TypeOf(value).String())
				}
				val.Field(i).SetInt(int64(v))
			case reflect.Float64:
				v, okVal := value.(float64)
				if !okVal {
					return New("field " + nameField + " - " + typeField.String() + ". map value - " + reflect.TypeOf(value).String())
				}
				val.Field(i).SetFloat(v)
			case reflect.String:
				v, okVal := value.(string)
				if !okVal {
					return New("field " + nameField + " - " + typeField.String() + ". map value - " + reflect.TypeOf(value).String())
				}
				val.Field(i).SetString(v)
			}
		}
	}

	return nil
}
