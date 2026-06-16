package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	v "github.com/go-playground/validator/v10"
)

var validate *v.Validate

func init() {
	validate = v.New()
}

// ValidateStruct validates a struct using the singleton validator instance.
// It returns a single aggregated error message with human-readable messages
// for supported tags: required, min, max.
func ValidateStruct(s interface{}) error {
	if s == nil {
		return nil
	}

	if err := validate.Struct(s); err != nil {
		if ves, ok := err.(v.ValidationErrors); ok {
			var msgs []string
			for _, fe := range ves {
				field := jsonFieldName(s, fe.StructField())
				switch fe.Tag() {
				case "required":
					msgs = append(msgs, fmt.Sprintf("%s is required", field))
				case "min":
					msgs = append(msgs, fmt.Sprintf("%s must be at least %s characters", field, fe.Param()))
				case "max":
					msgs = append(msgs, fmt.Sprintf("%s must be at most %s characters", field, fe.Param()))
				default:
					msgs = append(msgs, fmt.Sprintf("%s is invalid", field))
				}
			}
			return errors.New(strings.Join(msgs, "; "))
		}
		return err
	}

	return nil
}

// jsonFieldName attempts to return the JSON tag name for a struct field.
// Falls back to the struct field name when a JSON tag isn't present.
func jsonFieldName(s interface{}, structFieldName string) string {
	if s == nil || structFieldName == "" {
		return structFieldName
	}

	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return structFieldName
	}

	if f, ok := t.FieldByName(structFieldName); ok {
		tag := f.Tag.Get("json")
		if tag == "" {
			return structFieldName
		}
		name := strings.Split(tag, ",")[0]
		if name == "" || name == "-" {
			return structFieldName
		}
		return name
	}

	return structFieldName
}
