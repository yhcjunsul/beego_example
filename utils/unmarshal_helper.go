package utils

import (
	"encoding/json"
	"html"
	"reflect"
)

func UnmarshalRequestJson(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	value := reflect.ValueOf(v).Elem()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		if field.Type() != reflect.TypeOf("") {
			continue
		}

		str := field.Interface().(string)

		field.SetString(html.EscapeString(str))
	}

	return nil
}
