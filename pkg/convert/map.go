package convert

import (
	"fmt"
	"reflect"
)

func ObjectToMap[T ~map[string]interface{}](object interface{}, tagName string) (T, error) {
	defer func() {
		if recoveredError := recover(); recoveredError != nil {
			fmt.Println("recovered from: ", recoveredError)
		}
	}()

	values := reflect.ValueOf(object)
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}

	if values.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ObjectToMap only accepts structs; got %T", values)
	}

	resultMap := make(T, values.NumField())
	valueType := values.Type()

	for index := 0; index < values.NumField(); index++ {
		fieldType := valueType.Field(index)
		field := values.Field(index)

		if field.IsZero() {
			continue
		}

		if tagValue := fieldType.Tag.Get(tagName); tagValue != "" {
			resultMap[tagValue] = field.Interface()
		}
	}

	return resultMap, nil
}

func ObjectToStringMap(object interface{}, tagName string) (map[string]string, error) {
	mapObject, err := ObjectToMap[map[string]interface{}](object, tagName)
	if err != nil {
		return nil, err
	}

	mapString := make(map[string]string, len(mapObject))

	for key, value := range mapObject {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		mapString[strKey] = strValue
	}

	return mapString, nil
}
